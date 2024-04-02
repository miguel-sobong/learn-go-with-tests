package main

import "fmt"

const ENGLISH_HELLO_PREFIX = "Hello, "

func Hello(n string) string {
	if n == "" {
		return "Hello, World"
	}
	return fmt.Sprintf("%s%s", ENGLISH_HELLO_PREFIX, n)
}

func main() {
	fmt.Println(Hello("Chris"))
}
