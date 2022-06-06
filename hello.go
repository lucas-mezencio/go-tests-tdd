package main

import "fmt"

func main() {
	fmt.Println(Hello("Lucas"))
}

func Hello(s string) string {
	return "Hello, " + s
}
