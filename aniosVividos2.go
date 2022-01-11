package main

import "fmt"

func main() {
	cu := 1982
	for {
		if cu > 2021 {
			break
		}
		fmt.Println(cu)
		cu++
	}
}
