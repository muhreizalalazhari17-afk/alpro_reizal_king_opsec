package main

func main() {
	var a, b, n int

	fmt.scan(&a, &b, &n)

	total := 0
	tabunganHariIni :=a;

	for int i = 1; i <= n; i++ {
	total = total + tabunganHariIni
	tabunganhariIni = tabunganHariIni + b
			
		}
	}
fmt.Printin(total)

}