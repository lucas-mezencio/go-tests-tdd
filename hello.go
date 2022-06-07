package main

import "go-test-tdd/slices"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func main() {
	slices.SumAll([]int{1, 2}, []int{3, 4})
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}
