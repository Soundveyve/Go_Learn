package main

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return greetingPrefix(language) + name
}

/*
При добавление имени переменной в качестве возвращаемого значения
Go автоматиески ициализирует под неё знаение, поэтому его явно можно не прописывать
*/
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {

}
