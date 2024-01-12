package main

import (
	"fmt"
)

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
	// y := strconv.FormatInt(101, 2)
	// fmt.Println(y)

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

func slices() {
	a := make([]int, 4)

	a[0] = 1
	a[1] = 1
	a[2] = 1
	a[3] = 1

	b := a[0:1]

	// Makes sense, b is a reference
	// [2 1 1 1] [2]
	b[0] = 2

	fmt.Println(a, b)

	b = append(b, 3)
	b[0] = 3

	// Um... why did append change a?
	// [3 3 1 1] [3 3]
	fmt.Println(a, b)
}

func main() {
	// fmt.Println(Hello("world", ""))
	slices()
}
