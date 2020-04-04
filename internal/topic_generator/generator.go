package topic

func simpleTranslateNumberToHanzi(number int) []string {
	if number < 11 {
		return []string{numbersIHanzi[number]}
	}

	tens, unit := number/10, number%10

	if tens == 1 {
		return []string{numbersIHanzi[10], numbersIHanzi[unit]}
	}
	if unit == 0 {
		return []string{numbersIHanzi[tens], numbersIHanzi[10]}
	}

	return []string{numbersIHanzi[tens], numbersIHanzi[10], numbersIHanzi[unit]}
}
