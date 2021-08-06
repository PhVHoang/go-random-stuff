package main

import "fmt"

func fibonacci(n int) chan int {
    mychannel := make(chan int)
    go func() {
        defer close(mychannel)
        k := 0
        for i, j := 0, 1; k < n ; k++ {
            mychannel <- i
            i, j = i+j,i
            
        }
    }()
    return mychannel
}

func main() {
  
    for i := range fibonacci(10) {
       //do anything with the nth term while the fibonacci()
       //is computing the next term
        fmt.Println(i)
    }
}
