// Sabit Değişkenlerin değerleri daha sonra değiştirilemez.
package main

import "fmt"

const pi = 3.14

func main() {
	fmt.Println(pi)

	const kg = "1000"

	fmt.Println(kg)

	/*
		Sabitleri main fonksiyonu ya da dışında oluşturarak çağırabiliriz.
		Atanan değerler daha sonra değiştirilemez. Hata verir...
	*/
}
