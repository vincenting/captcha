package topic

const captchaExpectedLength = 7

var numbersIHanzi []string = []string{
	"零",
	"一",
	"二",
	"三",
	"四",
	"五",
	"六",
	"七",
	"八",
	"九",
	"十",
}

var mathOperationsIHanzi map[string][]string = map[string][]string{
	"+": []string{"加", "加 上"},
	"-": []string{"减", "减 掉"},
}

var equalSymbolIHanzi []string = []string{"是", "等 于", "是 多 少", "是 多 少 呢"}
