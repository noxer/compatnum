# compatnum
Data types for storing integers as JSON strings or numbers depending on the size.

## docs
https://pkg.go.dev/github.com/noxer/compatnum

## example
### code
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/noxer/compatnum"
)

type MyStruct struct {
	MyInt compatnum.Int
}

func main() {
	fmt.Println("Testing unmarshaling...")

	data1 := []byte(`{"MyInt": 123}`)
	data2 := []byte(`{"MyInt": "123"}`)

	var a, b MyStruct

	json.Unmarshal(data1, &a)
	json.Unmarshal(data2, &b)

	fmt.Printf("%s -> a.MyInt = %d\n", data1, a.MyInt)
	fmt.Printf("%s -> b.MyInt = %d\n", data2, b.MyInt)

	fmt.Println("\nTesting marshaling...")

	a = MyStruct{MyInt: 42}
	b = MyStruct{MyInt: 8589934591}

	p1, _ := json.Marshal(a)
	p2, _ := json.Marshal(b)

	fmt.Printf("%+v -> %s\n", a, p1)
	fmt.Printf("%+v -> %s\n", b, p2)

	fmt.Println("\nTesting custom MaxInt...")

	compatnum.MaxInt = 10
	fmt.Printf("compatnum.MaxInt = %d\n", compatnum.MaxInt)

	a = MyStruct{MyInt: 10}
	b = MyStruct{MyInt: 11}

	p1, _ = json.Marshal(a)
	p2, _ = json.Marshal(b)

	fmt.Printf("%+v -> %s\n", a, p1)
	fmt.Printf("%+v -> %s\n", b, p2)
}
```

### output
```
Testing unmarshaling...
{"MyInt": 123} -> a.MyInt = 123
{"MyInt": "123"} -> b.MyInt = 123

Testing marshaling...
{MyInt:42} -> {"MyInt":42}
{MyInt:8589934591} -> {"MyInt":"8589934591"}

Testing custom MaxInt...
compatnum.MaxInt = 10
{MyInt:10} -> {"MyInt":10}
{MyInt:11} -> {"MyInt":"11"}
```