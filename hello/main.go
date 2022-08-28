package main

import "fmt"

const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "

func Hello(name string, language string) string {
	var prefix string
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
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
