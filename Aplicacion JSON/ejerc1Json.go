//ejercicio 1 app json
//funcion Marshall

package main

import (
	"encoding/json"
	"fmt"
)

type Usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := Usuario{
		Nombre: "James",
		Edad:   32,
	}

	u2 := Usuario{
		Nombre: "MoneyPenny",
		Edad:   30,
	}

	//crear el slice de tipo Usuario
	usuarios := []Usuario{
		u1,
		u2,
	}
	fmt.Println(usuarios)
	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println("Error", err)
	}
	//impresion del slice en formato json
	fmt.Println(string(bs))

}
