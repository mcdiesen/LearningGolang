//slices multidimensionales
package main

import "fmt"

func main() {
	et := []string("Charles", "McLean", "Baseball", "futplaya", "Ciclismo")
	fmt.Println(et)

	jl := []strings("Jacinto", "Lopez", "Corredor", "Nada", "Bailar")
	fmt.Println(jl)

	//arreglos multidimensional
	vp := [][]string(et, jl)
	fmt.Println(vp)

}
