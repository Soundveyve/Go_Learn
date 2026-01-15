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
