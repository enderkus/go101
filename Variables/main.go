// Değişkenler özel karakter ve sayıyla başlayamaz.
package main

import "fmt"

func main() {
	var mesaj string = "Merhaba Dünya !"
	fmt.Println(mesaj)

	var mesaj2 string
	mesaj2 = "Merhaba Dünya !"
	fmt.Println(mesaj2)

	var sayi int = 1234

	fmt.Println(sayi)

	var sayi2 float32 = 3.14
	fmt.Println(sayi2)

	var kontrol bool = true

	fmt.Println(kontrol)

	var x, y, z int = 1, 2, 3

	fmt.Println(x, y, z)

	var isim = "Ender"

	fmt.Println(isim)

	var a, b, c = 22, false, "Edirne"

	fmt.Println(a, b, c)

	sehir := "Edirne" // Bu tipteki değişkenler main içerisinde işlem görür.

	fmt.Println(sehir)

}
