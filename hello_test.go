package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Проверка, что функция возвращает Hello, SOME_NAME", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Печать 'Hello World' если строка с именем была пустая", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("Возвращает: %q Ожидается: %q", got, want)
		}
	})
}
