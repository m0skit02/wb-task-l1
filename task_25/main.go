package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	ch := make(chan struct{})

	go func() {
		timer := time.NewTimer(d)
		<-timer.C
		close(ch)
	}()

	<-ch
}

func main() {
	fmt.Println("Start:", time.Now())
	mySleep(2 * time.Second)
	fmt.Println("End:", time.Now())
}
