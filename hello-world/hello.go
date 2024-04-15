package main

import (
	"fmt"
)

var HELLO_IN_LANGUAGES = map[string]string{
	"english": "Hello",
	"bisaya":  "Hoy",
}

func Hello(n, language string) string {
	if n == "" {
		n = "World"
	}

	prefix := HELLO_IN_LANGUAGES[language]
	if prefix == "" {
		prefix = HELLO_IN_LANGUAGES["english"]
	}

	return fmt.Sprintf("%s, %s", prefix, n)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
