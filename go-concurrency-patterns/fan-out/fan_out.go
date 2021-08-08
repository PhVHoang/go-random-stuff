package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var myNumbers [10]int
	for i := 0; i < 10; i++{ 
		rand.Seed(time.Now().UnixNano())
		myNumbers[i]=rand.Intn(50)
	}
	
	mychannelOut := channelGenerator(myNumbers)

	mychannel1 := double(mychannelOut)
	mychannel2 := double(mychannelOut)

	mychannelIn := fanIn(mychannel1, mychannel2)


	for i := 0; i < len(myNumbers); i++ {
		fmt.Println(<-mychannelIn)
	}
}

func channelGenerator(numbers [10]int) <-chan string {
	channel := make(chan string)
	go func() {
		for _, i := range numbers {
			channel  <-  strconv.Itoa(i)
		}
		close(channel)
	}()
	return channel 
}

func double(inputchannel <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := range inputchannel {
			num, err := strconv.Atoi(i)
			 if err != nil {
      			
  			 }
			 channel <- fmt.Sprintf("%d * 2 = %d", num,num*2)
		}
		close(channel)
	}()
	return channel
}


func fanIn(inputchannel1, inputchannel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case message1 := <-inputchannel1:  
				channel <- message1
			case message2 := <-inputchannel2:  
				channel <- message2
			}
		}
	}()
	return channel
}
