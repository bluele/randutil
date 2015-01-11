package randutil_test

import (
	"github.com/bluele/randutil"
	"testing"
	"time"
)

func TestChoice(t *testing.T) {
	dt := makeIntRange(1000)
	dv := randutil.Choice(dt).(int)

	if !containValue(dt, dv) {
		t.Error("choice: not contained error")
	}
}

func TestShuffle(t *testing.T) {
	dt := makeIntRange(1000)
	ar := make([]int, cap(dt))
	copy(ar, dt)
	randutil.Shuffle(dt)
	if !isChangedSlice(dt, ar) {
		t.Error("shuffle: not changed error")
	}
}

func makeIntRange(n int) []int {
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, int(time.Now().UnixNano()))
	}
	return ar
}

func isChangedSlice(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}

func containValue(a []int, dv int) bool {
	for _, v := range a {
		if v == dv {
			return true
		}
	}
	return false
}
