package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf("Hello, golang. I'm return :з \n")
	fmt.Println(Hello("world"))
}
