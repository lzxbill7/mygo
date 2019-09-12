package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func flipAndInvertImage(A [][]int) [][]int {
	for _, s := range A {
		for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
			s[i], s[j] = s[j]^1, s[i]^1
		}
	}

	return A
}

func TestFlipAndInvertImage(t *testing.T) {
	Convey("test", t, func() {
		t.Log(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
		t.Log(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
	})
}
