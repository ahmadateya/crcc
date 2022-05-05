package main

import (
	"fmt"
	"time"
)

func main() {
	// run infinite loop
	for {
		time.Sleep(5 * time.Second) // wait 1 sec
		fmt.Println("Hello World!")
	}
}
