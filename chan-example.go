package main

import "fmt"

func main() {
	i := make(chan int)
	go func(){
		i<- 1 //write to the channel
	}()
	//value_i := <-i //read from the channel
	fmt.Println(<-i)

}