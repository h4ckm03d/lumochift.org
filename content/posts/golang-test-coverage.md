---
title: "Golang Test Coverage"
date: 2018-01-03T20:10:36+07:00
type : "post"
description : "Golang Test Coverage"
hero: /images/hero-2.jpg
authors:
  - Moch. Lutfi
categories : ["Golang101"]
tags : ["programming", "golang", "golang101"]
draft: false
toc : true
---

Tulisan ini adalah lanjutan dari [Golang Unit Test]. Kali ini membahas tentang *test coverage*, yaitu untuk mengetahui apakah unit test sudah memenuhi semua jalur logika dari sistem yang kita uji. Contoh sederhana sebagai berikut:

1. Suatu fungsi untuk menentukan nilai maksimal dari dua integer input.

2. Input didefinisikan dalam variabel `a` dan `b`.

3. Jika `a` lebih besar atau sama dengan `b` maka return `a`, sebaliknya jika `b` lebih besar maka return `b`.

Sample code sebagai berikut:

```go
package main

func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
```

Kemudian *unit test* sebagai berikut:

```go
package main

import "testing"

func TestMax(t *testing.T) {
	if Max(1, 3) != 3 {
		t.Error("error, seharusnya 3")
	}

    // Jika pengecekan dibawah ini dihilangkan maka test coverage jadi 66.67%
	if Max(4, 3) != 4 {
		t.Error("error, seharusnya 4")
	}
}
```

Jika *test coverage* dijalankan menggunakan perintah `go test --cover` di cli maka hasilnya sebagai berikut:

```bash
$ go test --cover
PASS
coverage: 100.0% of statements
ok      github.com/h4ckm03d/blog-codes/test-coverage    0.002s
``` 

Coba kita review *unit testing* diatas.

1. Fungsi Max mempunya 2 cabang logika dan terdiri dari 3 statement `if a >= b`, `return a`, dan `return b`.

2. `Max(1, 3)` ini menghasilkan nilai 3, jika dirunut dari fungsi Max akan melewati 2 statement `if a >= b` dan karena kondisi tidak terpenuh maka langsung ke `return b`. Jadi hanya tercover 2 statement sehingga test coverage jika pengecekan `Max(4,3)` dihilangkan maka hasilnya **66.67%**.

3. Karena pemanggilan fungsi `Max(4,3)` menghasilkan jalur yang berbeda dengan proses pada no.2 yaitu `if a >= b` dan `return a`. Maka semua cabang logika dilewati, karena itulah hasil *test coverage* keseluruhan jadi **100%**.

Semoga yang sedikit ini bisa membantu pemahaman tentang *test coverage* menggunakan golang. Sampai jumpa lagi di tulisan selanjutnya.


[Golang Unit Test]: /posts/golang-unit-test/