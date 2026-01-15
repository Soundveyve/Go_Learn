package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

/*
Пример использования функции Add.
Комментарий внутри кода ВЛИЯЕТ на тесты: без этого комментария код только скоплириуется, но не выполнится
Удалите комментарий // Output: 6, затем запустите go test -v, и вы увидите, что ExampleAdd больше не выполняется.
*/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
