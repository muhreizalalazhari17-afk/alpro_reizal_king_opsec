TUGAS PENDAHULUAN MODUL 05

PERULANGAN FOR

Nama: Muh Reizal Al Azhari

1. Program "ganjil.go"

Kode Program

package main

import "fmt"

func main() {
    var n, jumlah int
    fmt.Scan(&n)

    for i := 1; i <= n; i += 2 {
        if i > 1 {
            fmt.Print(" ")
        }

        fmt.Print(i)
        jumlah += i
    }

    fmt.Println()
    fmt.Println(jumlah)
}

Output

Screenshot Output Program "ganjil.go"

Deskripsi Program

Program ini digunakan untuk menampilkan bilangan ganjil mulai dari 1 hingga batas nilai "n", sekaligus menghitung jumlah seluruh bilangan ganjil yang ditampilkan.

Variabel "i" diinisialisasi dengan nilai 1 dan ditambah 2 pada setiap iterasi menggunakan "i += 2". Dengan demikian, perulangan hanya memproses bilangan ganjil. Setiap bilangan ditampilkan dan ditambahkan ke variabel "jumlah". Setelah perulangan selesai, program menampilkan jumlah seluruh bilangan ganjil tersebut.

2. Program "perkalian.go"

Kode Program

package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    for i := 1; i <= b; i++ {
        fmt.Printf("%d x %d = %d\n", a, i, a*i)
    }
}

Output

Screenshot Output Program "perkalian.go"

Deskripsi Program

Program ini digunakan untuk menampilkan tabel perkalian dari bilangan "a", mulai dari perkalian dengan 1 hingga perkalian dengan "b".

Variabel "i" berfungsi sebagai penghitung iterasi yang dimulai dari 1 dan bertambah satu pada setiap putaran. Selama nilai "i" tidak melebihi "b", program menghitung hasil perkalian "a * i" dan menampilkannya dalam format "a x i = hasil". Perulangan berhenti ketika nilai "i" melebihi batas "b".

Kesimpulan

Berdasarkan pengerjaan tugas pendahuluan Modul 05 tentang Perulangan "for", dapat disimpulkan bahwa perulangan digunakan untuk menjalankan instruksi secara berulang tanpa perlu menuliskan perintah yang sama berkali-kali. Dalam bahasa Go, perulangan "for" dapat digunakan untuk mengolah data berdasarkan batas iterasi tertentu.

Pada program "ganjil.go", perulangan digunakan untuk menampilkan bilangan ganjil sekaligus menghitung jumlahnya. Sementara itu, pada program "perkalian.go", perulangan digunakan untuk menghasilkan tabel perkalian berdasarkan bilangan dan batas yang dimasukkan oleh pengguna. Kedua program tersebut menunjukkan bahwa perulangan "for" dapat membantu menyelesaikan permasalahan secara lebih terstruktur dan efisien.
