LAPORAN PRAKTIKUM MODUL 05

PERULANGAN FOR

Nama: Muh Reizal Al Azhari

I. Dasar Teori

A. Perulangan "for" dalam Bahasa Go

Perulangan "for" merupakan struktur kontrol yang digunakan untuk menjalankan sekumpulan perintah secara berulang selama kondisi tertentu terpenuhi. Dalam bahasa pemrograman Go, "for" merupakan satu-satunya struktur perulangan yang tersedia.

Bentuk umum perulangan "for" dengan tiga komponen adalah sebagai berikut:

for inisialisasi; kondisi; post {
    // Pernyataan yang diulang
}

Ketiga komponen tersebut memiliki fungsi yang berbeda. Inisialisasi dijalankan satu kali sebelum perulangan dimulai, kondisi diperiksa sebelum setiap iterasi, sedangkan pernyataan post dijalankan setelah setiap iterasi selesai.

B. Perulangan dengan Kondisi dan Perulangan Bersarang

Bahasa Go tidak menyediakan struktur "while" secara terpisah. Fungsinya dapat digantikan dengan perulangan "for" yang hanya menggunakan kondisi.

Selain itu, perulangan "for" dapat ditempatkan di dalam perulangan lainnya. Teknik ini disebut nested loop atau perulangan bersarang. Perulangan bersarang dapat digunakan untuk memproses data yang terdiri atas beberapa kelompok nilai, misalnya membaca enam nilai untuk setiap mahasiswa.

II. Guided

1. Program "tabungan.go"

Kode Program

package main

import "fmt"

func main() {
    var a, b, n int
    fmt.Scan(&a, &b, &n)

    total := 0
    tabunganHariIni := a

    for i := 1; i <= n; i++ {
        total = total + tabunganHariIni
        tabunganHariIni = tabunganHariIni + b
    }

    fmt.Println(total)
}

Deskripsi Program

Program ini digunakan untuk menghitung total tabungan selama "n" hari. Nilai "a" menunjukkan jumlah tabungan pada hari pertama, sedangkan "b" merupakan kenaikan tabungan setiap hari. Variabel "total" digunakan untuk menyimpan akumulasi tabungan. Pada setiap iterasi, tabungan hari tersebut ditambahkan ke total, kemudian jumlah tabungan harian dinaikkan sebesar "b".

2. Program "terbesar.go"

Kode Program

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var angka, max int
    fmt.Scan(&angka)
    max = angka

    for i := 2; i <= n; i++ {
        fmt.Scan(&angka)

        if angka > max {
            max = angka
        }
    }

    fmt.Println(max)
}

Deskripsi Program

Program ini digunakan untuk mencari bilangan terbesar dari sejumlah "n" bilangan yang dimasukkan oleh pengguna. Bilangan pertama disimpan sebagai nilai awal "max". Selanjutnya, program membaca bilangan berikutnya melalui perulangan. Jika bilangan yang dibaca lebih besar daripada nilai "max", nilai tersebut akan menggantikan nilai maksimum sebelumnya. Setelah seluruh bilangan diproses, program menampilkan bilangan terbesar.

3. Program "bioskop.go"

Kode Program

package main

import "fmt"

func main() {
    var t int
    fmt.Print("Input t (jumlah transaksi): ")
    fmt.Scan(&t)

    totalPendapatan := 0

    for i := 1; i <= t; i++ {
        var jenisFilm, hari string
        var jumlahTiket int

        fmt.Printf(
            "Input transaksi ke-%d (jenis, hari, tiket): ", i,
        )
        fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

        hargaSatuan := 0

        if jenisFilm == "Reguler" {
            if hari == "Senin" || hari == "Selasa" ||
                hari == "Rabu" || hari == "Kamis" {
                hargaSatuan = 40000
            } else {
                hargaSatuan = 50000
            }
        } else if jenisFilm == "3D" {
            if hari == "Senin" || hari == "Selasa" ||
                hari == "Rabu" || hari == "Kamis" {
                hargaSatuan = 60000
            } else {
                hargaSatuan = 70000
            }
        } else if jenisFilm == "IMAX" {
            if hari == "Senin" || hari == "Selasa" ||
                hari == "Rabu" || hari == "Kamis" {
                hargaSatuan = 80000
            } else {
                hargaSatuan = 90000
            }
        }

        totalHarga := hargaSatuan * jumlahTiket

        if jumlahTiket >= 5 {
            diskon := totalHarga * 10 / 100
            totalHarga = totalHarga - diskon
        }

        fmt.Printf(
            "Output total harga transaksi ke-%d: %d\n",
            i, totalHarga,
        )

        totalPendapatan += totalHarga
    }

    fmt.Println("Output total pendapatan:", totalPendapatan)
}

