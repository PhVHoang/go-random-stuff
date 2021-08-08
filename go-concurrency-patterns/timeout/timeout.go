package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dynamite := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		dynamite <- "Dynamite Diffused!"
	}()

	for {
		select {
		case s := <-dynamite:
			fmt.Println(s)
			return
		case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
			fmt.Println("Dynamite Explodes!")
			return
		}
	}
}
