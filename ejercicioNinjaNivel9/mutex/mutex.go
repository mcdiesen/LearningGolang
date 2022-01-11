//mutex uso del metodo lock y unlock
//para que las goroutines no se afecten en lectura y escritura
//no hay racecondition
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	incremento := 0
	gs := 100

	wg.Add(gs)

	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			//a partir del m.Lock() se bloquea el código a continuación
			v := incremento
			v++
			incremento = v
			fmt.Println(incremento)
			//desbloqueo del código
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El valor final de incremento es:", incremento)
}
