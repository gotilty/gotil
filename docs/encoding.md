## HexEncode

Encodes the given parameter to hex type. Returns string.

> ***all primitive types, arrays & slices are supported.***

```go
gotil.HexEncode(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/A2-AiVmzzHC)

```go
test_data := string("gotil is perfect!")
converted, err := gotil.HexEncode(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
676f74696c206973207065726665637421
```

## HexEncode

Decodes the given parameter to string type. Returns string.

> ***all primitive types, arrays & slices are supported.***

```go
gotil.HexDecode(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/3Jh3AcrS9vQ)

```go
test_data := string("676f74696c206973207065726665637421")
converted, err := gotil.HexDecode(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
gotil is perfect!
```