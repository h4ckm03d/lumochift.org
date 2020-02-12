---
title: "Tips dan trik unit test di Go"
date: 2018-02-01T21:54:00+07:00
description : "Tips dan trik menggunakan unit test di GO"
hero: /images/hero-2.jpg
authors:
  - Moch Lutfi
categories : ["Golang101"]
tags : ["programming", "golang", "golang101", "unit test", "Bahasa Indonesia", "tips"]
draft: false
toc : true
type : "post"
---
[Unit Test] memang tidak bisa dilepaskan dari proses pengembangan *software*. Namun seringkali dalam pembuatan [Unit Test] di Go terjadi banyak repetisi yang tidak perlu dan [Unit Test] yang tidak dikelola dengan baik. Salah satu contoh kasus yang paling banyak ditemui dalam pembuatan [Unit Test] yaitu tidak dipisahkannya *logic* dan data sehingga ketika penambahan data test terdapat penambahan pula *logic*.

# Contoh Kode

Sebelum memulai lebih lanjut berikut adalah contoh kode untuk membantu memahami tulisan ini.

- Struktur direktori

```bash
$ tree .
├── add.go
└── add_test.go
```

- File `add.go`

```go
package main

// Tambah merupakan fungsi sederhana penjumlahan
func Tambah(a, b int)int{
    return a+b
}
```

- File `add_test.go`

```go
package main

import "testing"

func TestBasic(t *testing.T){
    if Tambah(1,1) != 2{
        t.Error("seharusnya 2")
    }

    if Tambah(1,2) != 3{
        t.Error("seharusnya 3")
    }

    // kode berulang sebanyak jumlah data test
}

```

- Hasil running [Unit Test]

```bash
$ go test -v
=== RUN   TestBasic
--- PASS: TestBasic (0.00s)
PASS
ok      github.com/h4ckm03d/blog-codes/golang101/6-tips-trik-unit-test  0.002s
```

# Memanfaatkan *array of struct*

Contoh [Unit Test] diatas merupakan contoh paling sederhana, jika yang sederhana saja sudah terlalu verbose dan susah dikelola bagaimana jika kode dalam suatu `package` semakin besar? Salah satu cara untuk memisahkan *logic* unit test dengan data yaitu menggunakan unit test *array of struct* untuk data dan tinggal menambahkan looping data ketika pengecekan. Tambahkan fungsi dibawah ini ke dalam `add_test.go`

```go
func TestB(t *testing.T) {
    testData := []struct {
        name   string
        inputA int
        inputB int
        result int
    }{
        {"1+1", 1, 1, 2},
        {"1+2", 1, 2, 3},
        {"1+3", 1, 3, 4},
    }

    for _, tc := range testData {
        t.Run(tc.name, func(t *testing.T) {
            if Tambah(tc.inputA, tc.inputB) != tc.result {
                t.Errorf("Seharusnya %d", tc.result)
            }
        })
    }
}

```

Jika [Unit Test] dijalankan maka hasilnya seperti berikut:

```bash
$ go test -v
=== RUN   TestBasic
--- PASS: TestBasic (0.00s)
=== RUN   TestB
=== RUN   TestB/1+1
=== RUN   TestB/1+2
=== RUN   TestB/1+3
--- PASS: TestB (0.00s)
    --- PASS: TestB/1+1 (0.00s)
    --- PASS: TestB/1+2 (0.00s)
    --- PASS: TestB/1+3 (0.00s)
PASS
ok      github.com/h4ckm03d/blog-codes/golang101/6-tips-trik-unit-test  0.002s
```

Dalam `TestB` diatas ketika data test bertambah, kita hanya perlu menambahkan kedalam `array` `testData` tanpa perlu merubah *logic* sehingga kode lebih mudah dibaca dan dimodifikasi. `struct` dalam peubah `testData` terdiri dari 3 bagian label test(`name`), input data (`inputA`, `inputB`), dan output dari fungsi (`result`). Untuk input data dan output bisa diganti sesuai dengan fungsi yang ditest. Hasil running dari unit test juga lebih jelas jika menggunakan metode ini karena penggunaan label membantu tracking ketika ada kesalahan. Misalkan `testData` 1 ditambah 1 diubah menjadi 7 untuk mengetahui contoh jika unit test menghasilkan error, hasilnya sebagai berikut:

