package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())       // mostrando quantos nucleos tem o processados
	fmt.Println(runtime.NumGoroutine()) //mostrando os processos

	wg.Add(2) // adicionando duas gorotines no waitgroups.

	go func1() // chamando a execucao
	go func2() // chamando a execucao

	fmt.Println(runtime.NumGoroutine())

	wg.Wait() // esperando
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done() // avisando que terminou de ser executado.
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done() // avisando que terminou de ser executado.
}
