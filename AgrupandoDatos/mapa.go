//almacenamiento llave valor
//utiliza mapas para busquedas
//tipo de datos desordenados
package main

func main() {
m:=map[string]int("batman":32,
"Robin":27,)
fmt.Println(m)
fmt.Println("Batman")
fmt.Println("Eduar")
v,ok:=m["Eduar"]
fmt.Println(v)
fmt.Println(ok)

if v,ok:=m["Eduar";ok]{
	fmt.Println("Imprimiendo desde el if", v)
}

}
