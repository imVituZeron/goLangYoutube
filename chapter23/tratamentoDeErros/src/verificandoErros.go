package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("asdfasdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James")
	io.Copy(f, r)
}
