package main

import "fmt"

var greetings = map[string]string{
	"en": "Hello, ",
	"es": "Hola, ",
	"fr": "Bonjour, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	// Quirks of go, but seems nice
	//https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if greet, ok := greetings[language]; ok {
		return greet + name
	}
	return greetings["en"] + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
