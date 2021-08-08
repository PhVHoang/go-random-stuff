package main

import ( 
	"fmt"
	"math/rand"
	"time"
)

func updatePosition(name string) <-chan string { 
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name , i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return positionChannel
}

func fanIn(mychannel1, mychannel2 <-chan string) <-chan string {
	mychannel := make(chan string) 

	go func() { 
		for {
			mychannel <- <-mychannel1 
		}
	}()

	go func() { 
		for {
			mychannel <- <-mychannel2
		}
	}()

	return mychannel
}


func main() {
	positionsChannel := fanIn(updatePosition("Legolas :"), updatePosition("Gandalf :"))
	

	for i := 0; i < 10; i++ {
		fmt.Println(<-positionsChannel)
	}

	fmt.Println("Done with getting updates on positions.")
}



