//uso de struc anonimos
package main

import "fmt"

func main() {
	//estructuras anonimas
	s := struct {
		nombre  string
		amigos  map[string]int //tipo mapa
		bebidas []string       //tipo slice
	}{
		//asignación de valores a los campos
		nombre: "Charles",
		amigos: map[string]int{
			"Carito": 222,
			"Adrian": 444,
			"Condor": 555,
		},
		bebidas: []string{
			"agua",
			"chicha",
		},
	}
	fmt.Println(s.nombre)
	fmt.Println(s.amigos)
	fmt.Println(s.bebidas)

	for k, v := range s.amigos {
		fmt.Println(k, v)
	}
	for i, v := range s.bebidas {
		fmt.Println(i, v)
	}
}
