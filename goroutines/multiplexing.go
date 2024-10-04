package goroutines

import (
	"fmt"
	"time"
)

// Multiplex
// With the below code:
//
//	ch1 sends messages every 500 milliseconds.
//	ch2 sends messages every 1 second.
//
// This means you’re more likely to receive 2 messages from ch1 before getting 1 from ch2
// since ch1 produces messages twice as fast as ch2.
func Multiplex() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendMessagesToChannel(ch1, "Message from channel 1", 800*time.Millisecond)
	go sendMessagesToChannel(ch2, "Message from channel 2", time.Second)

	for i := 1; i <= 6; i++ {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println("Received: ", msg1)
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println("Received: ", msg2)
			}
		}
	}
	//time.Sleep(time.Second * 3)

	fmt.Println("All messages received form channels !! ")

}

func sendMessagesToChannel(ch chan<- string, message string, delay time.Duration) {
	for i := 1; i <= 3; i++ {
		//fmt.Printf("Value of i is %d from %s \n", i, message)
		time.Sleep(delay)
		// Add value to channel
		ch <- fmt.Sprintf("%s %d", message, i)

	}
	close(ch)
}
