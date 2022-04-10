## Each

Iterates an array of values by running each element in the given array thru iteratee.

> ***array & slice are supported***

```go
gotil.Each(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/sgH-q0eERn4)

```go
m1 := []int{123, 6787}

gotil.Each(m1, func(a interface{}, i int) {
    fmt.Println(i, a)
})
```

```go
//output:
0 123
1 6787
```

## Map

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***

```go
gotil.Map(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/sgH-q0eERn4)

```go
m1 := []float64{5, 10.5, 10, 10, 20, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}

if result, err := gotil.Map(m1, func(a interface{}, i int) interface{} {
    return math.Floor(a.(float64)) * 2
}); err == nil {
    for _, v := range result.([]float64) {
        fmt.Println(v)
    }
}
```

```go
//output:
10
20
20
20
40
40
20
200
8
10
10
200
```