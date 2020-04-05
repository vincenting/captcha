package topic

import (
	"math/rand"
	"strconv"
	"time"
)

// CaptchaTopic defines the content and expected result of the captcha
type CaptchaTopic struct {
	Subject []string // The captcha content
	Result  int      // The expected calculate result
}

func randInRange(from int, to int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(to-from) + from
}

func simpleTranslateNumberToHanzi(number int) []string {
	if number < 11 {
		return []string{numbersToHanzi[number]}
	}

	tens, unit := number/10, number%10
	if tens == 1 {
		return []string{numbersToHanzi[10], numbersToHanzi[unit]}
	}
	if unit == 0 {
		return []string{numbersToHanzi[tens], numbersToHanzi[10]}
	}
	return []string{numbersToHanzi[tens], numbersToHanzi[10], numbersToHanzi[unit]}
}

func convertIntToStringWithSeed(number int, sampleIndex int) []string {
	samples := [][]string{
		simpleTranslateNumberToHanzi(number),
		[]string{strconv.Itoa(number)},
	}
	return samples[sampleIndex]
}

func buildTopicSubject(left int, operator string, right int, randoms [3]int) []string {
	leftChars := convertIntToStringWithSeed(left, randoms[0])
	rightChars := convertIntToStringWithSeed(right, randoms[1])
	numberCharsLen := len(leftChars) + len(rightChars)

	if numberCharsLen == 6 {
		return append(append(leftChars, mathOperationsIHanzi[operator][0]...), rightChars...)
	}

	operationCharLen := randoms[2] + 1
	equalSymbol := equalSymbolIHanzi[captchaExpectedLength-numberCharsLen-operationCharLen-1]
	return append(
		append(leftChars, mathOperationsIHanzi[operator][operationCharLen-1]...),
		append(rightChars, equalSymbol...)...)
}

// RandGenerate generate a random captcha structure
func RandGenerate() CaptchaTopic {
	operator := supportOperations[randInRange(0, len(supportOperations))]
	randoms := [3]int{randInRange(0, 1), randInRange(0, 1), randInRange(0, 1)}

	if operator == "+" {
		left, right := randInRange(0, 30), randInRange(0, 30)
		return CaptchaTopic{
			buildTopicSubject(left, operator, right, randoms),
			left + right,
		}
	}

	left := randInRange(10, 50)
	right := randInRange(0, left-1)
	return CaptchaTopic{
		buildTopicSubject(left, operator, right, randoms),
		left - right,
	}
}
