package randutil

import (
	"math/rand"
	"reflect"
	"time"
)

func Choice(seq interface{}) interface{} {
	in := reflect.ValueOf(seq)
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

func Shuffle(seq interface{}) {
	in := reflect.ValueOf(seq)
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
