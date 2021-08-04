package main 
import "fmt"

func PrefixSum(my_array,my_output []int ,parent chan int) {
	if len(my_array) < 2 {
		parent<-my_array[0]
		my_output[0] = my_array[0] + <-parent
		
	} else if len(my_array)<1{
		parent<-0
		<-parent
	} else  {
		mid:=len(my_array)/2
		left:= make(chan int)
		right:=make(chan int)
		go PrefixSum(my_array[:mid],my_output[:mid],left)
		go PrefixSum(my_array[mid:],my_output[mid:],right)
		leftsum:=<-left
		parent<- leftsum +<-right
		fromleft:= <-parent
		left<-fromleft
		right<-fromleft + leftsum
		<-left
		<-right

	}
	parent<-0
}

func main () {
	data:= []int{1,2,3,4}
	output:= make([]int,len(data))
	parent :=make(chan int)
	go PrefixSum(data,output,parent)
	sum:= <-parent
	fromleft:=0
	parent<-fromleft
	donezero:=<-parent
	fmt.Println(data,output,sum,donezero)
}
