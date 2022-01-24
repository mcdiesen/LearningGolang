//Builtin error
//Manejo de Errores
package main

import "fmt"

type erroPer struct {
	info string
}

func (ep erroPer) Error() string {
	return fmt.Sprintf("Aquí esta el error %v", ep.info)
}
func main() {
	e1 := erroPer{
		info: "Necesito sentirme bien",
	}
	foo(e1)
}

func foo(e error) {
	//assertion - afirmar que e.info es de tipo erroPer
	fmt.Println("Foo corrio - ", e.(erroPer).info, "\n")
}
