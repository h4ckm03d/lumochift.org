---
title: "Golang Om Telolet Om Server"
date: 2018-01-03T08:49:20+07:00
type : "post"
description : "Menambahkan header x-telolet menggunakan golang"
author : "Moch. Lutfi"
categories : ["Golang101"]
tags : ["programming", "golang", "golang101"]
draft: false
toc : false
---

Tahun 2016 lalu ramai sekali tentang om telolet om, dan disalah satu website e-commerse di Indonesia menambahkan `x-om-telolet-om` di response header. Mungkin banyak yang tidak mengetahuinya, karena hanya tambahan response header dan tidak terlihat di tampilan website.

# Golang Om Telolet Om Server

Pertanyaannya bagaimana menambahkan respon http header menggunakan golang? Jawabnya tentu bisa dan sangat mudah sekali bahkan tidak perlu menggunakan eksternal dependensi tambahan.

Pertama perlu kita buat web server sederhana menggunakan `net/http` package seperti dibawah ini.

```go
package main

import (
    "fmt"
    "net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    http.ListenAndServe(":8080", nil)
}
// run dengan perintah `go run main.go`
```

Kita tes menggunakan `curl` hasilnya seperti di bawah ini masih tanpa ada tambahan om telolet om di header dan *response body* menampilkan string `Hello World!`

```bash
$curl -X HEAD -I localhost:8080
HTTP/1.1 200 OK
Date: Wed, 03 Jan 2018 02:13:14 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

$curl localhost:8080
Hello World!
```

Cara sederhana menambahkan header x-telolet hanya dengan menambahkan code `w.Header().Add("x-telolet", "OM telolet OM")` sebelum `fmt.Fprint(w, "Hello World!")` sehingga jadi seperti ini.

```go
package main

import (
    "fmt"
    "net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("x-telolet", "OM telolet OM")
    fmt.Fprint(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)
    http.ListenAndServe(":8080", nil)
}
```

Jika dicek menggunakan curl hasilnya sudah berubah.

```bash
$curl -X HEAD -I localhost:8080
HTTP/1.1 200 OK
X-Telolet: OM telolet OM
Date: Wed, 03 Jan 2018 02:18:58 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8
```

## Middleware

Kemudian muncul pertanyaan bagaimana jika handlernya banyak tidak hanya `HelloWorldHandler` saja? Untuk keperluan seperti ini bisa menggunakan `middleware`. Jadi menjalankan satu atau lebih fungsi tertentu sebelum menjalankan handler utama seperti `helloWorldHandler`.

Struktur `middleware` yang akan kita buat menggunakan `func(w http.ResponseWriter, r *http.Request)` sebagai tipe data. Jadi passing fungsi kedalam fungsi dan return fungsi lagi. Lebih jelasnya kita lihat code dibawah ini.

```go
package main

import (
    "fmt"
    "net/http"
)

// OmTeloletOm middleware
func OmTeloletOm(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("x-telolet", "OM telolet OM pake middleware")
        handler(w, r)
    }
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", OmTeloletOm(helloWorldHandler))
    http.ListenAndServe(":8080", nil)
}
```

Jika kita cek lagi dengan curl hasilnya seperti ini.

```bash
$curl -X HEAD -I localhost:8080
HTTP/1.1 200 OK
X-Telolet: OM telolet OM pake middleware
Date: Wed, 03 Jan 2018 02:34:48 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8
```

Sangat mudah ternyata menambahkan custom header menggunakan golang. Sampai jumpa lagi di [Golang 101] selanjutnya.

[Golang 101]: /categories/golang101/