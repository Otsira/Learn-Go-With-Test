package main

import "fmt"

const enGreetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enGreetingPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
