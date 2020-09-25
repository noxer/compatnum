# compatnum
Data types for storing integers as JSON strings or numbers depending on the size.

## docs
https://pkg.go.dev/github.com/noxer/compatnum

## example
```go
package main

type MyStruct struct {
    MyInt compatnum.Int
}

func main() {
    data1 := `{"MyInt": 123}`
    data2 := `{"MyInt": "123"}`

    var a, b MyStruct

    json.Unmarshal(data1, &a)
    json.Unmarshal(data2, &b)

    fmt.Printf("a.MyInt = %d; b.MyInt = %d\n", a.MyInt, b.MyInt)
}
```
