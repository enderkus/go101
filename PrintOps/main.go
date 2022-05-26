package main

import "fmt"

func main() {
	fmt.Println("Merhaba Dünya ! \n Merhaba Go !") // \n alt satıra geçmemize yardımcı oluyor.
	fmt.Println("Merhaba Dünya! \t Merhaba Go !")  // \r 4 karakter boşluk bırakmamıza yardımcı oluyor.
	fmt.Println("Merhaba \"Dünya\"")               // Strinng içerisinde çift tırnak kullanmak için \ kaçısını kullanıyoruz.
	fmt.Println("Merhaba" + "Dünya" + "!")         // String ifadeleri birleştiriyor
	fmt.Println("Merhaba", "Dünya", "!")           // + birleştiricisinden farklı olarak arada boşluk bırakıyor.
	fmt.Println("Merhaba" +
		"Dünya" +
		"!") // Çoklu satırların birleşimi için bu yöntem kullanılabilir.

	a := "Go"
	b := "Programlama"

	fmt.Println(a, b) // String tipinde değişkenlerin birleştirilmesi işlemi.

	stringuzunlugu := len(a)
	fmt.Println(stringuzunlugu) // String tipindeki değerlerin karakter sayılarını verir. Not: Boşluklar dahil.
}
