package main

import (
	"bufio" // Giriş bilgilerini taramamız için gerekli paket.
	"fmt"
	"os" // Os paketi işletim sisteminden gelen standart input değerini almak için kullanıyoruz.
	"strconv"
)

func main() {
	fmt.Print("Adınız nedir ? ")        // Burada kullanıcıdan örnek veri almak için yazdırdık...
	giris := bufio.NewScanner(os.Stdin) // Kullanıcıdan gelen verinin tutulacağı değişken tanımlamasını yaptık.
	// bufio.NewScanner fonksiyonu ile neyi alacağımızı belirliyoruz. os.Stdin yani standart input değerlerini alacağız dedik.
	giris.Scan()                               // Kullanıcıdan gelen veriyi burada tarıyoruz.
	gelen_veri := giris.Text()                 // Gelen veriyi alacğız. .Text fonksiyonu ile gelen veriyi Text formatında aldık.
	fmt.Printf("Hoşgeldin, %s \n", gelen_veri) // gelen veriler String olarak döndüğü için %s ile çağırdık.

	int_giris := bufio.NewScanner(os.Stdin)
	fmt.Print("Go kaç yılında duyuruldu ? ")
	int_giris.Scan()
	int_gelen_veri, _ := strconv.ParseInt(int_giris.Text(), 10, 64) // strconv kütüphanesinin ParseInt fonksiyonunu kullanıyoruz. Gelen veriyi int türüne dönüştürüyoruz.
	// int_giris.text() ile hangi veriyi int tipine dönüştüreceğimizi belirledik.
	// sonrasında ondalıklı değeri tanımladık.
	// Bit türünü tanımladık.
	fmt.Printf("Go programlama dili %d yılında duyuruldu.", int_gelen_veri)

}
