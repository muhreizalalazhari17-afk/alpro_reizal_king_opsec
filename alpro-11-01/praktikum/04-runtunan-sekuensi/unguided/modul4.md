LAPORAN PRAKTIKUM MODUL 04

RUNTUNAN/SEKUENSI

Nama: Muh Reizal Al Azhari

---

I. Dasar Teori

A. Operator Perbandingan dan Operator Logika pada Go

1. Operator Perbandingan

Operator perbandingan digunakan untuk membandingkan dua nilai. Hasil dari perbandingan tersebut berupa nilai boolean, yaitu "true" atau "false". Operator ini dapat digunakan untuk menentukan kondisi yang akan diperiksa dalam suatu program.

Beberapa operator perbandingan yang tersedia dalam bahasa Go adalah:

- ">" : lebih besar dari.
- "<" : lebih kecil dari.
- ">=" : lebih besar dari atau sama dengan.
- "<=" : lebih kecil dari atau sama dengan.
- "==" : sama dengan.
- "!=" : tidak sama dengan.

2. Operator Logika

Operator logika digunakan untuk menggabungkan atau membalik kondisi boolean. Dalam bahasa Go, terdapat tiga operator logika utama:

- "&&" (AND): menghasilkan "true" jika kedua kondisi bernilai benar.
- "||" (OR): menghasilkan "true" jika setidaknya salah satu kondisi bernilai benar.
- "!" (NOT): membalik nilai boolean, dari "true" menjadi "false" atau sebaliknya.

Urutan prioritas operator logika adalah "!", kemudian "&&", dan terakhir "||". Penggunaan operator ini membantu dalam menyusun kondisi yang melibatkan lebih dari satu perbandingan.

B. Percabangan

1. If, Else If, dan Else

Percabangan "if", "else if", dan "else" digunakan untuk menentukan instruksi yang dijalankan berdasarkan kondisi tertentu. Program memeriksa kondisi secara berurutan. Jika suatu kondisi bernilai benar, blok kode yang sesuai akan dijalankan dan pemeriksaan pada rangkaian percabangan tersebut selesai.

Struktur ini cocok digunakan untuk menentukan kategori nilai, menghitung pajak berdasarkan rentang penghasilan, dan menyelesaikan permasalahan yang memiliki beberapa kondisi.

2. Switch Case

"switch" merupakan struktur percabangan yang dapat digunakan sebagai alternatif dari rangkaian "if-else if". Pada bahasa Go, "switch" dapat ditulis tanpa ekspresi sehingga setiap "case" berisi kondisi yang diperiksa satu per satu.

Program menjalankan "case" pertama yang sesuai. Secara umum, Go tidak memerlukan perintah "break" pada setiap "case" karena eksekusi berhenti secara otomatis setelah satu cabang dijalankan. Struktur ini membantu membuat kode lebih teratur ketika terdapat banyak kondisi.

---

II. Guided

1. Program grade.go

Kode Program

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}

Deskripsi Program

Program "grade.go" digunakan untuk menentukan nilai huruf siswa berdasarkan nilai angka yang dimasukkan. Nama siswa dibaca menggunakan "bufio.Scanner" agar masukan dapat berupa nama lengkap yang mengandung spasi. Setelah itu, program menerima nilai angka bertipe "float64".

Penentuan nilai huruf dilakukan menggunakan percabangan "if-else if-else". Rentang nilai 90–100 memperoleh grade A, 80 sampai kurang dari 90 memperoleh B, 70 sampai kurang dari 80 memperoleh C, dan 60 sampai kurang dari 70 memperoleh D. Nilai di luar rentang tersebut memperoleh grade F.

Hasil akhirnya ditampilkan menggunakan "fmt.Printf", yang menggabungkan nama siswa dengan nilai huruf yang diperoleh.

2. Program penilaian.go

Kode Program

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		fmt.Println("Keluar dari program.")

	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

Deskripsi Program

Program "penilaian.go" merupakan pengembangan dari program penentuan grade. Program ini menampilkan menu yang menyediakan dua pilihan, yaitu menjalankan sistem penilaian atau keluar dari program.

Ketika pengguna memilih angka 1, program meminta nama dan nilai siswa, kemudian menentukan grade menggunakan percabangan "if-else if-else". Jika pengguna memilih angka 0, program menampilkan pesan bahwa program telah keluar. Sementara itu, masukan selain angka 1 dan 0 akan menghasilkan pesan bahwa pilihan tidak valid.

Melalui program ini, penggunaan percabangan diterapkan untuk mengatur alur program berdasarkan pilihan pengguna.

3. Program klasifikasi.go

Kode Program

package main

import "fmt"

