package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	fmt.Println("Misturando...")
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}
	fmt.Println("")

	for { // laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
