package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	koreanHelloPrefix  = "Annyeong, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingHelloPrefix(language) + name
}

func greetingHelloPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Korean":
		prefix = koreanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
