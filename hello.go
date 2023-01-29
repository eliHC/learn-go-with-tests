package main

import "fmt"

func main() {
	fmt.Println(Hello("Carl"))
}

func Hello(name string) string {
	return "Hello, " + name
}
