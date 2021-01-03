package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func mergeTwoChannels(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

// First way: using N goroutines to merge n channels

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, c := range cs {
		go func() {
			for v := range c {
				out <- v
			}
		}()
	}
	return out
}

func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := mergeTwoChannels(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
