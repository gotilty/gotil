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


## IsArray

Checks whether the given value is an array. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsArray(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/HUoMqpHPFbd)

```go
var a = []string{"gotil is perfect!"}
fmt.Println(gotil.IsArray(a))
```

```go
//output:
false
```

## IsSlice

Checks whether the given value is a slice. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsSlice(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/6M2sSgrv99n)

```go
var a = []string{"gotil is perfect!"}
fmt.Println(gotil.IsSlice(a))
```

```go
//output:
true
```


## IsString

Checks whether the given value is a string. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsString(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/Q0TTjR-4tKO)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"

fmt.Println(gotil.IsString(a))
fmt.Println(gotil.IsString(b))
```

```go
//output:
false
true
```

## IsBool

Checks whether the given value is a bool. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsBool(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/62rFwwFs_LW)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false

fmt.Println(gotil.IsBool(a))
fmt.Println(gotil.IsBool(b))
fmt.Println(gotil.IsBool(c))
```

```go
//output:
false
false
true
```

## IsInteger

Checks whether the given value is an integer. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsInteger(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/9IFrmupEr8q)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = 64

fmt.Println(gotil.IsInteger(a))
fmt.Println(gotil.IsInteger(b))
fmt.Println(gotil.IsInteger(c))
fmt.Println(gotil.IsInteger(d))
```

```go
//output:
false
false
false
true
```

## IsFloat

Checks whether the given value is a float. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsFloat(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/fjisF7eWHt9)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = 64.5

fmt.Println(gotil.IsFloat(a))
fmt.Println(gotil.IsFloat(b))
fmt.Println(gotil.IsFloat(c))
fmt.Println(gotil.IsFloat(d))
```

```go
//output:
false
false
false
true
```

## IsNumber

Checks whether the given value is a number. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsNumber(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/CJxGVWiHFka)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = 64.5

fmt.Println(gotil.IsNumber(a))
fmt.Println(gotil.IsNumber(b))
fmt.Println(gotil.IsNumber(c))
fmt.Println(gotil.IsNumber(d))
```

```go
//output:
false
false
false
true
```


## IsFunc

Checks whether the given value is a function. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsFunc(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/7Cr23xAywC1)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = func(a, b int) {}

fmt.Println(gotil.IsFunc(a))
fmt.Println(gotil.IsFunc(b))
fmt.Println(gotil.IsFunc(c))
fmt.Println(gotil.IsFunc(d))
```

```go
//output:
false
false
false
true
```

## IsStruct

Checks whether the given value is a struct. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsStruct(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ufHq4CRiDm8)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = tools{
    name: "gotil",
}

fmt.Println(gotil.IsStruct(a))
fmt.Println(gotil.IsStruct(b))
fmt.Println(gotil.IsStruct(c))
fmt.Println(gotil.IsStruct(d))
```

```go
//output:
false
false
false
true
```


## IsPointer

Checks whether the given value is a pointer. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsPointer(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ilhhPAixO_t)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = &tools{
    name: "gotil",
}

fmt.Println(gotil.IsPointer(a))
fmt.Println(gotil.IsPointer(b))
fmt.Println(gotil.IsPointer(c))
fmt.Println(gotil.IsPointer(d))
```

```go
//output:
false
false
false
true
```

## IsChan

Checks whether the given value is a channel. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsChan(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/mMOZfEkvHYJ)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = make(chan int)

fmt.Println(gotil.IsChan(a))
fmt.Println(gotil.IsChan(b))
fmt.Println(gotil.IsChan(c))
fmt.Println(gotil.IsChan(d))
```

```go
//output:
false
false
false
true
```

## IsMap

Checks whether the given value is a map. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsMap(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/vzo56qsQ3vI)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = map[string][]int{}

fmt.Println(gotil.IsMap(a))
fmt.Println(gotil.IsMap(b))
fmt.Println(gotil.IsMap(c))
fmt.Println(gotil.IsMap(d))
```

```go
//output:
false
false
false
true
```


## IsByteArray

Checks whether the given value is a byte array. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsByteArray(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/ma4pU4caYXs)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"
var c = false
var d = []byte{}

fmt.Println(gotil.IsByteArray(a))
fmt.Println(gotil.IsByteArray(b))
fmt.Println(gotil.IsByteArray(c))
fmt.Println(gotil.IsByteArray(d))
```

```go
//output:
false
false
false
true
```


## IsEqual

Checks whether the given parameters have equal value. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsEqual(val1, val2)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/0BWX3yW6FNb)

```go
var a = []string{"gotil is perfect!"}
var b = "gotil is perfect!"

var c = "1234"
var d = 1234

fmt.Println(gotil.IsEqual(a, b))
fmt.Println(gotil.IsEqual(c, d))
```

```go
//output:
false
false
```


## IsEmpty

Checks whether the given parameters have equal value. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsEmpty(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/-YUojLxKEuJ)

```go
var a = []string{"gotil is perfect!"}
var c = ""

fmt.Println(gotil.IsEmpty(a))
fmt.Println(gotil.IsEmpty(c))
```

```go
//output:
false
true
```


## IsNil

Checks whether the given parameters have equal value. Returns boolean.

> ***all primitive types, array, slice, map & struct are supported.***

```go
gotil.IsNil(val)
```

### examples

>💻 [Try on Playground](https://go.dev/play/p/unzv1GqwpOK)

```go
var a = []string{"gotil is perfect!"}
var c = ""

fmt.Println(gotil.IsNil(a))
fmt.Println(gotil.IsNil(c))
```

```go
//output:
false
true
```