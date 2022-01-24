//Error manejo de errores

package main

import "fmt"

func main() {

	var s1, s2 string
	//la funcion Println, devuelve dos valores n (#byes) y error, por tanto,
	//se pueden verificar errores
	n, err := fmt.Println("Ingrese dos valores de tipo cadena")

	if err != nil {
		fmt.Println("Error", n, err)
	}
	fmt.Println("Nombre:")
	_, err = fmt.Scan(&s1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Apellido:")
	_, err = fmt.Scan(&s2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hola", s1, s2)

}
