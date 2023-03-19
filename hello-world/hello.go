package main

import (
	"fmt"
)

const (
	englishHelloPrefix  = "Hello, "
	japaneseHelloPrefix = "こんにちは、"
	frenchHelloPrefix   = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
