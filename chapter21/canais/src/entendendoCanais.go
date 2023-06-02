package main

import "fmt"

func main() {

	channel := make(chan []int) // crando um canal com valor int

	go func() {
		channel <- []int{50, 70, 77} // passando o int 50 para o canal
	}()

	fmt.Println(<-channel) // passando o canal para o println

}
