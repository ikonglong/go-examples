package time

import (
	"fmt"
	"testing"
	"time"
)

func useTickerIncorrectly() *time.Ticker {
	ticker := time.NewTicker(3 * time.Second)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			fmt.Println("Tick ... ")
		}

		fmt.Println("Ticker is about to stop")
	}(ticker)

	return ticker
}

func TestUseTickerIncorrectly(t *testing.T) {
	ticker := useTickerIncorrectly()
	time.Sleep(15 * time.Second)
	// Stop turns off a ticker. After Stop, no more ticks will be sent.
	// Stop does not close the channel, to prevent a concurrent goroutine
	// reading from the channel from seeing an erroneous "tick".
	ticker.Stop()
}

func useTickerCorrectly() chan bool {
	ticker := time.NewTicker(3 * time.Second)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer func() {
			fmt.Println("Before calling stop()")
			ticker.Stop()
			fmt.Println("After calling stop()")
		}()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick ...")
			case stop := <-stopChan:
				if stop {
					fmt.Println("Ticker is about to stop")
				}
			}
		}
	}(ticker)

	return stopChan
}

func TestUseTickerCorrectly(t *testing.T) {
	ch := useTickerCorrectly()
	time.Sleep(15 * time.Second)
	ch <- true
	close(ch)
}
