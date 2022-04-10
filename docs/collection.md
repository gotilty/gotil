## Every

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.Every(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/sgH-q0eERn4)

```go
 m1 := []float64{5, 10.5, 10, 10, 20, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}

    if result, err := gotil.Every(m1, func(a interface{}, i interface{}) interface{} {
        return int
    }); err == nil {
        for k, v := range result.(map[float64][]float64) {
            fmt.Println(k, v)
        }
    }
```

```go
//output:
false
```

## Filter

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.Filter(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w)

```go
data := []int64{-100, -5, 30, 100}
result, _ := gotil.FilterBy(data, func(val interface{}, i int) bool {
    if val.(int64) > 0 {
        return true
    } else {
        return false
    }
})
fmt.Println(result)
```

```go
//output:
[30 100]
```

## FindBy

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.Filter(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w)

```go
data := []int64{-100, -5, 30, 100}
result, _ := gotil.FilterBy(data, func(val interface{}, i int) bool {
    if val.(int64) > 0 {
        return true
    } else {
        return false
    }
})
fmt.Println(result)
```

```go
//output:
[30 100]
```
