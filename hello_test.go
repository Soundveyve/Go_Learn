package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Проверка, что функция возвращает Hello, SOME_NAME", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Печать 'Hello World' если строка с именем была пустая", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// t testing.TB - интерфейс для вспомогательных функций в тесте или бенчмарке
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Возвращает: %q Ожидается: %q", got, want)
	}
}
