## Every TODO

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.Every(val, f)
```

### examples TODO

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

## FindByFromIndex

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.Filter(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w) 

```go
data := []int64{-100, -5, 30, 100, 200, 300, 400}
result, _ := gotil.FindByFromIndex(data, func(val interface{}, i int) bool {
    if val.(int64) > 30 {
        return true
    } else {
        return false
    }
}, 5)
fmt.Println(result)
```

```go
//output:
400
```
## FindLastBy

Inversely checks whether the predicate is true for all elements of the collection. Iteration is stopped once predicate returns false.

> ***array & slice are supported***
```go
gotil.FindLastBy(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ay1huuG2ipx) 

```go
data := []int64{-100, -5, 30, 100}
result, _ := gotil.FindLastBy(data, func(val interface{}, i int) bool {
    if val.(int64) > 29 {
        return true
    } else {
        return false
    }
})
fmt.Println(result)
```

```go
//output:
100
```

## GroupBy

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***
```go
gotil.GroupBy(val, f)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ay1huuG2ipx) 

```go
m1 := []string{"eight", "nine", "four", "seven"}

if result, err := gotil.GroupBy(m1, func(a interface{}, i interface{}) interface{} {
    return len(strings.Split(a.(string), ""))
}); err == nil {
    for k, v := range result.(map[int][]string) {
        fmt.Println(k, v)
    }
}
```

```go
//output:
5 [eight seven]
4 [nine four]
```

## Includes

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***
```go
gotil.Includes(array, val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd) 

```go
m1 := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
result, _ := gotil.Includes(m1, "gotil")
fmt.Println(result)
```

```go
//output:
false
```

## Reduce TODO

Creates a map composed of keys generated from the results of running each element of collection through the iteratee function.

> ***array & slice are supported***
```go
gotil.Includes(array, val)
```

### examples TODO

>💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd) 

```go
m1 := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
result, _ := gotil.Includes(m1, "gotil")
fmt.Println(result)
```

```go
//output:
false
```


## Shuffle

Shuffles the list randomly. A new array is returned.

> ***array & slice are supported***
```go
gotil.Shuffle(data)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/jxNSYukWBZR) 

```go
data := []int64{-100, -5, 30, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9}
newData, _ := gotil.Shuffle(data)
fmt.Println(newData)
```

```go
//output:
[8 6 4 -100 7 2 9 -5 100 30 3 1 5]
```
## ShuffleSeed

Shuffles the list based on the given seed. A new array is returned.

> ***array & slice are supported***
```go
gotil.ShuffleSeed(data, seed)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd) 

```go
// Use "seed := time.Now().UnixNano()" for random result
seed := int64(58239238)
//Seed you will get the same sequence of pseudo­random numbers
// each time you run the program.

data := []int64{-100, -5, 30, 100}
newData, _ := gotil.ShuffleSeed(data, seed)
fmt.Println(newData)
```

```go
//output:
[-5 100 -100 30]
```

## Size

Returns size of given array or slice.

> ***array & slice are supported***
```go
gotil.Size(data)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/EY4ca2fMJhR) 

```go
data := []int64{-100, -5, 30, 100}
newData, _ := gotil.Size(data)
fmt.Println(newData)
```

```go
//output:
4
```

## Sort

Returns sorted version of given array or slice.

> ***array, slice, map & struct are supported***
```go
gotil.Sort(data)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/jmZZjaS6gqC) 

```go
data := []int64{100, 30, -100, -5}
newData, _ := gotil.Sort(data)
fmt.Println(newData)
```

```go
//output:
[-100 -5 30 100]
```

## SortBy

Returns sorted version of given array or slice.

> ***array, slice, map & struct are supported***
```go
gotil.SortBy(data)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/jmZZjaS6gqC) 

```go
data := []int64{100, 30, -100, -5}
newData, _ := gotil.Sort(data)
fmt.Println(newData)
```

```go
//output:
[-100 -5 30 100]
```