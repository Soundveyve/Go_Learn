package massiveandslice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Сложение 5 чисел в коллекции", func(t *testing.T) {
		/*синтаксис для массивов:
		numbers := [5]int{1, 2, 3, 4, 5}

		альтернативные объявления массивов:

		С указанием типа отдельно
		var numbers [5]int = [5]int{1, 2, 3, 4, 5}

		Без указания размера (компилятор выведет его сам)
		numbers := [...]int{1, 2, 3, 4, 5}

		Объявление и последующее присваивание
		var numbers [5]int
		numbers = [5]int{1, 2, 3, 4, 5}

		Частичная инициализация (остальные элементы будут равны нулю)
		numbers := [5]int{1, 2, 3}

		*/
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %v want %v, given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Корректное высчитывание сумм", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{3, 9})
		var want []int = []int{2, 9}
		checkSum(t, got, want)
	})
	t.Run("Корректное высчитывание пустого слайса", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 9})
		var want []int = []int{0, 9}
		checkSum(t, got, want)
	})
}
