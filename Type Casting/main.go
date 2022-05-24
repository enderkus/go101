package main

import (
	"fmt"
	"reflect" // reflect.TypeOf ile değişken türü alınıyor.
	"strconv" // strconv.Atoi ya da strconv.Itoa ile stringten integer değerine ya da integer dan string türüne dönüşüm sağlanıyor.
)

func main() {
	/*
		Type Casting - Uyumlu veri tiplerinin bir birine dönüşümü.
	*/
	var a, b int = 100, 42               // İki integer tipte değişken tanımladık.
	var bolum = a / b                    // Bunları bolum adında bir değişken içerisinde böldük.
	fmt.Println(bolum)                   // Sonuç 2 ancak burada bir kayıp var. 2,380952380952381 gibi bir sonuç çıkıyor aslında. Biz bunu elde etmek istersek float dönüşümü yapmalıyız.
	var bolum2 = float32(a) / float32(b) // int olarak tanımladığımız a ve b değişkenlerini float32 türüne dönüştürdük.
	fmt.Println(bolum2)                  // Sonuç : 2.3809524 Dolayısıyla bolum2 değişkeni bir float türü olarak tanımlandı.
	var bolum3 = float64(a) / float64(b) // int olarak tanımladığımız a ve b değişkenlerini float64 türüne dönüştürdük.
	fmt.Println(bolum3)                  // Sonuç : 2.380952380952381 bolum3 değişkeni de artık float türünden bir değişken.
	var sonucint int = int(bolum3)       // bolum3 değerini tekrar int değerine çevirmek için bu yöntemi uyguluyoruz.
	fmt.Println(sonucint)                // Sonuç : 2

	/*
		Type Conversion - Uyumsuz veri tiplerinin dönüşümü.
	*/

	// String türünden int tipine dönüşüm.
	c := "3"                     // c isminde bir değişken tanımladık ve string tipinde ("3") değişken atadık.
	cToint, _ := strconv.Atoi(c) // String olan c değişken değerini strconv kütüphanesinin Atoi metodu ile int tipine dönüştürdük. Hata yakalama yerine ikincil değer olarak blank identifier tanımladık.
	cToint++                     // dönüşen 3 değerini 1 arttırdık.
	fmt.Println(cToint)          // sonuç 4
	// Şimdi cToint değerini tekrar string türüne  dönüştürelim.
	cTostr := strconv.Itoa(cToint) // Herhangi bir hata dndürmeyeceği için blank identifier tanımlamadık.
	fmt.Println(cTostr)            // Ekrana yazdırıyoruz sonuç 4
	// Reflect kütüphanesi ilede değişken türünü kontrol edelim.
	fmt.Println(reflect.TypeOf(cTostr)) // reflect.TypeOf fonksiyonu bize değişken türünü veriyor. Sonuç string

}
