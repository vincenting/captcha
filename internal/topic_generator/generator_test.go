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
