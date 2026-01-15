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
