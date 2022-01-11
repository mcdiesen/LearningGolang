//Eliminando elamentos de mapas
package main
func main() {
m:=map[]int(
	"batman":32,
	"Robin":27,

)
fmt.Println(m)
delete(m,"Batman")
fmt.Println(m)
if v, ok:=m["Robin"];ok {
	fmt.Println("Se borro la llave con valor", v)
	delete(m,"Batman")
}
fmt.Println(m)
}