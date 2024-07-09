package main

import "fmt"

const (
	salutationEN = "Hello, "
	salutationES = "Hola, "
	salutationFR = "Bonjour, "
	salutationIT = "Ciao, "

	es = "Spanish"
	en = "English"
	fr = "French"
	it = "Italian"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return salutationSelector(lang) + name + "!"
}

func salutationSelector(lang string) (salutation string) {
	switch lang {
	case en:
		salutation = salutationEN
	case es:
		salutation = salutationES
	case fr:
		salutation = salutationFR
	case it:
		salutation = salutationIT
	default:
		salutation = salutationEN
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
