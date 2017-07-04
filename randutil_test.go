/*
The MIT License (MIT)

Copyright (c) 2015 Jun Kimura

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package randutil_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bluele/randutil"
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
	ar := make([]int, len(dt))
	copy(ar, dt)
	randutil.Shuffle(dt)
	if !isDiffSlice(dt, ar) {
		t.Error("shuffle: not changed error")
	}
	sort.Slice(dt, func(i, j int) bool {
		return dt[i] < dt[j]
	})
	if !reflect.DeepEqual(ar, dt) {
		t.Errorf("%v != %v", ar, dt)
	}
}

func TestSmallSample(t *testing.T) {
	size := 50
	sampleSize := 10
	dt := makeIntRange(size)
	ar := randutil.Sample(dt, sampleSize).([]int)

	if len(ar) != sampleSize {
		t.Error("sample: sample size is invalid")
	}

	if len(dt) != size {
		t.Error("sample: sample source size has changed")
	}
}

func TestBigSample(t *testing.T) {
	size := 10000000
	sampleSize := 10
	dt := makeIntRange(size)
	ar := randutil.Sample(dt, sampleSize).([]int)

	if len(ar) != sampleSize {
		t.Error("sample: sample size is invalid")
	}

	if len(dt) != size {
		t.Error("sample: sample source size has changed")
	}
}

func makeIntRange(n int) []int {
	ar := []int{}
	for i := 0; i < n; i++ {
		ar = append(ar, i)
	}
	return ar
}

func isDiffSlice(a []int, b []int) bool {
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
