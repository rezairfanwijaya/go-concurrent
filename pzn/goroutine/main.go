package main

import "fmt"

func main() {
	i := 1
	for {
		go HelloWorld(i)
		i++
		if i == 100000 {
			break
		}
	}
	fmt.Println("pres enter")
	fmt.Scanln()
}

func HelloWorld(index int) {
	fmt.Println("Hello world ke ", index)
}
