//Package mate nos ayuda a comprobar que sabes sumar
package mate

//Sum suma una cantidad ilimitada de numeros enteros
func sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
