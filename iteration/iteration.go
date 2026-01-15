package iteration

import (
	"strings"
)

func Repeat(symbol string, countRepeat int) string {

	if countRepeat == 0 || countRepeat == 1 {
		return symbol
	}

	/*
		Строки в Go неизменяемы, то есть при каждой операции конкатенации, как в функции Repeat,
		происходит копирование памяти для размещения новой строки.
		Это влияет на производительность, особенно при интенсивной конкатенации строк.
		Стандартная библиотека предоставляет тип strings.Builder,
		который минимизирует копирование данных в памяти.
		Он реализует метод WriteString, который можно использовать для объединения строк:
	*/
	var repeated strings.Builder
	for i := 0; i < countRepeat; i++ {
		repeated.WriteString(symbol)
	}
	return repeated.String()
}
