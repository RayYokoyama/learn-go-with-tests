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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
