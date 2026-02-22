package validation

func IsPaymentValid(harga int, uang int) bool {
	return uang >= harga
}

func HitungKembalian (harga int, uang int) int {
	return uang - harga
}