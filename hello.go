package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

func main() {
	fmt.Println(Hello("Lucas", "Spanish"))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}
