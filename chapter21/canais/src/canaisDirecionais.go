package main

import "fmt"

func main() {

	channelBi := make(chan int) // um canal bidirecional
	//channelSend := make(<-chan int) // um canal que so manda valores
	//channelReceive := make(chan<- int) // um canal que so recebe valores

	go send(channelBi) // uma funcao que manda valor para o canal

	receive(channelBi) // uma funcao que recebe valor do canal

}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor do canal é:", <-r)
}
