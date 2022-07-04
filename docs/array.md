## Map

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***

```go
gotil.Map(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/6-D3kaa2UNS)

```go
result := gotil.Map([]int{10, 20}, func(v, i int) int {
	return v * v
})
fmt.Println(result)
```

```go
//output:
[100 400]
```