package main

import "fmt"

//iota inicia en 0 y continua inscrementando las constantes
const (
	a = 2020 + iota
	b = 2020 + iota
	c = 2020 + iota
	d = 2020 + iota
)

func main() {
	//iota incrementará los valores de la constante
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
