//Ejercicio Ninja Nivel 9
//go rutine, la primera goroutine es main
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Numero de CPU", runtime.NumCPU())
	fmt.Println("Sistema operativo: ", runtime.GOOS)
	fmt.Println("Numero de Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera goroutine")
		wg.Done()
	}()
	func() {
		fmt.Println("Hola desde la segunda goroutine")
		wg.Done()
	}()
	fmt.Println("Numero de CPU", runtime.NumCPU())
	fmt.Println("Numero de Goroutine", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Numero de CPU", runtime.NumCPU())
	fmt.Println("Numero de Goroutine", runtime.NumGoroutine())
}
