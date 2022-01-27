//Package mate nos ayuda a comprobar que sabes sumar
package mate

//Sum suma una cantidad ilimitada de numeros enteros
fun sum(xi ...int) int{
	var s intfor _, v:=range xi{
		s+=v
	}
	return s
}