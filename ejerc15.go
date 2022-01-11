package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 10 {
			//break sale del ciclo for si cumple la condicion if
			break
		}
		fmt.Println(i)
		i++
	}
}
