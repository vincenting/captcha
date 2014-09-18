package captcha

import (
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Draw("十 二 减 去 13 等 于", "../images/result.gif")
	}
}
