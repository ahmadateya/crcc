package main

import (
	"time"
)

func main() {
	// run infinite loop
	l1 := listenToMalPort("7222")
	for {
		time.Sleep(5 * time.Second) // wait 5 sec
	}
	l1.Close()
}
