package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("При передаче одного символа, он же и возвращается", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "a"
		getMeFeedback(t, got, want)
	})
	t.Run("Возвращается строка с нужным кол-вом повторов", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		getMeFeedback(t, got, want)
	})
}

func getMeFeedback(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

/*
Бенчмарки используются для теста средней скорости выполнения функции
Loop() возвращает true до тех пор, пока выполняется бенчмарк.
go test -bench=. (или, если Windows Powershell, go test -bench=".")
Измеряется время только выполнения тела цикла;
код настройки и очистки автоматически исключается из измерения времени выполнения теста
*/
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 0)
	}
}
