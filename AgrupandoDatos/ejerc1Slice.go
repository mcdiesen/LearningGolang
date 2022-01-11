package main

import "fmt"

func main() {
	//Arreglo sin la capacidad es Slice
	//x := [5]int(41, 42, 43, 44, 45)
	x := []int(41, 42, 43, 44, 45)

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("%T", x)
}
