package main

import "fmt"

//uso de switch
func main() {
	deporteFav := "surfing"
	switch deporteFav {
	case "baseball":
		fmt.Println("Ve al estadio")
	case "Natacion":
		fmt.Println("Ve a la playa")
	case "surfing":
		fmt.Println("Ve a Hawai")
	}
}
