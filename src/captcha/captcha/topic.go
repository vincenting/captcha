package captcha

import (
	"strconv"
	"strings"
)

const (
	captchaLen int = 7
)

var num2chinese []string = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
var operator2chinese map[string][]string = map[string][]string{
	"+": []string{"加", "加 上"},
	"-": []string{"减", "减 掉"},
}
var eql2chinese []string = []string{"是", "等 于", "是 多 少", "是 多 少 呢"}

type NumSt struct {
	Size int
	Cn   string
}

type Topic struct {
	Subject string
	Result  string
}

func Num2CN(num int) NumSt {
	if num < 11 {
		return NumSt{1, num2chinese[num]}
	}
	a, b := num/10, num%10
	if a == 1 {
		return NumSt{
			2,
			strings.Join([]string{num2chinese[10], num2chinese[b]}, " "),
		}
	}
	if b == 0 {
		return NumSt{
			2,
			strings.Join([]string{num2chinese[a], num2chinese[10]}, " "),
		}
	}
	return NumSt{
		3,
		strings.Join([]string{num2chinese[a], num2chinese[10], num2chinese[b]}, " "),
	}
}

func TopicParse(le NumSt, rt NumSt, operator string, optLen int) string {
	currentLen := le.Size + rt.Size
	optArr := operator2chinese[operator]
	if currentLen == 6 {
		return strings.Join([]string{le.Cn, optArr[0], rt.Cn}, " ")
	}
	if currentLen == 5 && optLen == 2 {
		return strings.Join([]string{le.Cn, optArr[1], rt.Cn}, " ")
	}
	eqlLen := captchaLen - optLen - currentLen
	return strings.Join([]string{le.Cn, optArr[optLen-1], rt.Cn, eql2chinese[eqlLen-1]}, " ")
}

func RandNumParse(num int) NumSt {
	rd := random(0, 5)
	if rd < 2 {
		return Num2CN(num)
	}
	return NumSt{1, strconv.Itoa(num)}
}

func RandTopic() Topic {
	operateInt := random(0, 5)
	var le, rt, result int
	var operator string
	if operateInt < 3 {
		operator = "+"
		le, rt = random(0, 30), random(0, 30)
		result = le + rt
	} else {
		operator = "-"
		le = random(10, 50)
		rt = random(0, le)
		result = le - rt
	}
	optLen := 2
	if random(0, 5) < 3 {
		optLen = 1
	}
	return Topic{
		TopicParse(RandNumParse(le), RandNumParse(rt), operator, optLen),
		strconv.Itoa(result),
	}
}
