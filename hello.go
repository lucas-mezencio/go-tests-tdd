package main

import "fmt"

const helloFunctionPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Lucas"))
}

func Hello(s string) string {
	if s == "" {
		s = "World"
	}
	return helloFunctionPrefix + s
}
