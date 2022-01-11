//tipo String
package main

import "fmt"

func main() {
	s1 := "Hola Mundo"
	s2 := `Esto es una linea
	partida`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es %T\n", s2)

	//slice
	bs := []byte(s1)
	fmt.Println(bs)
	//imprime el esquema de código almacenado en el slice
	fmt.Printf("El tipo del slice es: %T\n", bs)

	//imprimirlo en utf8
	fmt.Println("")
	for i := 0; i < len(s1); i++ {
		fmt.Printf(" %#U ", s1[i])
	}
	//con range
	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("El indice  %d el valor es %x\n", i, v)
	}

}
