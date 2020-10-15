---
title: "Golang SQL Database"
date: 2020-09-10T20:10:36+07:00
type : "post"
excerpt : "Dasar SQL database menggunakan golang"
authors:
  - Moch Lutfi
categories : ["Golang101"]
tags : ["programming", "golang", "golang101", "sql", "database"]
draft: false
toc : true
---

Dalam tulisan ini kita akan mempelajari bagaimana menggunakan SQL database di Golang. Mulai dari dasar [database/sql](http://golang.org/pkg/database/sql/), implementasi dalam pembuatan aplikasi, dan sedikit berbagi pengalaman pribadi dalam menggunakan database.

Database yang digunakan dalam tulisan ini menggunakan PostgreSQL, jangan khawatir untuk semua database driver hampir sama sintaksnya jika spesifik driver saya akan jelaskan juga.

Kita akan membuat simple `Pokedex` application dengan operasi CRUD pada tabel `pokemons`.

## Persiapan

Pertama-tama kita perlu membuat database `pokedex` dengan script sql sebagai berikut:
``` sql
-- migration.sql
CREATE TABLE pokemons (
    id int NOT NULL,
    name varchar(255) NOT NULL,
    species varchar(255) NOT NULL,
    height decimal(5,2) NOT NULL,
    weight decimal(5,2) NOT NULL
);

INSERT INTO pokemons (id, name, species, height, weight)  VALUES
(1, 'Bulbasaur', 'Seed Pokémon', 0.7, 6.9),
(2, 'Ivysaur', 'Seed Pokémon', 1, 12),
(3, 'Venusaur', 'Seed Pokémon', 2, 100),
(132, 'Ditto', 'Transform Pokémon', 0.3, 4),
(808, 'Meltan', 'Hex Nut Pokémon', 0.2, 8);

ALTER TABLE pokemons ADD PRIMARY KEY (id);
```

Kemudian tinggal buat database dan jalankan `migration.sql` diatas
```sh
# buat database pokedex
$ psql -h localhost -Upostgres -W -c "CREATE DATABASE pokedex";

# import migration sql ke database pokedex
$ psql -h localhost -Upostgres -W -d pokedex -a -f migration.sql
-- migration.ssql
CREATE TABLE pokemons (
    id int NOT NULL,
    name varchar(255) NOT NULL,
    species varchar(255) NOT NULL,
    height decimal(5,2) NOT NULL,
    weight decimal(5,2) NOT NULL
);
CREATE TABLE
INSERT INTO pokemons (id, name, species, height, weight)  VALUES
(1, 'Bulbasaur', 'Seed Pokémon', 0.7, 6.9),
(2, 'Ivysaur', 'Seed Pokémon', 1, 12),
(3, 'Venusaur', 'Seed Pokémon', 2, 100),
(132, 'Ditto', 'Transform Pokémon', 0.3, 4),
(808, 'Meltan', 'Hex Nut Pokémon', 0.2, 8);
INSERT 0 5
ALTER TABLE pokemons ADD PRIMARY KEY (id);
ALTER TABLE
```
Setelah persiapan database selesai sekarang dilanjutkan dengan bootstrap code.

``` sh
$ mkdir pokedex && cd pokedex
$ touch main.go
$ go mod init example.com/pokedex
go: creating new go.mod: module example.com/pokedex
```

## Memilih database driver

Alhamdulillah tidak sesulit memilih jodoh, untuk PostgreSQL tidak banyak pilihan dan [pq](https://github.com/lib/pq) merupakan jawaranya untuk PostgreSQL database. Untuk pilihan driver database lain bisa dicari sendiri di [daftar driver yang tersedia](https://github.com/golang/go/wiki/SQLDrivers). Untuk MYSQL driver yang paling kondang yaitu [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) 

## Membuat koneksi database

```go
// main.go
package main

import (
	"database/sql"
	"log"

  // Import the pq driver.
	_ "github.com/lib/pq"
)

func main() {
    // init koneksi database instance, tetapi masih belum konek ke database server
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/pokedex?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

    // validasi konfigurasi dengan ping ke database server
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("connected successfully!")
}
```

Penggunaan library `pq` dalam kode diatas menggunakan `blank identifier` `_` karena kita tidak menggunakan library tersebut secara langsung, tetapi secara tidak langsung via `sql` package. Pemanggilan dengan cara tersebut hanya menjalankan method `init()` didalam package `pq` ini berfungsi untuk [mendaftarkan dirinya sendiri](http://golang.org/pkg/database/sql/#Register) ke dalam `database/sql`. Pola seperti ini merupakan pendekatan baku untuk hampir semua Go sql driver.

Fungsi `sql.Open` mengembalikan nilai pointer [sql.DB](https://golang.org/pkg/database/sql/#DB), objek value `sql.DB` ini bukanlah sebuah koneksi database tapi merupakan pool koneksi dengan maksimum koneksi yg bisa diatur menggunakan `db.SetMaxOpenConns(integer_value)` dan `db.SetMaxIdleConns(int_value)`. Ilustrasi sederhana mengenai koneksi pool ini yaitu ketika kita menggunakan `sql.DB` maka driver sql akan mengambil 1 koneksi dari pool untuk digunakan dan kondisi total koneksi di pool sejumlah `N-1`, jika sudah selesai maka koneksi tersebut dikembalikan ke pool untuk digunakan dalam operasi yang lain.

Jika kita menjalankan kode sederhana diatas maka hasilnya seperti dibawah ini:
```sh
$ go run main.go
go: finding module for package github.com/lib/pq
go: downloading github.com/lib/pq v1.8.0
go: found github.com/lib/pq in github.com/lib/pq v1.8.0
19:54:15 connected successfully!
```

## Dasar SQL

Kita mulai dengan query receh `SELECT * FROM pokemons` kemudian kita tampilkan kedalam `stdout`.

```go
// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import the pq driver.
	_ "github.com/lib/pq"
)

type Pokemon struct {
	ID      int
	Name    string
	Species string
	Height  float64
	Weight  float64
}

func (p Pokemon) String() string {
	return fmt.Sprintf("%d, %s, %s, %.2f m, %.2f Kg", p.ID, p.Name, p.Species, p.Height, p.Weight)
}

func main() {
	// init koneksi database instance, tetapi masih belum konek ke database server
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/pokedex?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// validasi konfigurasi dengan ping ke database server
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("connected successfully!")

	rows, err := db.Query("SELECT * FROM pokemons")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	pokemons := make([]*Pokemon, 0)
	for rows.Next() {
		p := new(Pokemon)
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Species,
			&p.Height,
			&p.Weight,
		)
		if err != nil {
			log.Fatal(err)
		}
		pokemons = append(pokemons, p)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range pokemons {
		fmt.Println(p.String())
	}
}
```
Apa yang terjadi dengan kode diatas? Kita mendefinisikan tipe `Pokemon` untuk menampung hasil query database dari tabel `pokemons`. Tipe data masih menggunakan standar `int`, `string`, dan `float64` karena didalam database kita _paksa_ nilainya tidak boleh `nil` dengan `NOT NULL` ketika mendefinisikan tabel. Jika menggunakan nilai yang nullable maka kita perlu menampungnya kedalam `sql.NullString`, `sql.NullInt32`, dan `sql.NullFloat64`. Untuk menyederhanakan tulisan ini kita hindari dulu penggunaan nullable field.

Kita coba bedah melakukan apa saja kode diatas:

1. Mengambil resulset dari tabel `pokemons` menggunakan fungsi `DB.Query()` dan diletakkan di variabel `rows`. Setelah itu `defer rows.Close()` untuk memastikan resulset telah menutup koneksi kedalam database sebelum fungsi parent selesai. **Menutup resultset ini sangat penting**. Karena jika dibiarkan saja maka koneksi yang digunakan dalam mengambil data diatas tidak dikembalikan ke koneksi pool sehingga mempercepat kehabisan koneksi ke database.
2. Kemudian menggunakan `rows.Next()` untuk iterasi semua baris dalam resultset dan dilanjutkan dengan `rows.Scan()` untuk memindahkan data. Urutan ketika `rows.Scan()` ini sesuai query, dalam hal ini kita menggunakan `SELECT *` berarti urutannya sesuai dengan tabel didatabase, jika querynya `SELECT c, b, a` maka urutanya sesuai dengan deklarasi di `SELECT` statement yaitu c, b, a.
3. Ketika `rows.Next()` loop selesai kita panggil `rows.Err()`. Ini untuk memastikan jika ada error ketika melakukan iterasi, karena tidak semua iterasi diatas pasti selalu berakhir bahagia tanpa error.
4. Jika semuanya aman tanpa error maka tinggal loop variabel `pokemons` dan kita tampilkan informasi kedalam stdout.

Tampilan kode diatas sebagai berikut,

```sh
go run main.go
2020/10/09 20:09:18 connected successfully!
1, Bulbasaur, Seed Pokémon, 0.70 m, 6.90 Kg
2, Ivysaur, Seed Pokémon, 1.00 m, 12.00 Kg
3, Venusaur, Seed Pokémon, 2.00 m, 100.00 Kg
132, Ditto, Transform Pokémon, 0.30 m, 4.00 Kg
808, Meltan, Hex Nut Pokémon, 0.20 m, 8.00 Kg
```

Konversi ke web app

Saatnya aplikasi sederhana diatas _henshin_ ke REST ala-ala dengan 3 routes dan hanya menerima form request:

- GET /pokemons – daftar semua pokemon di pokedex
- GET /pokemons/show – menampilkan spesifik pokemon berdasarkan pokemon id
- POST /pokemons/create – menambahkan pokemon baru ke pokedex

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// Import the pq driver.
	_ "github.com/lib/pq"
)

type Pokemon struct {
	ID      int
	Name    string
	Species string
	Height  float64
	Weight  float64
}

func (p Pokemon) String() string {
    return fmt.Sprintf("%d, %s, %s, %.2f m, %.2f Kg", 
        p.ID, 
        p.Name, 
        p.Species, 
        p.Height, 
        p.Weight)
}

type Env struct {
	db *sql.DB
}

func main() {
	// init koneksi database instance, tetapi masih belum konek ke database server
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/pokedex?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// validasi konfigurasi dengan ping ke database server
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connected successfully!")

	env := &Env{db: db}
	http.Handle("/pokemons", pokemonsIndex(env))
	http.Handle("/pokemons/show", pokemonsShow(env))
	http.Handle("/pokemons/create", pokemonsCreate(env))
	http.ListenAndServe(":3000", nil)
}
```

HTTP handler untuk pembuatan web app kali ini menggunakan pendekatan _closure_, keutunganya tiap handler lebih bebas dalam menggunakan parameter input. Dalam aplikasi kali ini hanya `env` yang digunakan sebagai parameter input. Dalam `pokemonsShow` parsing data hanya menggunakan `r.FormValue` untuk menyederhanakan pembahasan kali ini, karena lebih fokus ke penggunaan database.

```go
func pokemonsIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		rows, err := env.db.Query("SELECT * FROM pokemons")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		pokemons := make([]*Pokemon, 0)
		for rows.Next() {
			p := new(Pokemon)
			err := rows.Scan(
				&p.ID,
				&p.Name,
				&p.Species,
				&p.Height,
				&p.Weight,
			)
			if err != nil {
				log.Fatal(err)
			}
			pokemons = append(pokemons, p)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}

		for _, p := range pokemons {
			fmt.Fprintln(w, p.String())
		}
	})
}
```

Untuk fungsi `pokemonsShow` memerlukan data dari user yaitu `id` yang didapat dari `r.FormValue("id")` dan passing parameter query menggunakan `$1` karena menggunakan `postgresql`, jika menggunakan Mysql maka placeholder `$1` perlu diganti dengan `?` agar tidak error. Pengambilan datanya hanya sekali tidak seperti di `pokemonsIndex`.
```go
func pokemonsShow(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if id == 0 || err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		row := env.db.QueryRow("SELECT * FROM pokemons WHERE id = $1", id)

		p := new(Pokemon)
		err = row.Scan(&p.ID, &p.Name, &p.Species, &p.Height, &p.Weight)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintf(w, p.String())
	})
}

