package main

import "fmt"

const englishHelloPrefix string = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name

}

func main() {
	fmt.Println(Hello("World"))
}
