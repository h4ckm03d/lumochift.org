---
author : "Moch. Lutfi"
categories : ["Golang101"]
tags : ["programming", "golang", "golang101"]
date : 2020-08-30T18:00:00+07:00
excerpt : "Binary Operator Hack and Tricks"
type : "post"
authors : 
    - Moch Lutfi
title : "Binary Operator Hack and Tricks"
---

Binary operator is a way to manipulate binary data. We already know there are `&`, `|`, `^`, `<<` and `>>` operators, but not all of us know the secret of each operators. Let's explore what tricks behind those operator using go language.

## Multiply or Divide By 2

We already know how multiply 2 using `* 2` or divide using `/ 2`, but how we can achieve same with binary operator?

|          |       |             |
|----------|-------|------------|
| divide by 2 | shift right by 1|`someNumber >> 1` |
| multiply by 2 | shift left by 1 | `someNumber << 1` |

```go
    // multiply by 2
	fmt.Println(4 << 1)
    // Output: 8
    
    // divide by 2
    fmt.Println(4 >> 1)
	// Output: 2
```

## Change case of character

|          |       |             |
|----------|-------|-------------|
| change to uppercase | use `&` with `underscore` |`'c' & '_'` |
| change to lowercase | use `|` with `space`  |`'A' | ' '` |

```go

    // to upper case char
	fmt.Println((string)('c' & '_'))
	// Output: C

    // to lower case char
	fmt.Println(string('A' | ' '))
	// Output: a
```

## Invert case of character

Invert char can be achieved by `xor` with space

```go
	fmt.Println(string('A' ^ ' '), string('b' ^ ' '))
	// Output: a B
```

## Get letter position

Get letter's position in alphabet (1-26) using `and with 31`

```go
	fmt.Println('z' & 31)
	// Output: 26
```

## Check number odd or even

Simple check if number is odd/even using `and with 1`, odd number will return true

```go
    // odd number return true
    fmt.Println(7 & 1 > 0)
    // Output: true
    
    // even number return false
    fmt.Println(8 & 1 > 0)
	// Output: false

```

Try it yourself at https://play.golang.org/p/-wsIlDgBTmF