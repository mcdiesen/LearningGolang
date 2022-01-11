//concurrencia
//waitGroup
//pkg runtime
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Variable wg para acceder a los métodos Add, Wait y Done de Waitgroup
var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())

	wg.Add(1) //agrega una 1 goroutine

	go foo()
	bar()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GOROUTINE\t", runtime.NumGoroutine())

	wg.Wait() //vigila y cuando no hay rutinas main puede finalizar
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	//Done indica la finalización de la Goroutine
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {

		fmt.Println("bar", i)
	}
}
