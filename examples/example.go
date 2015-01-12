package main

import (
	"fmt"
	"github.com/bluele/randutil"
)

func main() {
	var ss []string = []string{"a", "b", "c", "d", "e"}
	fmt.Println("before:", ss)
	randutil.Shuffle(ss)
	fmt.Println("after:", ss)

	fmt.Println("choice:", randutil.Choice(ss))

	samples := randutil.Sample(ss, 3).([]string)
	fmt.Println("3 samples:", samples)
}
