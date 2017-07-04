/*
The MIT License (MIT)

Copyright (c) 2017 Jun Kimura

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

package randutil

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

func makeDefaultRand(seed int64) *rand.Rand {
	return rand.New(&LockedSource{src: rand.NewSource(1).(rand.Source64)})
}

// rand function
var Rand = makeDefaultRand(time.Now().UnixNano())

// Choose a random element from a non-empty slice.
func Choice(slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("choice: not slice")
	}
	n := in.Len()
	if n == 0 {
		panic("choice: length is 0")
	}
	i := Rand.Intn(n)
	return in.Index(i).Interface()
}

// Shuffle elements order in slice.
func Shuffle(slice interface{}) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("shuffle: not slice")
	}
	n := in.Len()
	swap := reflect.Swapper(slice)
	for i := n - 1; i > 0; i-- {
		j := int(Rand.Float32() * float32(i+1))
		swap(i, j)
	}
}

// Chooses k unique random elements from a population slice.
// Returns a new list containing elements from the population while leaving the original population unchanged.
// The resulting list is in selection order so that all sub-slices will also be valid random samples.
func Sample(slice interface{}, k int) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("sample: not slice")
	}
	n := in.Len()
	if !(0 <= k && k <= n) {
		panic("sample: k is larger than slice size")
	}
	nf := float32(n)

	setSize := 21
	if k > 5 {
		setSize += int(math.Pow(4, math.Ceil(_log(float64(k*3), 4))))
	}

	sl := reflect.SliceOf(in.Type().Elem())
	result := reflect.MakeSlice(sl, k, k)
	if n <= setSize {
		pin := reflect.MakeSlice(sl, n, n)
		reflect.Copy(pin, in)
		for i := 0; i < k; i++ {
			j := int(Rand.Float32() * float32(n-i))
			result.Index(i).Set(pin.Index(j))
			pin.Index(j).Set(pin.Index(n - i - 1))
		}
	} else {
		selected := make(map[interface{}]interface{})
		for i := 0; i < k; i++ {
			j := int(Rand.Float32() * nf)
			for {
				if _, ok := selected[j]; !ok {
					break
				}
				j = int(Rand.Float32() * nf)
			}
			selected[j] = true
			result.Index(i).Set(in.Index(j))
		}
	}
	return result.Interface()
}

func _log(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}
