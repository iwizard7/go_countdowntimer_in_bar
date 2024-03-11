package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Countdown Timer")

	countdown(99)
}

func countdown(seconds int) {
	for i := seconds; i >= 0; i-- {
		fmt.Printf("%d seconds left\n", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Time's up!")
}
