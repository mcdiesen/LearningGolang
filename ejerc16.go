package main

import "fmt"

func main() {
	//control de flujo iterativo
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
