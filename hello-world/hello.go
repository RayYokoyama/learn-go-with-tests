package main

import (
	"fmt"
)

const (
	englishHelloPrefix  = "Hello, "
	japaneseHelloPrefix = "こんにちは、"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Japanese" {
		return japaneseHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
