package topic

const captchaExpectedLength = 7

var numbersToHanzi []string = []string{
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

var supportOperations []string = []string{"+", "-"}

var mathOperationsIHanzi map[string][][]string = map[string][][]string{
	"+": [][]string{
		[]string{"加"},
		[]string{"加", "上"},
	},
	"-": [][]string{
		[]string{"减"},
		[]string{"减", "掉"},
	},
}

var equalSymbolIHanzi [][]string = [][]string{
	[]string{"是"},
	[]string{"等", "于"},
	[]string{"是", "多", "少"},
	[]string{"是", "多", "少", "呢"},
}
