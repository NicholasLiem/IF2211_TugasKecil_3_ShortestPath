<h1> <b> IF2211 Tugas Kecil 3 Strategi Algoritma - Implementasi Algoritma UCS dan A* untuk
Menentukan Lintasan Terpendek</b> </h1>

## **Daftar Isi**
* [Deskripsi Program](#deskripsi-program)
* [Dependencies](#dependency)
* [Cara Menjalankan Program](#cara-menjalankan-program)
* [Penulis](#penulis)
* [Ekstra](#meme-section)

## **Deskripsi Program**
<p>Algoritma UCS (Uniform cost search) dan A* (atau A star) dapat digunakan untuk menentukan lintasan terpendek dari suatu titik ke titik lain. Pada tugas kecil
3 ini, anda diminta menentukan lintasan terpendek berdasarkan peta Google Map
jalan-jalan di kota Bandung. Dari ruas-ruas jalan di peta dibentuk graf. 

Simpul menyatakan persilangan jalan (simpang 3, 4 atau 5) atau ujung jalan. Asumsikan jalan dapat dilalui dari dua arah. Bobot graf menyatakan jarak (m atau km) antar simpul. Jarak antar dua simpul dapat dihitung dari koordinat kedua simpul menggunakan rumus jarak Euclidean (berdasarkan koordinat) atau dapat menggunakan
ruler di Google Map, atau cara lainnya yang disediakan oleh Google Map. Langkah pertama di dalam program ini adalah membuat graf yang merepresentasikan peta (di area tertentu, misalnya di sekitar Bandung Utara/Dago). 

Berdasarkan graf yang dibentuk, lalu program menerima input simpul asal dan simpul tujuan, lalu menentukan lintasan terpendek antara keduanya menggunakan algoritma UCS dan A*. Lintasan terpendek dapat ditampilkan pada peta/graf (misalnya jalan-jalan yang menyatakan lintasan terpendek diberi warna merah). Nilai
heuristik yang dipakai adalah jarak garis lurus dari suatu titik ke tujuan.</p>

## **Dependency**
1. Go 1.20.3

## **Cara Menjalankan Program**
1. Klon repositori ini <br>
`$ git clone https://github.com/NicholasLiem/Tucil3_13521083_13521135`
2. Pergi ke folder src <br>
`$ cd src`
3. Jalankan program <br>
`$ go run main.go`
4. Pergi ke localhost dibrowser anda <br>
`$ localhost:8080`

## **Penulis**
Moch. Sofyan Firdaus - 13521083 <br>
Nicholas Liem - 13521135

## **Meme Section**
![alt text](https://res.cloudinary.com/practicaldev/image/fetch/s--pxxN7gvW--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/rhmldpyrr2nwrmmcxo7k.png)
