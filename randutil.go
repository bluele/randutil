package randutil

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

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
	i := rand.Intn(n)
	return in.Index(i).Interface()
}

// Shuffle elements order in slice.
func Shuffle(slice interface{}) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("shuffle: not slice")
	}
	n := in.Len()
	for i := n - 1; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())
		j := int(rand.Float32() * float32(i+1))
		vi, vj := in.Index(i), in.Index(j)
		v := vi.Interface()
		vi.Set(vj)
		vj.Set(reflect.ValueOf(v))
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
			rand.Seed(time.Now().UnixNano())
			j := int(rand.Float32() * float32(n-i))
			result.Index(i).Set(pin.Index(j))
			pin.Index(j).Set(pin.Index(n - i - 1))
		}
	} else {
		selected := make(map[interface{}]interface{})
		for i := 0; i < k; i++ {
			j := int(rand.Float32() * nf)
			for {
				if _, ok := selected[j]; !ok {
					break
				}
				j = int(rand.Float32() * nf)
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
