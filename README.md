# randutil

Random variable utility library for golang.

## Install

```
$ go install github.com/bluele/randutil
```

## Example

```go
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

/* output:

before: [a b c d e]
after: [c e a d b]
choiced: e
3 samples: [d e b]
*/
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>