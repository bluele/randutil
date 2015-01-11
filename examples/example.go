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
}
