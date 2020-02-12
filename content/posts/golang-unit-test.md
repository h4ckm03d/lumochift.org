---
title: "Golang Unit Test"
date: 2018-01-02T08:00:00+07:00
description : "Bagaimana menggunakan Unit Test pada bahasa pemrograman go?"
hero: /images/hero-2.jpg
authors:
  - Moch Lutfi
categories : ["Golang101"]
tags : ["programming", "golang", "golang101", "unit test", "Bahasa Indonesia"]
type : "post"
---

*Unit test* merupakan salah satu cara untuk validasi sebuah unit terkecil dalam suatu aplikasi, misalnya *global variable*, fungsi, fungsi dalam suatu *class* (dalam *context* golang fungsi dalam `struct`). Adanya *unit test* ini juga mempermudah pengujian suatu aplikasi/*library* yang secara terus menerus/diperlukan repetisi pengujian sehingga tidak perlu membuat aplikasi yang lain untuk menggunakan *library* kemudian dicek satu-persatu secara manual output program sudah sesuai atau belum.

# Golang Unit Test

`GOLANG` sudah mempunyai *standard library* untuk *unit testing*, jadi tidak perlu menggunakan *third-party* untuk *unit test*. Meskipun demikian untuk kenyamanan bisa menggunakan *third-party tools* misalkan untuk penyederhanaan unit test menggunakan assert bisa menggunakan `github.com/stretchhr/testify/assert`.

Untuk demonstrasi penggunaan *unit test* dan cara penggunaannya, disini menggunakan 2 berkas yaitu `SimpleMath.go` dan untuk unit test menggunakan `SimpleMath_test.go`.
*Unit test* dalam golang diletakkan dalam package yang sama dan nama yg sama dengan nama berkas yang akan ditest dengan menambahkan `_test` pada nama berkas. Berikut adalah contoh sederhana penggunaan unit test.

```go
// math_test.go
package main

import(
    "testing"
)

func TestKotak(t *testing.T){
    // init kotak
    kotak := NewKotak(4, 8)

    if p.Luas() != 32{
        t.Error("Seharusnya 32")
    }
}

```

Untuk memastikan apakah *unit test* berjalan lancar bisa digunakan `go test` di `console`. Misalkan letak berkas di `$GOPATH/src/math`. Maka `~cd $GOPATH/src/math && go test`.

Contoh yang lebih realistis yaitu dengan membuat aplikasi untuk menghitung kotak dengan spesifikasi seperti berikut:

- Input merupakan variabel `panjang` dan `lebar`

- Menghitung luas dengan mengalikan `panjang` dengan `lebar`

- Menghitung keliling kotak dengan `2 * (panjang + lebar)`

- Mengecek apakah kotak merupakan persegi atau bukan dengan membandingkan panjang dengan lebar, jika sama merupakan persegi.

Spesifikasi sudah ditentukan sekarang saatnya mengubahnya menjadi code seperti dibawah ini.

```go
// SimpleMath.go
package main

type Kotak struct {
    Panjang int
    Lebar   int
}

func (p *Kotak) Luas() int {
    return p.Panjang * p.Lebar
}

func (p *Kotak) Keliling() int {
    return 2 * (p.Panjang + p.Lebar)
}

func (p *Kotak) IsPersegi() bool {
    if p.Panjang == p.Lebar {
        return true
    }
    return false
}

func NewKotak(p, l int) *Kotak {
    return &Kotak{p, l}
}
```

Sedangkan untuk *unit test* lengkapnya seperti dibawah ini. Untuk penamaan fungsi menggunakan prefix Test dan parameter argumen `t *testing.T`. Contoh di bawah ini  menggunakan *standard library* dan *third-party*.

```go
// SimpleMath_test.go
package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestKotak(t *testing.T) {
    p := NewKotak(4, 8)

    if p.IsPersegi() == true {
        t.Error("Seharusnya false")
    }

    if p.Luas() != 32 {
        t.Error("Seharusnya 32")
    }

    if p.Keliling() != 24 {
        t.Error("seharusnya 24")
    }

    persegi := NewKotak(4, 4)

    if persegi.IsPersegi() == false {
        t.Error("seharusnya true")
    }
}

func TestKotakPakeAssert(t *testing.T) {
    assert := assert.New(t)
    p := NewKotak(4, 8)

    assert.Equal(p.IsPersegi(), false, "seharusnya false")
    assert.Equal(p.Luas(), 32, "seharusnya 32")
    assert.Equal(p.Keliling(), 24, "seharusnya 24")

    persegi := NewKotak(4, 4)
    assert.Equal(persegi.Luas(), 16, "seharusnya 16")
    assert.Equal(persegi.Keliling(), 16, "seharusnya 16")
    assert.Equal(persegi.IsPersegi(), true, "seharusnya true")
}
```

Untuk menjalankan unit test bisa menggunakan perintah `go test` di *command line*, kalau misalkan perlu verbose bisa ditambahkan flag `-v` seperti ini `$ go test -v`.

Cukup mudah bukan penggunaan unit test di Golang? Selanjutnya mungkin akan membahas *test coverage* dan *benchmark* menggunakan golang. Sampai jumpa lagi di tulisan selanjutnya.

