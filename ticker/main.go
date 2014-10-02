package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	id := 0
	c := func() chan int {
		c := make(chan int)
		go func() {
			for {
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				id++
				c <- id
			}
		}()
		return c
	}()
	for {
		select {
		case res := <-c:
			log.Printf("Got %v\n", res)
		case <-time.Tick(time.Second):
			log.Printf("Still Alive")
		}
	}
}
