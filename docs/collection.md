## Filter

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> **_array & slice are supported_**

```go
gotil.Filter(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w)

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

> **_array & slice are supported_**

```go
gotil.Filter(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w)

```go
data := []int64{-100, -5, 30, 100}
// Input: [-100 -5 30 100]
result, _ := gotil.FindBy(data, func(val interface{}, i int) bool {
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
30
```

## FindByFromIndex

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> **_array & slice are supported_**

```go
gotil.FindByFromIndex(val, f, i)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/A98eBoD6Nxi)

```go
data := []int64{-100, -5, 30, 100}
// Input: [-100 -5 30 100]
result, _ := gotil.FindByFromIndex(data, func(val interface{}, i int) bool {
    if val.(int64) > 0 {
        return true
    } else {
        return false
    }
}, 3)
fmt.Println(result)
```

```go
//output:
100
```

## FindLastBy

Inversely checks whether the predicate is true for all elements of the collection. Iteration is stopped once predicate returns false.

> **_array & slice are supported_**

```go
gotil.FindLastBy(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/ay1huuG2ipx)

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

> **_array & slice are supported_**

```go
gotil.GroupBy(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/ay1huuG2ipx)

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

It checks whether the given value exists in the given parameter. Returns boolean.

> **_array & slice are supported_**

```go
gotil.Includes(array, val)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd)

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

// Reduce iterates given collection and returns the accumulated result of running each element the last param is initial value of accumulator.

> **_array & slice are supported_**

```go
data := []int{5, 10}
result, _ := gotil.Reduce(data, func(accumulator, val interface{}, i int) interface{} {
    return accumulator.(int) + val.(int)
}, 0)
fmt.Println(result)
// Output: 15
```

### examples TODO

> 💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd)

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

> **_array & slice are supported_**

```go
gotil.Shuffle(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/jxNSYukWBZR)

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

> **_array & slice are supported_**

```go
gotil.ShuffleSeed(data, seed)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/IeA4GbdIqjd)

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

> **_array & slice are supported_**

```go
gotil.Size(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/EY4ca2fMJhR)

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

> **_array, slice, map & struct are supported_**

```go
gotil.Sort(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/jmZZjaS6gqC)

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

> **_array, slice, map & struct are supported_**

```go
gotil.SortBy(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/8BrQFoMJAVm)

```go
var data = []user{
    {
        name: "Micheal",
        age:  27,
        location: location{
            city: "New York",
        },
    },
    {
        name: "Joe",
        age:  30,
        location: location{
            city: "Detroit",
        },
    },
    {
        name: "Olivia",
        age:  42,
        location: location{
            city: "New York",
        },
    },
    {
        name: "Kevin",
        age:  10,
        location: location{
            city: "Boston",
        },
    },
}
newData, _ := gotil.SortBy(data, "location.city")
fmt.Println(newData)
```

```go
//output:
[{Kevin 10 {Boston}} {Joe 30 {Detroit}} {Olivia 42 {New York}} {Micheal 27 {New York}}]
```