```bash
$ go test -v
=== RUN   TestBasic
--- PASS: TestBasic (0.00s)
=== RUN   TestB
=== RUN   TestB/1+1
=== RUN   TestB/1+2
=== RUN   TestB/1+3
--- FAIL: TestB (0.00s)
    --- FAIL: TestB/1+1 (0.00s)
        add_test.go:34: Seharusnya 7
    --- PASS: TestB/1+2 (0.00s)
    --- PASS: TestB/1+3 (0.00s)
FAIL
exit status 1
FAIL    github.com/h4ckm03d/blog-codes/golang101/6-tips-trik-unit-test  0.005s
```

Karena output diubah menjadi 7 sehingga unit testnya fail, akan tetapi karena menggunakan label kita segera mengetahui data apa yang salah sehingga bug fix lebih mudah.

# Fatal VS Error

Pada fungsi `TestBasic` diatas menggunakan `t.Error("seharusnya 2")`, kenapa menggunakan `t.Error` bukan menggunakan `t.Fatal` ? Penggunaan *error* diatas karena ketika terjadi kesalahan misalkan output salah atau tidak sesuai maka unit test tidak akan berhenti tetapi melanjutkan sampai semua proses dalam fungsi tersebut selesai. Jika menggunakan *fatal* maka ketika `t.Fatal` atau `t.Fatalf` dipanggil maka proses dihentikan. Dalam hal ini ketika 1 + 1 terjadi error maka proses berhenti dan tidak dilanjutkan pengecekan 1 + 2. Berikut contoh perbedaan penggunaan *fatal* dan *error*.

```go
func TestBasicFatal(t *testing.T) {
    if Tambah(1, 1) != 3 {
        t.Error("seharusnya 2")
    }

    if Tambah(1, 2) != 4 {
        t.Error("seharusnya 3")
    }

    // kode berulang sebanyak jumlah data test
}

func TestBasicError(t *testing.T) {
    if Tambah(1, 1) != 3 {
        t.Error("seharusnya 2")
    }

    if Tambah(1, 2) != 4 {
        t.Error("seharusnya 3")
    }

    // kode berulang sebanyak jumlah data test
}
```
Hasil `go test -v` sebagai berikut

```bash
$ go test -v
=== RUN   TestBasic
--- PASS: TestBasic (0.00s)
=== RUN   TestBasicFatal
--- FAIL: TestBasicFatal (0.00s)
        add_test.go:21: seharusnya 2
=== RUN   TestBasicError
--- FAIL: TestBasicError (0.00s)
        add_test.go:33: seharusnya 2
        add_test.go:37: seharusnya 3
=== RUN   TestB
=== RUN   TestB/1+1
=== RUN   TestB/1+2
=== RUN   TestB/1+3
--- PASS: TestB (0.00s)
    --- PASS: TestB/1+1 (0.00s)
    --- PASS: TestB/1+2 (0.00s)
    --- PASS: TestB/1+3 (0.00s)
FAIL
exit status 1
FAIL    github.com/h4ckm03d/blog-codes/golang101/6-tips-trik-unit-test  0.004s
```

Ketika `TestBasicFatal` dijalankan maka ketika ada error proses berhenti berbeda dengan `TestBasicError`, meskipun ada error proses dilanjutkan sampai selesai. 

Dengan pemisahan kode [Unit Test] dan penggunaan error/fatal yang tepat kita bisa lebih optimal memanfaatkan unit test. Semoga bermanfaat dan sampai jumpa lagi di tulisan selanjutnya.


[Unit Test]: /posts/golang-unit-test/