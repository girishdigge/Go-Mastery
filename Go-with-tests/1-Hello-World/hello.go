package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const enHelloPrefix = "Hello, "
const spHelloPrefix = "Ola, "
const frHelloPrefix = "Bonjure, "

func Hello(lang string, name string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + name
}
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frHelloPrefix
	case spanish:
		prefix = spHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return prefix
}
func main() {
	fmt.Println(Hello("", "Girish"))
}
