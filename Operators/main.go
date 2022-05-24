package main

import "fmt"

func main() {
	a := 5
	b := 10

	/*
		Matematiksel operatörler.
	*/
	fmt.Println("a + b", a+b) // Toplama işlemi için +
	fmt.Println("a-b", a-b)   // Çıkarma işlemi için -
	fmt.Println("b/a", b/a)   // Bölme işlemi için /
	fmt.Println("b%a", b%a)   // Bölündenden kalan için %

	a++            // a değişkenine +1 ekler.
	fmt.Println(a) // Çıktısı 6 olacak.

	b--            //b değişkeninden 1 çıkarır.
	fmt.Println(b) // Çıktısı 9 olacak.

	/*
		Atama operatörleri
	*/
	c := 59
	d := 22

	var f = c + d  // Değişken içeriği tanımlamak için kullandığımız = bir atama operatörüdür.
	fmt.Println(f) // Çıktısı 81 olacak.

	c += d         // += Atama operatörü değişkenin kendisi ve belirttiğimiz değişkeni toplar ve belirttiğimiz değişkene yeni değeri atar.
	fmt.Println(c) // Bu işlemin açıklaması şöyle : c = c + d Sonuç yine 81 olacak ve artık c değişkeni 81 değerini alacak.

	c -= d         // -= Atama operatörü += atama operatörünün tam tersi. Yani değişkenin kendisinden belirttiğimiz değişkeni çıkarıp yeni değerini alır.
	fmt.Println(c) // Bu işlemin açıklaması şöyle : c = c - d Sonuç 59 olacak ve artık c 59 değerini alacak.

	c *= d         // *= Atama operatörü c değişkeninin kendisinin değeri ile belirttiğimiz değişkenin değerini çarparak c değişkeni için yeni tanımlama yapar.
	fmt.Println(c) // Bu işlemin açıklaması şöyle : c = c * d Sonuç 1298 olacak.

	c /= d         // /= Atama operatörü c değişkeninin değerini belirttiğimiz değişkene bölerek çıkan sonucu tekrar c değişkenine atar.
	fmt.Println(c) // Bu işlemin açıklaması şöyle : c = c / d Sonuç tekrar 59 oldu.

	c %= d         // %= Atama operatörü c değişkeni ile d değişkeninin bölümünden kalanı c değişkenine atar.
	fmt.Println(c) // Bu işlemin açıklaması şöyle : c = c % d Sonuç 15 olacak.

	/*
		İlişkisel Operatörler
	*/
	e := 10
	g := 10
	h := 5

	// İlişkisel operatörler == , != , < , > , >= , <=
	// İlişkisel operatörler iki değişkenin bir biri çeşitli denklik ya da denk olmama durumlarını kontrol eder.
	// İlişkisel operatör sonuçları boolean tipinde döner. Yani true ya da false.

	fmt.Println(e == g) // Burada e ve g değişkenlerinin değerleri bir birine eşit mi diye sorduk ve true döndü.
	fmt.Println(e == h) // Burada e ve h değişkenlerinin değerleri bir birine eşit mi diye sorduk ve false döndü.
	fmt.Println(e != g) // Burada 'e ve g birbirine eşit değil değil mi ?' diye sorduk ve eşit oldukları için false döndü.
	fmt.Println(e != h) // Burada 'e ve h birbirine eşit değil değil mi ?' diye sorduk ve eşit olmadığı için true döndü.
	fmt.Println(e > g)  // Burada e, g den büyük mü diye sorduk eşit oldukları için false döndü.
	fmt.Println(e > h)  // Burada e, h den büyük mü diye sorduk koşul karşılandığı için true döndü.
	fmt.Println(e < g)  // Burada e küçük mü g den diye sorduk ancak eşit oldukları için false döndü.
	fmt.Println(h < e)  // Burada h küçük mü e den diye sorduk koşul karşılandığı için true döndü.
	fmt.Println(e >= g) // Burada e, g den büyük mü ya da e,g ye eşit mi diye sorduk. Büyük değil ancak eşit olduklar için true döndü.
	fmt.Println(h <= g) // Burada h, g den küçük mü ya da eşit mi diye sorduk. h, değeri g değerinden küçük olduğu için true döndü.

	/*
		Mantıksal Operatörler
	*/
	var i, j, k, l = 1, 2, 3, 1
	// || , && , ! olarak kullanılırlar.

	fmt.Println(i == l || k == j) // Ya da operatörü. Şartlardan biri doğru olarak karşılanıyorsa true karşılanmıyorsa false döndürür.
	// Sorumuz şu : i ve l birbirine eşit mi ? Ya da... k ve j birbirine eşit mi ? İlk sorumuzun cevabı true, ikinci sorumuzun cevabı false ancak şartlardan biri true ve dönen değer true olacak.
	fmt.Println(i == l && k == j) // Ve operatörü iki şartında true olması halinde true iki şartında false olması halinda false iki şarttan biri true ve diğeri false ise false döndürür.
	// sorumuz şu : i ve l birbirine eşit mi ? Ve... k ve j bir birine eşit mi ? i ve l eşit yani true k ve j eşit değil yani false dönen değer false olacak.
	fmt.Println(i == l && !(k == j)) // ! yani NOT operatörü true değeri false, false değerini true olarak değiştirir. Yani tam tersi sonuç almamıza yarar.
	// İşlem şöyle oldu : i ve l bir birine eşit mi ? (true) Ve ... k ve j bir birine eşit mi ? (false) ama ikinci sorgumuzu parantez ile kapsadık ve başına ! koyduğumuzda false değeri true ya döndü ve sonuç true döndü.
}
