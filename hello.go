package main

import "fmt"

const salutationEN = "Hello, "
const salutationES = "Hola, "
const salutationFR = "Bonjour, "

const es = "Spanish"
const en = "English"
const fr = "French"

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
	if lang == fr {
		salutation = salutationFR
	}
	return salutation + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
