# gotil

[methods](https://github.com/gotilty/gotil/wiki/methods) | [wiki](https://github.com/gotilty/gotil/wiki)

 utility library for go 


[![CodeFactor](https://www.codefactor.io/repository/github/gotilty/gotil/badge)](https://www.codefactor.io/repository/github/gotilty/gotil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gotilty/gotil)](https://goreportcard.com/report/github.com/gotilty/gotil)
![SymfonyInsight Grade](https://img.shields.io/symfony/i/grade/3a820096-5934-4bac-ab0e-a55e0a78141e)
![Scrutinizer code quality (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/quality/g/gotilty/gotil/master)
![GitHub](https://img.shields.io/github/license/gotilty/gotil)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gotilty/gotil)

<img src="images/gotil.png" width="250" height="100">


## Table of Contents

- [gotil](#gotil)
  - [Table of Contents](#table-of-contents)
  - [Description](#description)
  - [Features](#features)
  - [Run on Local](#run-on-local)
    - [Step 1: git clone this repository gotilty/gotil@github](#step-1-git-clone-this-repository-gotiltygotilgithub)
    - [Step 2: run test](#step-2-run-test)
    - [Step 3: run benchmarks](#step-3-run-benchmarks)
  - [TODO](#todo)

---


## Description

utility library for go 

## Features

* No thirdparty libraries.
* all features [wiki](https://github.com/gotilty/gotil/wiki/methods)

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
```
└─(02:55:58 on feature/readme ✹ ✭)──> go test -v ./... -bench=. -run=xxx -benchmem -v                                                                                                                       ──(Sun,Mar27)─┘
=== RUN   TestIsAssigned
--- PASS: TestIsAssigned (0.00s)
goos: darwin
goarch: arm64
pkg: github.com/gotilty/gotil
BenchmarkIsAssignedString
BenchmarkIsAssignedString-8     107678102               10.78 ns/op
BenchmarkIsAssignedInteger
BenchmarkIsAssignedInteger-8    235145172                5.076 ns/op
BenchmarkIsAssignedStruct
BenchmarkIsAssignedStruct-8      9016149               131.4 ns/op
PASS
ok      github.com/gotilty/gotil       5.155s
```
## TODO

- [x] Readme File
- [ ] Go Doc
- [x] ApiDoc
- [x] Tests
