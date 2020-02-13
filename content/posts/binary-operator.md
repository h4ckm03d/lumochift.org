---
author : "Moch. Lutfi"
categories : ["Golang101"]
tags : ["programming", "golang", "golang101"]
date : 2020-02-10T18:00:00+07:00
excerpt : "Binary Operator Hack and Tricks"
type : "post"
authors : 
    - Moch Lutfi
title : "Binary Operator Hack and Tricks"
draft: true
toc : true
---

Binary operator is a way to manipulate binary data. We already know there are `&`, `|`, `^`, `<<` and `>>` operators, but not all of us know the secret of each operators. Let's explore what tricks behind those operator.

## Multiply or Divide By 2

divide by 2 : shift right 1 >> 1 
multiply by two: with shift left 1 << 1

## Change case of character

	"c to uppercase ": {'c' & '_', 'C'},
		"A to lowercase ": {'A' | ' ', 'a'},

## Invert case of character

## Get letter position

## Check number odd or even

