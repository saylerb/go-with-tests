package main

import "fmt"

const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "

func Hello(name string, language string) string {
	var prefix string
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	if name == "" {
		return prefix + "World"
	}
	return prefix + name

}

func main() {
	fmt.Println(Hello("World", ""))
}
