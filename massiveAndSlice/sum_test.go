package massiveandslice

import (
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
		letMeFeedBack(t, got, want, numbers)
	})
}

func letMeFeedBack(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
