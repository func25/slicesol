# slicesol

slicesol is a set of utils which adapts to generic go 1.18, makes slice easier to utilize

- [Getting Started](#getting-started)
  * [Installation](#installation)
  * [Samples](#samples)
- [Functions](#functions)
  * [Built-in functions](#built-in-functions)
  * [Using as PIPE](#using-as-pipe)

## Installation

go get github.com/func25/slicesol

## Samples

Insert at index 1 with 2 elements 2 & 3
```go
arr := slicesol.Insert([]int{1, 4, 5}, 1, 2, 3)
// result: [1 2 3 4 5]
```

Remove element at index 0 of int slice
``` go
arr := slicesol.Remove([]int{1, 2, 3, 4, 5}, 0)
// result: [2 3 4 5]
```

Stimulate queue/ stack structure

```go
popArr := &[]int{1, 2, 3, 4, 5}
p, err := slicesol.Pop(arr)
// p: 5
// popArr: [1, 2, 3, 4]

deqArr := &[]int{1, 2, 3, 4, 5}
d, err := slicesol.Dequeue(deqArr)
// d: 1
// deqArr: [2, 3, 4, 5]
```

Or you can pipe the slice with Sliol
```go
arr := slicesol.Sliol[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
newSlice := arr.
            Filter(func(e int) bool { return e%2 == 0 }).
            Map(func(e int) int { return e * 2 })
// newSlice: [4, 8, 12, 16, 20]
```

## Built-in functions
  - `Any`: returns index of element that meet fn condition, or -1 if not
  - `IsAny`: returns true if any element meets fn condition, or -1 if not
  - `Duplicate`: creates a new slice base on input slice
  - `Insert`: inserts elems to slice at i (index)
  - `Remove`: will CHANGE the original slice but remain its element order
  - `RemoveUnor`: will CHANGE the original slice and its element order also
  - `Filter`: creates new slice contains elements that meet fn condition
  - `Map`: transforms each element of input slice through fn and return a new transformed slice
  - `Shuffle`: will return a NEW slice as shuffled version of input slice
  - ... 

## Using as PIPE

This package has a type called Sliol supporting creating function pipes 
```go
arr := slicesol.Sliol[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
newSlice := arr.
            Filter(func(e int) bool { return e%2 == 0 }).
            Map(func(e int) int { return e * 2 })
// newSlice: [4, 8, 12, 16, 20]
```
