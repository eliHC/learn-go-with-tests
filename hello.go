package main

import (
	"fmt"
)

const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const ptBrHelloPrefix = "Olá, "

func main() {
	fmt.Println(Hello("Carl", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "Friend"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	case "pt-br":
		prefix = ptBrHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}
