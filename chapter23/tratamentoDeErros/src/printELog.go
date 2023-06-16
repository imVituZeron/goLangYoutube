package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happend", err)
		//log.Println("err happend", err) // o log marca um time
		//log.Fatalln(err) // da um exit status code
		log.Panicln(err) // tambem da um exit status code
	}
}
