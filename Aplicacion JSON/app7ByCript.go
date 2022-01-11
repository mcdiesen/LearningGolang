//BCrypt
//PkG para encriptación
//Hashing para contraseñas
//para importar el pkg bcrypt usar: go get - u golang.org/x/crypto/bcrypt
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	//string de la contraseña la que se debe almacenar en la BD
	fmt.Println(string(bs))

	claveLogin := `clave123`

	//func CompareHashAndPassword
	//compara la clave del hash inicial contra la nueva clave "ingresada"
	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("No te puedes loguear", err)
		return
	}
	fmt.Println("Te haz logueado")
}
