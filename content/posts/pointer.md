---
title: "Pointer"
date: 2017-11-05T15:36:49+07:00
type : "post"
description : "Pointer explanation in golang"
author : "Moch. Lutfi"
categories : ["Golang101"]
tags : ["programming", "golang", "golang101"]
draft: false
toc : false
---

Pointers gave me some nostalgic moment when learning C++ on first semester. They become source of headache on my day. Pointers are one of main tools to achieve high performance code in non-garbage-collected languages. So what about pointers in Go? Luckily Go's pointer have achieved the best of both worlds by providing high-performance pointers with garbage-collector capabilities and easiness.

# Pointers in real world

Unconsiously we are already knew main concept of pointers in real world. Nowadays online markets are popular, we'd like to shopping using smartphone or computer. After checkout if you want to receive package in your house, it's far easier to simply send the address of your house(pointer) instead of sending the entire house to the seller so that your package is deposited inside. The problem is, if you send wrong address of yours. You will not get the package.

Let's back to programming world. For example I have simple program to convert mp3 files to other format. The first step I must read the mp3 files and save to the variable, let's say the variable size is 1GB and I need pass it to another function. Without a pointer, the entire variable is cloned to the scope of the function that is going to use it. So I will have 2GB of memory occupied by using this variable twice. If the second function will pass again to another function, the memory occupied will raised.

If I use a pointer to pass a very small reference to the another function so that just the small reference is cloned and I can keep memory usage low.

Different with C or C++ pointers, in GO very limited. We can't use pointer arithmetic nor can create a pointer to reference an exact posision in the stack.

Here the basic of [Example] of pointers:

```go
package main

import (
	"fmt"
	"unsafe"
)

type Music struct {
	SongName string
	TrackNo  int
	Singer   string
}

func main() {
    // declare a variable
    music := Music{"November Rain", 20, "Gun n Roses"}
    // set pointer of music to variable p
	p := &music
	fmt.Printf("music data %v\n", music)
	fmt.Printf("Pointer of music %p\n", p)
	// access original value using pointer
	fmt.Printf("Get music from pointer %v\n", *p)

	fmt.Printf("Original size: %T, %d\n", music, unsafe.Sizeof(music))
	fmt.Printf("Pointer size: %T, %d\n", p, unsafe.Sizeof(p))
}

// output:
// music data {November Rain 20 Gun n Roses}
// Pointer of music 0x10444240
// Get music from pointer {November Rain 20 Gun n Roses}
// Original size: main.Music, 20
// Pointer size: *main.Music, 4
```

`music := Music{"November Rain", 20, "Gun n Roses"}` code represent our 1GB variable and `p` contains the reference with value `0x10444240` (represented by an ampersand). We can use asterisk `*` to take value from referenced by the pointer. With those example we have origical size of music data is 20 bytes and pointer only 4 bytes. Pointer size have same size with int variable in Go. Even if I increase the size of struct/variable the pointer remain same size.

I hope you can understand. See ya on the next [Golang 101] post...
[Example]: https://play.golang.org/p/FqVdZ6ntLN
[Golang 101]: /categories/golang101/