package main

import "fmt"

const salutationEN = "Hello, "
const salutationES = "Hola, "

const es = "Spanish"
const en = "English"

func Hello(name string, lang string) string {
	salutation := ""
	if name == "" {
		name = "World"
	}
	if lang == en || lang == "" {
		salutation = salutationEN
	}
	if lang == es {
		salutation = salutationES
	}
	return salutation + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
