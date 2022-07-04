## Every

Checks if the predicate returns true for all of the collection's elements. When the predicate returns false, the
iteration comes to an end.

> ***array, slice, map & struct are supported***

```go
gotil.EveryBy(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/4ALoPEWqMug)

```go
input := []user{
	{
		name: "Micheal",
		age:  27,
	},
	{
		name: "Joe",
		age:  30,
	},
	{
		name: "Olivia",
		age:  42,
	},
	{
		name: "Mike",
		age:  43,
	},
}
result := gotil.EveryBy(input, func(v user) bool {
	return v.age < 30
})
fmt.Println(result)

```

```go
//output:
false
```

## Each

Iterates an array of values by running each element in the given array thru iteratee.

> ***array, slice, map & struct are supported***

```go
gotil.Each(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/5JsdmpiqQIi)

```go
gotil.Each([]string{"gotilty", "gotil"}, func(v string) {
	fmt.Fprint(os.Stdout, v)
})
```

```go
//output:
gotiltygotil
```

## Filter

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> **_array & slice are supported_**

```go
gotil.Filter(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/aLVfuwExKGO)

```go
result := gotil.FilterBy([]int64{-100, -5, 30, 100}, func(v int64, i int) bool {
	return v > 0
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
gotil.FindBy(val, f()bool)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/9CHhvvipA8w)

```go
input2 := []user{
	{
		name: "Micheal",
		age:  27,
	},
	{
		name: "Joe",
		age:  30,
	},
	{
		name: "Olivia",
		age:  42,
	},
}

result := gotil.FindBy(input2, func(val user) bool {
	return val.name == "Olivia"
})
fmt.Println(result)
```

```go
//output:
{Olivia 42}
```

## FindByFromIndex

Checks if predicate returns true for all elements of collection. Iteration is stopped once predicate returns false.

> **_array & slice are supported_**

```go
gotil.FindByFromIndex(val, f, i)
```

### examples
> 💻 [Try on Playground](https://go.dev/play/p/vnnIiD_Luce)

```go
data := []int64{-100, -5, 30, 100}
// Input: [-100 -5 30 100]
result := gotil.FindByFromIndex(data, func(val int64) bool {
return val > 0
}, 3)
fmt.Println(result)
```

```go
//output:
100
```

## FindLastBy

Inversely checks whether the predicate is true for all elements of the collection. Iteration is stopped once predicate
returns false.

> **_array & slice are supported_**

```go
gotil.FindLastBy(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/_Td7R7uEyyP)

```go
data := []int64{-100, -5, 30, 100}
// Input: [-100 -5 30 100]
newData := gotil.FindLastBy(data, func(val int64) bool {
	return val == 30
})
	fmt.Println(newData)
```

```go
//output:
30
```

## GroupBy

Creates a map composed of keys generated from the results of running each element of collection through the iteratee
function.

> **_array & slice are supported_**

```go
gotil.GroupBy(val, f)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/8qQclTSuIf8)

```go
input := []int{10, 20, 30}
result := gotil.GroupBy(input, func(v int, i int) int {
	return v / 10
})
fmt.Println(result)
```

```go
//output:
map[1:[10] 2:[20] 3:[30]]
```

## Reduce

// Reduce iterates given collection and returns the accumulated result of running each element the last param is initial
value of accumulator.

> **_array & slice are supported_**

### examples

> 💻 [Try on Playground](https://go.dev/play/p/8Nl5N9y9_kF)

```go
result := gotil.Reduce([]int{5, 10}, func(accumulator, v, i int) int {
	return accumulator + v
}, 0)
fmt.Println(result)
```

```go
//output:
15
```


## Shuffle

Shuffles the list randomly. A new array is returned.

> **_array & slice are supported_**

```go
gotil.Shuffle(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/XWtS7si2i3i)

```go
data := []int64{-100, -5, 30, 100}
newData := gotil.Shuffle(data)
fmt.Println(newData)
```

```go
//output:
[100 30 -5 -100]
```

## ShuffleSeed

Shuffles the list based on the given seed. A new array is returned.

> **_array & slice are supported_**

```go
gotil.ShuffleSeed(data, seed)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/jnVTdRE9Dsc)

```go
// Use "seed := time.Now().UnixNano()" for random result
seed := int64(58239238)
//Seed you will get the same sequence of pseudo­random numbers
// each time you run the program.

data := []int64{-100, -5, 30, 100}
newData := gotil.ShuffleSeed(data, seed)
fmt.Println(newData)
```

```go
//output:
[-5 100 -100 30]
```


## Sort

Returns sorted version of given array or slice.

> **_array, slice, map & struct are supported_**

```go
gotil.Sort(data)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/LuSbBJFiG8F)

```go
data := []int64{-100, -5, 30, 100}
newData := gotil.SortDesc(data)
fmt.Println(newData)
```

```go
//output:
[100 30 -5 -100]
```

## SortBy

Returns sorted version of given array or slice.

> **_array, slice, map & struct are supported_**

```go
gotil.SortBy(data, path)
```

### examples

> 💻 [Try on Playground](https://go.dev/play/p/6CthKZcvLk7)

```go
input := []user{
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
result := gotil.SortBy(input, "age")
fmt.Println(result)
```

```go
//output:
[{Kevin 10 {Boston}} {Micheal 27 {New York}} {Joe 30 {Detroit}} {Olivia 42 {New York}}]
```
