package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[5] = 100
	fmt.Println(x)
	x = append(x, 200)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 300)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
