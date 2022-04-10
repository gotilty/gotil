## Md5

Hashes the given parameter with md5. Returns string.

> ***all primitive types, arrays & slices are supported.***

```go
gotil.Md5(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/jpk-1HKLyw9)

```go
test_data := string("gotil is perfect!")
converted, err := gotil.Md5(test_data)
if err != nil {
    fmt.Println(err)
}
fmt.Println(converted)
```

```go
//output:
f9beacbfe9c1d1d9a067c7f14d4e7e37
```