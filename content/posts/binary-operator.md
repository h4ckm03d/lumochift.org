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
---

Binary operator is a way to manipulate binary data. We already know there are `&`, `|`, `^`, `<<` and `>>` operators, but not all of us know the secret of each operators. Let's explore what tricks behind those operator using go language.

## Multiply or Divide By 2

We already know how multiply 2 using `* 2` or divide using `/ 2`, but how we can achieve same with binary operator? 

|          |       |             |
|----------|-------|------------|
| divide by 2 | shift right by 1|`someNumber >> 1` |
| multiply by 2 | shift left by 1 | `someNumber << 1` |

## Change case of character

|          |       |             |
|----------|-------|-------------|
| change to uppercase | use `&` with `underscore` |`'c' & '_'` |
| change to lowercase | use `|` with `space`  |`'A' | ' '` |

## Invert case of character

## Get letter position

## Check number odd or even

