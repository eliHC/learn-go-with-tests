package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Carl"))
}

func Hello(name string) string {
	if name == "" {
		name = "Friend"
	}

	return englishHelloPrefix + name
}
