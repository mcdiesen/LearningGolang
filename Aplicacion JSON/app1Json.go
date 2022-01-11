//uso de JSON
//json.org
//marshall y unmarshall
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Rojo", "Amarillo", "Ruby", "Pink"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("Error", err)
	}
	os.Stdout.Write(b)
}
