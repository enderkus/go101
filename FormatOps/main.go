package main

import "fmt"

func main() {
	plaka := 22
	fmt.Printf("Edirne'nin plaka kodu: %d \n", plaka) // Format işlemlerinde Printf kullanılıyor. Int tipteki değerler için %d

	fmt.Printf("Edirne'nin plaka kodu: %b \n", plaka) // Binary olarak yazdırmak için %b

	pi := 3.1415
	fmt.Printf("Pi sayısı : %f \n", pi)   // Float değerler için %f
	fmt.Printf("Pi sayısı : %.2f \n", pi) // Noktadan sonra kaç karakter almasını istiyorsak %.2 gibi belirtiyoruz. Bu 2 karakter demek.

	sonuc := false
	fmt.Printf("Dönen değer : %t \n", sonuc) // Boolean değerler için %t kullanılıyor.

	dil := "Go"
	fmt.Printf("Programlama dilinin adı : %s \n", dil) // String değerler için %s
	cikis_yili := 2009
	fmt.Printf("%s programlama dili %d yılında çıkmıştır \n", dil, cikis_yili) // Birden fazla ve farklı türlerde değişken dahil edilebilir. Sırasıyla çağırılır.

	fmt.Printf("dil değişkeninin tipi %T \n", dil) // Değişkenlerin tipini almak için %T kullanılıyor. Type...

	go_hakkinda := fmt.Sprintf("%s programlama dili %d yılında çıkmıştır.", dil, cikis_yili) // Formatlanmış ifadeleri Sprtintf ile bir değişkene hapsedebiliyoruz.
	fmt.Println(go_hakkinda)                                                                 // Formatlanmış ifadeye sahip değişkeni Println ie çağırdık.

}
