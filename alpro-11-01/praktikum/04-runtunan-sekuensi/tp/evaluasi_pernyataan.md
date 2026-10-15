Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi? 25

Apa output yang dihasilkan oleh program? Nilai akhir result: 25

Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan: ◦ Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya? Benar, karena x bernilai 10 sehingga 10 > 5. Program lalu mengecek if yang ada di dalamnya, yaitu y < 10. Kondisi ini juga benar (5 < 10), maka result diisi dengan x + y = 15

◦ Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result? Benar, sebab kedua sisinya bernilai true: 15 > 10 dan 10 == 10. Karena itu program menjalankan result += z, sehingga result bertambah 15 dan menjadi 30. Blok else tidak dijalankan

◦ Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi? Benar. Operator || hanya butuh satu sisi yang true, dan x == 10 memenuhinya, walaupun y > 10 bernilai false. Program menjalankan result += 5, sehingga result menjadi 35, lalu else if dan else dilewati

◦ Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa dampaknya pada result? Kondisi dihitung mulai dari dalam kurung: x < 15 true dan y < 10 true, sehingga hasil && adalah true. Tanda ! membalikkannya menjadi false. Karena kondisi if bernilai false, program masuk ke blok else dan menjalankan result -= 10. Nilai result turun dari 35 menjadi 25, dan itulah nilai akhirnya