```
Endpoint untuk menambahkan data juga sedikit berbeda dalam pengecekannya karena menggunakan `POST` request maka selain `POST` request dianggap error. Kali ini kita menggunakan execute statement karena bukan merupakan query.

```go
func pokemonsCreate(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		idStr := r.FormValue("id")
		name := r.FormValue("name")
		species := r.FormValue("species")
		heightStr := r.FormValue("height")
		weightStr := r.FormValue("weight")
		if idStr == "" || name == "" || species == "" || heightStr == "" || weightStr == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		height, err := strconv.ParseFloat(heightStr, 64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		result, err := env.db.Exec("INSERT INTO pokemons VALUES($1, $2, $3, $4, $5)", id, name, species, height, weight)
		if err != nil {
			fmt.Fprintf(w, "something wrong", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintf(w, "Pokemon %d created successfully (%d row affected)\n", id, rowsAffected)
	})
}
```

Jika semua digabungkan code diatas maka hasilnya seperti ini
```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// Import the pq driver.
	_ "github.com/lib/pq"
)

type Pokemon struct {
	ID      int
	Name    string
	Species string
	Height  float64
	Weight  float64
}

func (p Pokemon) String() string {
    return fmt.Sprintf("%d, %s, %s, %.2f m, %.2f Kg", 
        p.ID, 
        p.Name, 
        p.Species, 
        p.Height, 
        p.Weight)
}

