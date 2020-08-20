package moreduplicatedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalCase(t *testing.T) {
	arr := []int{2, 4, 3, 1, 0, 4, 6, 5}
	ans := Findaduplicate(arr)
	assert.EqualValues(t, "Duplicate found 4", ans)

}
func TestCornerCase(t *testing.T) {
	arr := []int{0, 0}
	ans := Findaduplicate(arr)
	assert.EqualValues(t, "Duplicate found 0", ans)

}
func TestBadInputCase(t *testing.T) {
	arr := []int{2, 4, 3, 1, 0, 4, 16, 5}
	ans := Findaduplicate(arr)
	assert.EqualValues(t, "Negative or too big numbers!", ans)

}
