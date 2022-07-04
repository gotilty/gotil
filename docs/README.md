#

## ⚙️ Installation

Make sure you have Go installed ([download](https://go.dev/dl/)). Version `1.18` or higher is required.

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://go.dev/blog/using-go-modules)) inside the folder. Then install Gotil with the [`go get`](https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/gotilty/gotil
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