type Env struct {
	db *sql.DB
}

func main() {
	// init koneksi database instance, tetapi masih belum konek ke database server
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/pokedex?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// validasi konfigurasi dengan ping ke database server
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("connected successfully!")

	env := &Env{db: db}
	http.Handle("/pokemons", pokemonsIndex(env))
	http.Handle("/pokemons/show", pokemonsShow(env))
	http.Handle("/pokemons/create", pokemonsCreate(env))
	http.ListenAndServe(":3000", nil)
}

func pokemonsIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		rows, err := env.db.Query("SELECT * FROM pokemons")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		pokemons := make([]*Pokemon, 0)
		for rows.Next() {
			p := new(Pokemon)
			err := rows.Scan(
				&p.ID,
				&p.Name,
				&p.Species,
				&p.Height,
				&p.Weight,
			)
			if err != nil {
				log.Fatal(err)
			}
			pokemons = append(pokemons, p)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}

		for _, p := range pokemons {
			fmt.Fprintln(w, p.String())
		}
	})
}

func pokemonsShow(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if id == 0 || err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		row := env.db.QueryRow("SELECT * FROM pokemons WHERE id = $1", id)

		p := new(Pokemon)
		err = row.Scan(&p.ID, &p.Name, &p.Species, &p.Height, &p.Weight)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintf(w, p.String())
	})
}

func pokemonsCreate(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		idStr := r.FormValue("id")
		name := r.FormValue("name")
		species := r.FormValue("species")
		heightStr := r.FormValue("height")
		weightStr := r.FormValue("weight")
		if idStr == "" || name == "" || species == "" || heightStr == "" || weightStr == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		height, err := strconv.ParseFloat(heightStr, 64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		result, err := env.db.Exec("INSERT INTO pokemons VALUES($1, $2, $3, $4, $5)", id, name, species, height, weight)
		if err != nil {
			fmt.Fprintf(w, "something wrong", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		fmt.Fprintf(w, "Pokemon %d created successfully (%d row affected)\n", id, rowsAffected)
	})
}
```

Sekian dulu pengenalan penggunaan database kali ini, pada tulisan selanjutkan kita akan membuat ReST API yang menggunakan `json` tapi tetap menggunakan studi kasus pokedex. 