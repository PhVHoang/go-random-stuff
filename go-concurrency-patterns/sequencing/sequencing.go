package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CookInfo struct {
	foodCooked     string
	waitForPartner chan bool
}

func cookFood(name string) <-chan CookInfo {
	cookChannel := make(chan CookInfo)
	wait := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			cookChannel <- CookInfo{
				fmt.Sprintf("%s %s", name, "Done"), wait,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-wait
		}
	}()

	return cookChannel
}

func fanIn(channel1, channel2 <-chan CookInfo) <-chan CookInfo {
	channel := make(chan CookInfo)
	go func() {
		for {
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()
	return channel
}

func main() {
	gameChannel := fanIn(cookFood("Player 1: "), cookFood("Player 2: "))

	for round := 0; round < 3; round++ {
		food1 := <-gameChannel
		fmt.Println(food1.foodCooked)

		food2 := <-gameChannel
		fmt.Println(food2.foodCooked)

		food1.waitForPartner <- true
		food2.waitForPartner <- true

		fmt.Printf("Done with round %d\n", round+1)
	}

	fmt.Printf("Done with the competition")
}
