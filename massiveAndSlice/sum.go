package massiveandslice

func Sum(numbers []int) int {
	var sum int = 0
	/*
		range позволяет выполнять итерацию по массиву.
		На каждой итерации range возвращает два значения — индекс и значение.
		Мы решили игнорировать значение индекса, используя _ пустой идентификатор
	*/
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
Для написания функции, которая принимает разное количество аргументов одного типа
достаточно использовать ... перед указанием типа.
Для всех входящих значений будет инициализироваться срез с переданными в функцию значениями
*/
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] += Sum(numbers)
	}
	return sums
}
