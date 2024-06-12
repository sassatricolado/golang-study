package main

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const espanhol = "Spanish"
const french = "French"
const portuguese = "Português"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixLanguage(language) + name
}

func prefixLanguage(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case espanhol:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
