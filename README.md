# go-bytes
![build status](https://github.com/j-dumbell/go-bytes/actions/workflows/build_test.yml/badge.svg)

A Golang package for working with quantities of bytes, inspired by `time.Duration`. It provides functionality for 
ergonomically representing and converting quantities of bytes in the range -2^63 bytes to 2^63 bytes.

## Installation
```shell
go get github.com/j-dumbell/go-bytes
```

## Usage
Initialize a `Bytes` instance using the base 10 or base 2 unit constants, or from an existing integer or float
value:
```go
a := 100 * Megabyte     // 100 MB
b := 3 * GibiByte       // 3 GiB
c := Bytes(1_600_000)   // 1.6 MB
```

Convert a `Bytes` instance to the desired units of bytes:
```go
a := (10 * MibiByte).Mebibytes() // 10 
b := (3 * GigaByte).Terabytes() // 0.003
```
