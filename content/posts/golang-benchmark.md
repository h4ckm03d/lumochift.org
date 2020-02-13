---
title: "Membuat Benchmark di Golang"
date: 2018-09-02T08:00:00+07:00
excerpt : "Bagaimana melihat performa kode menggunakan benchmark?"
hero: /images/hero-2.jpg
type : "post"
authors:
  - Moch Lutfi
categories : ["Golang101"]
tags : ["programming", "golang", "golang101", "benchmark", "Bahasa Indonesia"]
---

Ketika membuat suatu aplikasi tidak dipungkiri salah satu faktor yang sangat penting yaitu kecepatan. Baik kecepatan load data, kecepatan dalam menjalankan suatu perintah ataupun ketika membuka dan menutup aplikasi. Tentunya perlu instrument untuk mengetahui seberapa cepat kode kita, dalam hal ini perasaan tidak dapat digunakan sebagai tolak ukur. Menariknya pada bahasa pemrograman golang sudah ada *library* standar untuk mengukur seberapa cepat perintah dalam kode yang sering disebut *benchmarking*.

Kali ini contoh kasus untuk komparasi performa saya menggunakan 2 sorting sederhana yaitu *bubble sort* dan *shell sort*. Kira-kira mana yang lebih cepat ya? Ah iya, jangan pake perasaan tapi pake hasil *benchmark* untuk menentukan siapa yang paling cepat. Berikut contoh 2 sorting tersebut.

```go
// sort.go
package benchmark

import "math/rand"

// BubbleSort sorting array of integer using bubble sort
func BubbleSort(arr []int) []int {
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

// ShellSort sorting int using shell sort
func ShellSort(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}

// RandArray helper for create random array
func RandArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
```

Pada artikel sebelumnya tentang [Unit Test] sudah dibahas tentang bagaimana caranya membuat unit test pada suatu package, penggunaan *benchmark* juga tetap menggunakan *package* *testing* namun menggunakan variabel `B` bukan `T` seperti yg digunakan pada [Unit Test]. Langsung saja pada penggunaanya dalam kode berikut.

```go
package benchmark_test

import (
	"testing"

	"github.com/h4ckm03d/blog-codes/golang101/benchmark"
)

func BenchmarkBubbleSorting(b *testing.B) {
	arr := benchmark.RandArray(100)
	for n := 0; n < b.N; n++ {
		benchmark.BubbleSort(arr)
	}
}

func BenchmarkShellSorting(b *testing.B) {
	arr := benchmark.RandArray(100)
	for n := 0; n < b.N; n++ {
		benchmark.ShellSort(arr)
	}
}
```

Untuk menjalankan *benchmark* sama dengan unit test, hanya saja menggunakan parameter tambahan `-bench=.` untuk semua *benchmark*. Jika ingin menjalankan salah satu bisa menggunakan `-bench=ShellSort`, menggunakan nama fungsi *benchmark* tanpa menggunakan kata `Benchmark`. Berikut hasil *benchmark* dari 2 fungsi sorting diatas.

```bash
➜  benchmark git:(master) ✗ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/h4ckm03d/blog-codes/golang101/benchmark
BenchmarkBubbleSorting-12         300000              4181 ns/op
BenchmarkShellSorting-12         3000000               433 ns/op
PASS
ok      github.com/h4ckm03d/blog-codes/golang101/benchmark      3.049s
➜  benchmark git:(master) ✗ go test -bench=BubbleSort
goos: darwin
goarch: amd64
pkg: github.com/h4ckm03d/blog-codes/golang101/benchmark
BenchmarkBubbleSorting-12         300000              4188 ns/op
PASS
ok      github.com/h4ckm03d/blog-codes/golang101/benchmark      1.306s
➜  benchmark git:(master) ✗
```
Pada hasil perintah `go test -bench=.` diatas menghasilkan 3 kolom:

1. Nama benchmark, contohnya `BenchmarkBubbleSorting-12`

2. Total operasi yg dijalankan, `300000`

3. waktu yang dibutuhkan untuk menjalankan 1 fungsi dalam nanoseconds. `4181 ns/op`

Jadi `BubbleSort` perlu `4181 ns/op` dan `ShellSort` memerlukan `433 ns/op`. Sudah jelas kalau pemenangnya adalah `ShellShort`. Mudah bukan? 

Sampai jumpa lagi di tulisan selanjutnya.

Code lengkapnya ada di [github](https://github.com/h4ckm03d/blog-codes/tree/master/golang101/benchmark)


[Unit Test]: /posts/golang-unit-test/