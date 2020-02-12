---
hero: /images/hero-2.jpg
authors:
  - Moch Lutfi
categories : ["computer science"]
tags : ["computer science", "story", "async","concept"]
date : 2013-10-12T20:02:21+07:00
description : "Antara aku, async, dan sync"
featuredalt : ""
featuredpath : "date"
linktitle : ""
title : "Konsep synchronous dan asynchronous"
---
Dalam postingan pertama ini saya akan sedikit mengulas tentang konsep asynchronous dalam analogi sederhana sehingga bisa memudahkan pembaca dalam memahaminya.

Sebelum itu perlu kita ketahui ada proses synchronous dan asynchronous, apa bedanya? Mari kita simak analogi sederhana berikut.

Untuk penjelasan selanjutnya saya menggunakan istilah sync untuk proses synchronous dan async untuk proses asynchronous.

Pertama-tama akan saya gambarkan dalam bentuk proses yg berurutan:

## State awal

1. Kondisi awal, sedang main DOTA 2. Lagi seru-serunya nih, dapet hero favorit juga.
2. Perut udah dangdutan gara-gara seharian belum makan dan di dapur semua makan ludes, apesnya lagi sendiri juga dirumah.

## proses sync

1. Demi kesehatan akhirnya meninggalkan peperangan dan pergi ke toko terdekat beli mie instant (kira-kira butuh waktu 10 menit, maklum jalan kaki cuy).
2. Sesampainya dirumah langsung masak mie (butuh waktu 10 menit), udah dicacat-cacatin temen nih di chat T_T
3. Akhirnya bisa makan juga, meski dengan penuh pengorbanan 1 match. :(
4. Setelah makan balik ngadep PC dan memelas serta minta maaf ke anggota clan. Maaf ane leave gan.

## Proses async

1. Disaat lapar pun tetap tenang dan dengan sigap ambil hape trus pesen makanan cepat saji, dan diakhir pembicaraan dengan lantang bilang “ga pake lama ya mbak :3”. (ga sampai 1 menit)
2. Lanjut main DOTA deh sambil nunggu pesenan datang.
3. Ting tong bunyi bel rumah bunyi. Alhamdulillah pesenan dah dateng dan pas match baru aja selesai.
4. Ambil pesenan dan ga lupa bayar. Trus makan deh.
5. Setelah makan balik WAR lagi. uyeaah.

> Dari cuplikan proses diatas dapat diambil kesimpulan pada saat proses sync, tiap proses harus diselesaikan terlebih dahulu agar bisa melanjutkan ke proses selanjutnya dan harus berurutan.

Berbeda dengan async proses, pada async proses masih tetap bisa pada state awal tetapi dalam kondisi menunggu proses yang berjalan diluar system seperti proses masak makanan cepat saji sampai makanan sampai di rumah. Dan pada saat proses lain tersebut selesai ia akan memberi tanda bahwa proses sudah selesai.

Dalam konteks async diatas, notifikasi atau tanda jika proses lain adalah bel rumah berbunyi berarti tanda kurir makanan cepat saji sudah datang. Notifikasi seperti ini dalam bahasa pemrograman disebut callback yang berisi informasi proses dan proses tambahan seperti mbayar makanan. :p

Semoga sedikit kisah ga jelas tersebut bisa membantu memahami konsep asynchronous.
