# gotil

[methods](https://github.com/gotilty/gotil/wiki/methods) | [wiki](https://github.com/gotilty/gotil/wiki) | [documentation](https://gotilty.github.io/gotil/#/) 



lodash inspired utility library for go



<img src="images/gotil.png" width="250" height="100" style="display: block; margin: 0 auto">


------------


[![Build Status](https://scrutinizer-ci.com/g/gotilty/gotil/badges/build.png?b=master)](https://scrutinizer-ci.com/g/gotilty/gotil/build-status/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gotilty/gotil)](https://goreportcard.com/report/github.com/gotilty/gotil)
![Scrutinizer code quality (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/quality/g/gotilty/gotil/master)
![GitHub](https://img.shields.io/github/license/gotilty/gotil)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gotilty/gotil)

------------
## ⚙️ Installation

Make sure you have Go installed ([download](https://go.dev/dl/)). Version `1.18` or higher is required.

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://go.dev/blog/using-go-modules)) inside the folder. Then install Gotil with the [`go get`](https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/gotilty/gotil
```

## Run on Local

### Step 1: git clone this repository [gotilty/gotil@github](https://github.com/gotilty/gotil)

```
git clone https://github.com/gotilty/gotil.git
```

Go to project folder

```
cd gotil
```

### Step 2: run test

```
go test -v ./...
```

### Step 3: run benchmarks

```
go test -v ./... -bench=. -run=xxx -benchmem
```

## 🎯 Why GOtil?

GOtil makes go easier by taking the hassle out of working with arrays,
numbers, objects, strings, etc. GOtil's modular methods are great for:

- Iterating arrays, objects, & strings
- Manipulating & testing values
- Creating composite functions

## 👀 Examples

Listed below are some of the common examples. If you want to see more code examples , please visit our [Recipes repository](https://github.com/gotilty/gotil) or visit our hosted [API documentation](https://gotilty.github.io/gotil/#/).

#### 📖 [**Each**]

```go
gotil.Each([]string{"gotilty", "gotil"}, func(v string) {
	fmt.Fprint(os.Stdout, v)
})
// Output: gotiltygotil
```



## TODO

- [x] Readme File
- [x] Go Doc
- [x] ApiDoc
- [x] Tests
