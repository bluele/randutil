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
  fmt.Println("choiced:", randutil.Choice(ss))
}

/* output:

before: [a b c d e]
after: [c e a d b]
choiced: e
*/
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>