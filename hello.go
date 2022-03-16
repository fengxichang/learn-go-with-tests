package main

import "fmt"

func main() {
	fmt.Println(Hello("neil"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}