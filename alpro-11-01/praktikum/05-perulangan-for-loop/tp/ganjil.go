or Tiga Komponen vs for Hanya Kondisi Go tidak menyediakan while dan do-while, sehingga perannya diambil alih oleh for.

for tiga komponen cocok ketika jumlah putaran sudah diketahui. Inisialisasi, kondisi, dan perubahan nilai ditulis dalam satu baris

for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
for hanya kondisi bekerja seperti while. Inisialisasi ditulis di luar for, dan perubahan nilai ditulis di dalam badan perulangan. Cocok ketika jumlah putaran belum diketahui di awal

i := 1
for i <= 5 {
	fmt.Println(i)
	i++
}
| Aspek | for tiga komponen | for hanya kondisi | | Penulisan | Inisialisasi, kondisi, dan post dalam satu baris | Hanya kondisi di samping for | | Inisialisasi | Di dalam for | Di luar, sebelum for | | Perubahan nilai | Otomatis oleh post statement | Ditulis sendiri di dalam badan | | Cocok untuk | Jumlah putaran sudah pasti | Jumlah putaran belum pasti |
