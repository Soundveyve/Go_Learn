package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Функция возвращает Hello, SOME_NAME при передаче имени", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Печать 'Hello World' если строка с именем была пустая", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("При передаче пустого префикса используется английский", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Ответ на испанском", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Ответ на французском", func(t *testing.T) {
		got := Hello("Petro", "French")
		want := "Bonjour, Petro"
		assertCorrectMessage(t, got, want)
	})
}

// t testing.TB - интерфейс для вспомогательных функций в тесте или бенчмарке
func assertCorrectMessage(t testing.TB, got, want string) {

	/*
		Указание, что функция вспомогательная
		При ошибке в коде ссылка строку будет не во вспомогательной функции, а в тесте
	*/

	t.Helper()
	if got != want {
		t.Errorf("Возвращает: %q Ожидается: %q", got, want)
	}
}