func main() {
	var usia int
	var gaji int
	var keterangan string

	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	switch {
	case usia < 18:
		keterangan = "Masih sekolah"

	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"

	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"

	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"

	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"

	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"

	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	fmt.Println(keterangan)
}

Deskripsi Program

Program "klasifikasi.go" digunakan untuk mengelompokkan seseorang berdasarkan usia dan jumlah gaji yang dimasukkan. Pengguna memasukkan dua bilangan bulat, yaitu usia dan gaji. Selanjutnya, program memeriksa kondisi menggunakan "switch" tanpa ekspresi.

Klasifikasi dilakukan berdasarkan beberapa rentang usia dan batas gaji. Seseorang yang berusia di bawah 18 tahun dikategorikan sebagai masih sekolah. Pada rentang usia 18–25 tahun, kategori ditentukan berdasarkan apakah gajinya mencapai 50 atau tidak. Untuk usia 26–40 tahun, batas gaji yang digunakan adalah 100. Sementara itu, untuk usia di atas 40 tahun, batas yang digunakan adalah 150.

Setiap kondisi memiliki keterangan yang berbeda. Setelah kondisi yang sesuai ditemukan, hasil klasifikasi disimpan pada variabel "keterangan" dan ditampilkan menggunakan "fmt.Println".

---

III. Unguided

1. Program grade dengan Switch

Kode Program

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"

	case nilai >= 80 && nilai < 90:
		grade = "B"

	case nilai >= 70 && nilai < 80:
		grade = "C"

	case nilai >= 60 && nilai < 70:
		grade = "D"

	default:
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}

Output Program

[Masukkan screenshot hasil eksekusi program grade dengan switch di sini.]

Deskripsi Program

Program ini merupakan pengembangan dari program "grade.go" dengan mengganti rangkaian percabangan "if-else if-else" menjadi "switch" tanpa ekspresi.

Program menerima nama dan nilai angka siswa sebagai masukan. Selanjutnya, nilai diperiksa berdasarkan rentang yang telah ditentukan. Nilai 90–100 memperoleh grade A, 80 sampai kurang dari 90 memperoleh B, 70 sampai kurang dari 80 memperoleh C, dan 60 sampai kurang dari 70 memperoleh D. Apabila tidak ada kondisi yang sesuai, bagian "default" menetapkan grade F.

Penggunaan "switch" membuat pengelompokan kondisi nilai menjadi lebih terstruktur. Hasil penilaian kemudian ditampilkan melalui "fmt.Printf".

2. Program pajak.go

Kode Program

package main

import "fmt"

func main() {
	var penghasilan float64
	var pajak float64

	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = 0.05*50 + 0.10*(penghasilan-50)
	} else if penghasilan <= 200 {
		pajak = 0.05*50 + 0.10*50 + 0.15*(penghasilan-100)
	} else {
		pajak = 0.05*50 + 0.10*50 +
			0.15*100 + 0.20*(penghasilan-200)
	}

	fmt.Println(pajak)
}

Output Program

[Masukkan screenshot hasil eksekusi program pajak di sini.]

Deskripsi Program

Program "pajak.go" digunakan untuk menghitung jumlah pajak berdasarkan nilai penghasilan yang dimasukkan pengguna. Variabel "penghasilan" dan "pajak" menggunakan tipe data "float64" sehingga dapat menyimpan nilai pecahan.

Penghitungan pajak dilakukan dengan percabangan "if-else if-else" menggunakan sistem tarif bertingkat. Penghasilan sampai dengan 50 dikenai tarif 5%. Untuk penghasilan di atas 50 sampai dengan 100, penghitungan menggunakan tarif 5% pada lapisan pertama dan 10% pada bagian berikutnya. Untuk penghasilan di atas 100 sampai dengan 200, tarif yang digunakan pada lapisan tambahan adalah 15%. Penghasilan di atas 200 dikenai tarif 20% pada bagian yang melebihi batas tersebut.

Pajak dihitung secara bertahap sesuai lapisan penghasilan, bukan dengan menerapkan satu tarif terhadap seluruh penghasilan. Nilai pajak yang diperoleh kemudian ditampilkan menggunakan "fmt.Println".

Perhitungan ini mengikuti asumsi batas penghasilan dan tarif yang digunakan dalam soal praktikum, bukan ketentuan pajak resmi.

---

IV. Kesimpulan

Berdasarkan praktikum Modul 04 tentang Runtunan/Sekuensi, dapat dipahami bahwa alur program tidak hanya terdiri dari instruksi yang dijalankan secara berurutan, tetapi juga dapat melibatkan percabangan untuk menentukan tindakan berdasarkan kondisi tertentu.

Percabangan "if-else if-else" dapat digunakan untuk memproses kondisi bertingkat, seperti penentuan grade dan penghitungan pajak. Sementara itu, "switch" tanpa ekspresi dapat digunakan sebagai alternatif untuk memeriksa beberapa kondisi, seperti pada program klasifikasi usia dan gaji maupun program penilaian.

Selain itu, operator perbandingan dan operator logika berperan dalam membentuk kondisi yang diperiksa oleh program. Penggunaan struktur tersebut membantu menghasilkan program yang lebih terorganisasi dan sesuai dengan kebutuhan penyelesaian masalah.

V. Referensi

1. The Go Authors. (2024). The Go Programming Language Specification. https://go.dev/ref/spec
2. Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. New York: Addison-Wesley.
3. Laboratorium Praktikum Informatika. (2025). Modul 04: Runtunan/Sekuensi – Algoritma dan Pemrograman. Bandung: Fakultas Informatika, Universitas Telkom.
