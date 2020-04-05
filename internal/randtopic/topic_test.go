package topic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTranslateNumberToHanzi(t *testing.T) {
	assert.Equal(t, simpleTranslateNumberToHanzi(5), []string{"五"})
	assert.Equal(t, simpleTranslateNumberToHanzi(15), []string{"十", "五"})
	assert.Equal(t, simpleTranslateNumberToHanzi(20), []string{"二", "十"})
	assert.Equal(t, simpleTranslateNumberToHanzi(25), []string{"二", "十", "五"})
}

func TestConvertIntToStringWithSeed(t *testing.T) {
	assert.Equal(t, convertIntToStringWithSeed(5, 0), []string{"五"})
	assert.Equal(t, convertIntToStringWithSeed(5, 1), []string{"5"})
	assert.Equal(t, convertIntToStringWithSeed(25, 0), []string{"二", "十", "五"})
	assert.Equal(t, convertIntToStringWithSeed(25, 1), []string{"25"})
}

func TestBuildTopicSubject(t *testing.T) {
	assert.Equal(t, buildTopicSubject(25, "+", 25, [3]int{0, 0, 0}), []string{"二", "十", "五", "加", "二", "十", "五"})

	assert.Equal(t, buildTopicSubject(5, "+", 5, [3]int{0, 0, 0}), []string{"五", "加", "五", "是", "多", "少", "呢"})
	assert.Equal(t, buildTopicSubject(5, "+", 5, [3]int{1, 0, 0}), []string{"5", "加", "五", "是", "多", "少", "呢"})
	assert.Equal(t, buildTopicSubject(5, "+", 5, [3]int{1, 0, 1}), []string{"5", "加", "上", "五", "是", "多", "少"})
	assert.Equal(t, buildTopicSubject(5, "+", 5, [3]int{1, 1, 1}), []string{"5", "加", "上", "5", "是", "多", "少"})
}
