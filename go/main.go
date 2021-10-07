package main

import (
	"fmt"
	"time"
)

func doThing(reason string) {
	fmt.Printf("do thing because %s...\n", reason)
	time.Sleep(time.Second)
	fmt.Println("done")
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	trigger := make(chan bool)

	go func() {
		for {
			select {
			case <-trigger:
				doThing("trigger")
			case <-ticker.C:
				doThing("ticker")
			}
		}
	}()

	time.Sleep(time.Second)
	trigger <- true
	time.Sleep(5 * time.Second)
	trigger <- true
	time.Sleep(1 * time.Second)
	trigger <- true
	trigger <- true
	trigger <- true
	time.Sleep(5 * time.Second)
}
