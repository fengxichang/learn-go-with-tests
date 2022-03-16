package hello_world

import "fmt"

func main() {
	fmt.Println(Hello("neil", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const franceHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "France":
		prefix = franceHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}