package iteration

func Repeat(symbol string, countRepeat int) (repeatString string) {

	if countRepeat == 1 || countRepeat == 0 {
		return symbol
	}

	for i := 0; i < countRepeat; i++ {
		repeatString += symbol
	}

	return repeatString
}
