package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		//_ = listener.Close()
		listener.Close()
	}()
	t.Logf("Bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
			//return err
		}

		go func(c net.Conn) {
			defer c.Close()

			// Your code would handle the connection here.

		}(conn)
	}
}
