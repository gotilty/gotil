## Each

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***

```go
gotil.GroupBy(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/sgH-q0eERn4)

```go
 m1 := []float64{5, 10.5, 10, 10, 20, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}

    if result, err := gotil.GroupBy(m1, func(a interface{}, i interface{}) interface{} {

        return math.Floor(a.(float64))
    }); err == nil {
        for k, v := range result.(map[float64][]float64) {
            fmt.Println(k, v)
        }
    }
```

```go
//output:
10 [10.5 10 10 10.75]
20 [20 20]
100 [100 100.0001]
4 [4.23]
5 [5 5.15 5.99]
```
