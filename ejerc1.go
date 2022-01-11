package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {

	a := fmt.Sprintf("%v\t %v\t %v", x, y, z)
	fmt.Println(a)
}
