package main

import (
	"fmt"
	"time"
	"sync"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// TODO: completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		// TODO: simular trabajo
		time.Sleep(300 * time.Millisecond)

		ch <- x * x
	}()
	return ch
}

func main() {
	// TODO: crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)

	// TODO: Opción 1: esperar cada futuro secuencialmente
	fmt.Println("f1:", <-f1)
    fmt.Println("f2:", <-f2)
    fmt.Println("f3:", <-f3)

	
	// TODO: Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.
	f4 := asyncCuadrado(5)
    f5 := asyncCuadrado(6)
    combinado := fanIn(f4, f5)

    for v := range combinado {
	    fmt.Println("fan-in recibió:", v)
		}
}

func fanIn(canales ...<-chan int) <-chan int {
	salida := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(canales))

	for _, c := range canales {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				salida <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(salida)
	}()

	return salida
}
	

