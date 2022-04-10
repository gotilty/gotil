## IsAssigned

It checks if there is a value corresponding to the given parameter. Returns boolean.

> ***all primitive types, arrays & slices are supported.***

```go
gotil.IsAssigned(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ejXoOFiSa1n)

```go
var a = []string{"gotil is perfect!"}
fmt.Println(gotil.IsAssigned(a))
```

```go
//output:
true
```

## IsNotAssigned

It checks if there is a value corresponding to the given parameter. Returns boolean.

> ***all primitive types, arrays & slices are supported.***

```go
gotil.IsNotAssigned(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/b0CeMGU32HT)

```go
var a = []string{"gotil is perfect!"}
fmt.Println(gotil.IsNotAssigned(a))
```

```go
//output:
false
```