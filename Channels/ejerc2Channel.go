//Canal send only chan<- int
//Canal recive only <-chan int
package main

import "fmt"

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 43
	}()
	fmt.Println(cs)
}
