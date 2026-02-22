package main

import(
"fmt"
"github.com/fatih/color"
"transaksi-cust/validation"
)

func main() {
var harga int
var uang int

fmt.Println("harga barang: ")
fmt.Scan(&harga)

fmt.Println("uang pembeli: ")
fmt.Scan(&uang)

if validation.IsPaymentValid(harga,uang){
	kembalian := validation.HitungKembalian(harga, uang)
	color.Green("[SISTEM]: Transaksi Berhasil! Kembalian Anda:", kembalian)
} else {
	color.Red("[SISTEM]: Transaksi Ditolak! Uang kurang", harga-uang)
}
}