Deskripsi Program

Program ini menghitung total harga tiket untuk setiap transaksi bioskop sekaligus menghitung total pendapatan dari seluruh transaksi. Harga tiket ditentukan berdasarkan jenis film, yaitu Reguler, 3D, atau IMAX, serta kategori hari. Harga pada hari Senin sampai Kamis berbeda dengan harga pada hari lainnya.

Jika jumlah tiket dalam satu transaksi mencapai lima tiket atau lebih, program memberikan diskon sebesar 10% dari total harga transaksi. Setelah itu, total harga setiap transaksi ditampilkan dan ditambahkan ke variabel "totalPendapatan". Pada akhir program, total pendapatan dari seluruh transaksi ditampilkan.

III. Unguided

1. Program "prima.go"

Kode Program

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    jumlah := 0

    for i := 1; i <= n; i++ {
        if n%i == 0 {
            if jumlah > 0 {
                fmt.Print(" ")
            }

            fmt.Print(i)
            jumlah++
        }
    }

    fmt.Println()

    if jumlah == 2 {
        fmt.Println("Ya")
    } else {
        fmt.Println("Tidak")
    }
}

Output

Screenshot Output Unguided 1

Deskripsi Program

Program ini digunakan untuk menampilkan seluruh faktor dari bilangan "n" dan menentukan apakah bilangan tersebut termasuk bilangan prima. Perulangan dilakukan dari angka 1 sampai dengan "n". Jika "n" habis dibagi oleh nilai "i", nilai tersebut merupakan faktor dari "n" dan akan ditampilkan.

Variabel "jumlah" digunakan untuk menghitung banyaknya faktor yang ditemukan. Bilangan dinyatakan prima apabila memiliki tepat dua faktor positif, yaitu 1 dan bilangan itu sendiri. Program menampilkan "Ya" jika bilangan memenuhi kriteria tersebut dan "Tidak" jika tidak memenuhi.

2. Program "ecct.go"

Kode Program

package main

import "fmt"

func main() {
    var m, nilai, lulus int
    fmt.Scan(&m)

    for i := 0; i < m; i++ {
        jumlah := 0

        for j := 0; j < 6; j++ {
            fmt.Scan(&nilai)
            jumlah += nilai
        }

        rata := float64(jumlah) / 6

        if rata >= 3.33 {
            fmt.Println("Ya")
            lulus++
        } else {
            fmt.Println("Tidak")
        }
    }

    fmt.Println(lulus)
}

Output

Screenshot Output Unguided 2

Deskripsi Program

Program ini digunakan untuk menentukan kelulusan mahasiswa dalam memperoleh sertifikat ECCT berdasarkan enam nilai yang dimiliki setiap mahasiswa. Variabel "m" menunjukkan jumlah mahasiswa yang akan diproses.

Perulangan luar digunakan untuk memproses setiap mahasiswa, sedangkan perulangan dalam digunakan untuk membaca dan menjumlahkan enam nilai mahasiswa tersebut. Setelah seluruh nilai dijumlahkan, program menghitung rata-rata dengan mengonversi hasil penjumlahan ke tipe "float64".

Mahasiswa dinyatakan lulus apabila memperoleh rata-rata nilai minimal 3,33. Program menampilkan "Ya" jika mahasiswa memenuhi kriteria dan "Tidak" jika tidak memenuhi. Variabel "lulus" digunakan untuk menghitung jumlah mahasiswa yang berhasil memperoleh sertifikat, kemudian hasilnya ditampilkan setelah seluruh mahasiswa selesai diproses.

IV. Kesimpulan

Berdasarkan praktikum Modul 05 tentang Perulangan "for", dapat disimpulkan bahwa perulangan digunakan untuk menjalankan instruksi secara berulang sehingga proses pengolahan data menjadi lebih efisien. Perulangan dapat dimanfaatkan untuk menghitung akumulasi tabungan, mencari nilai terbesar, menghitung transaksi, menentukan bilangan prima, serta memproses nilai mahasiswa.

Praktikum ini juga menunjukkan bahwa perulangan dapat dikombinasikan dengan percabangan untuk menentukan kondisi tertentu. Selain itu, penggunaan perulangan bersarang memungkinkan program memproses beberapa kelompok data secara terstruktur. Dengan memahami konsep perulangan "for", program dapat dibuat lebih sistematis, efisien, dan mudah dipahami.

V. Referensi

1. The Go Authors. (2024). The Go Programming Language Specification. Diakses melalui https://go.dev/ref/spec
2. Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. New York: Addison-Wesley.
3. Laboratorium Praktikum Informatika. (2025). Modul 05: Perulangan For – Algoritma dan Pemrograman. Bandung: Fakultas Informatika, Universitas Telkom.
