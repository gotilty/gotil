## ToString

Converts the given parameter to string type. Returns string.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.ToString(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/NlfntmMO9f3)

```go
test_data := float64(922.2022)
converted, err := gotil.ToString(test_data)
if err != nil {
fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
922.2022
```

## Join

Joins slices or arrays with given separator. Returns string.

> ***all primitive types are supported.***

```go
gotil.Join(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/tWHN5QiKVNJ)

```go
stringArray := []int{1, 2, 3, 4, 5}
converted, err := gotil.Join(stringArray, "***")
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
1***2***3***4***5
```

## ToInt64

Converts the given parameter to int64 type.

> ***all primitive types are supported.***

```go
gotil.ToInt64(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/73gbSw9SWlk)

```go
test_data := string("9223372036854775807")
converted, err := gotil.ToInt64(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
9223372036854775807
```

## ToInt32

Converts the given parameter to int32 type.

> ***all primitive types are supported.***

```go
gotil.ToInt32(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/joNZaWGoctY)

```go
test_data := ("65537")
converted, err := gotil.ToInt32(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
65537
```

## ToInt16

Converts the given parameter to int16 type.

> ***all primitive types are supported.***

```go
gotil.ToInt16(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/-hJu2TuE1bh)

```go
test_data := string("32766")
converted, err := gotil.ToInt16(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
32766
```

## ToInt8

Converts the given parameter to int8 type.

> ***all primitive types are supported.***

```go
gotil.ToInt8(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/IA57XgBpzpX)

```go
test_data := string("123")
converted, err := gotil.ToInt8(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
123
```

## ToInt

Converts the given parameter to int type.

> ***all primitive types are supported.***

```go
gotil.ToInt(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/tG3oG4y3vxx)

```go
test_data := string("123")
converted, err := gotil.ToInt(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
123
```

## ToUint64

Converts the given parameter to uint64 type.

> ***all primitive types are supported.***

```go
gotil.ToUint64(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/cLNUvSztDLx)

```go
test_data := string("18446744073709551615")
converted, err := gotil.ToUint64(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
18446744073709551615
```

## ToUint32

Converts the given parameter to uint32 type.

> ***all primitive types are supported.***

```go
gotil.ToUint32(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/EX3r_xcrwIb)

```go
test_data := string("123")
converted, err := gotil.ToUint32(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
123
```

## ToUint16

Converts the given parameter to uint16 type.

> ***all primitive types are supported.***

```go
gotil.ToUint16(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/mZFVzm4VcgV)

```go
test_data := string("32766")
converted, err := gotil.ToUint16(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
32766
```

## ToUint8

Converts the given parameter to uint8 type.

> ***all primitive types are supported.***

```go
gotil.ToUint8(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/t37E-ixaDb0)

```go
test_data := string("123")
converted, err := gotil.ToUint8(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
123
```

## ToUint

Converts the given parameter to uint type.

> ***all primitive types are supported.***

```go
gotil.ToUint(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/FlJWI5pyON7)

```go
test_data := string("156")
converted, err := gotil.ToUint(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
156
```


## ToFloat64

Converts the given parameter to float64 type.

> ***all primitive types are supported.***

```go
gotil.ToFloat64(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/OeQudordvEp)

```go
test_data := string("32766.1254")
converted, err := gotil.ToFloat64(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
32766.1254
```

## ToFloat32

Converts the given parameter to float32 type.

> ***all primitive types are supported.***

```go
gotil.ToFloat32(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/zxBIJ_iBXtz)

```go
test_data := string("32766.125")
converted, err := gotil.ToFloat32(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
32766.125
```