package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Mutlak değer alma
	*/

	fmt.Println(math.Abs(-10))   // Bir sayının mutlak değerini almak için Aps fonksiyonundan yararlanırız.
	fmt.Println(math.Abs(-12.2)) // Float tipinde değerler için de kullanabiliriz.
	fmt.Println(math.Abs(5))     // Eksi yönlü olmayan sayılar içinde kullanılabilir.

	/*
		Sinüs kosinüs ve tanjant
	*/

	fmt.Println(math.Sin(4))  // Sinüs hesaplama. Float tipinde döner.
	fmt.Println(math.Cos(12)) // Kosinüs hesaplama. Float tipinde döner.
	fmt.Println(math.Tan(24)) // Tanjant hesaplama. Float tipinde döner.

	/*
		Log almak
	*/

	fmt.Println(math.Log(89))   // Doğal yani e olarak log alma işlemi gerçekleşir. Float tipinde döner.
	fmt.Println(math.Log10(90)) // 10lu logaritma alınır. Float tipinde döner.

	/*
		Mod hesaplama
	*/

	fmt.Println(math.Mod(42, 25)) // Mod alma işlemi için .Mod fonksiyonundan yararlanılır. Float tipinde döndürülebilir.

	/*
		Yuvarlama işlemi
	*/

	fmt.Println(math.Ceil(25.4)) // Float değerleri yukarı yuvarlar. Int tipinde değer döndürür.
	fmt.Println(math.Ceil(25.5))
	fmt.Println(math.Ceil(25.6))

	/*
		Minimum ve maximum bulma
	*/

	fmt.Println(math.Min(14.5, 25)) // Minimumu bulur iki değer alır. Float değerler girilebilir.
	fmt.Println(math.Max(14.5, 25)) // Maximum değeri bulur. Float değerler kullanabilir.

	/*
		Üs hesaplama
	*/

	fmt.Println(math.Pow(10.5, 2)) // İki değer alır, float girilebilir. İlk değer sayı ikinci değer üssüdür.
}
