package ch03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// Start a server
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	// Start the listener's goroutine.
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			// When you close the listener, the listener's Accept method will immediately unblock and return an error.
			// This error isn't necessarily a failure, so you simply log it and move on. It doesn't cause your test to fail.
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			// Spin off each connection into its own goroutine.
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					// After receiving the FIN packet, the Read method returns the io.EOF error,
					// indicating to the listener’s code that you closed the client side of the connection.
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					t.Logf("Received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	// Start a client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	// Initiate a graceful termination of the connection from the client side.
	conn.Close()
	<-done
	listener.Close()
	<-done
}
