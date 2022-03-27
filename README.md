# gotil
 utility library for go 

 ![gotil-logo](images/gotil.png)

## Table of Contents

- [gotil](#gotil)
  - [Table of Contents](#table-of-contents)
  - [Description](#description)
  - [Features](#features)
  - [Run on Local](#run-on-local)
    - [Step 1: git clone this repository gotilty/gotil@github](#step-1-git-clone-this-repository-gotiltygotilgithub)
    - [Step 2: download dependencies](#step-2-download-dependencies)
    - [Step 3: run test](#step-3-run-test)
    - [Step 3: run benchmarks](#step-3-run-benchmarks)
  - [TODO](#todo)

---


## Description

utility library for go 

## Features

* No thirdparty libraries.

## Run on Local 

### Step 1: git clone this repository [gotilty/gotil@github](https://github.com/gotilty/gotil)
```
git clone https://github.com/gotilty/gotil.git
```
Go to project folder
```
cd gotil
```
### Step 2: download dependencies
```
go mod download
```
### Step 3: run test
```
go test -v ./... 
```

### Step 3: run benchmarks
```
go test -bench . -v
```
```
└─(02:55:58 on feature/readme ✹ ✭)──> go test -bench . -v                                                                                                                       ──(Sun,Mar27)─┘
=== RUN   TestIsAssigned
--- PASS: TestIsAssigned (0.00s)
goos: darwin
goarch: arm64
pkg: github.com/gotility/gotil
BenchmarkIsAssignedString
BenchmarkIsAssignedString-8     107678102               10.78 ns/op
BenchmarkIsAssignedInteger
BenchmarkIsAssignedInteger-8    235145172                5.076 ns/op
BenchmarkIsAssignedStruct
BenchmarkIsAssignedStruct-8      9016149               131.4 ns/op
PASS
ok      github.com/gotility/gotil       5.155s
```
## TODO

- [x] Readme File
- [ ] Go Doc
- [x] ApiDoc
- [x] Tests
