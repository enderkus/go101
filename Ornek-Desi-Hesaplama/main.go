package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("### Desi Hesaplama Programı ###")
	fmt.Println("Lütfen ondalıklı değerleri nokta ile ayırın.")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Lütfen en değeri giriniz (Cm cinsinden) : ")
	scanner.Scan()
	en, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Lütfen boy değeri giriniz (Cm cinsinden ) : ")
	scanner.Scan()
	boy, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Lütfen yükseklik değeri giriniz (Cm cinsinden ) : ")
	scanner.Scan()
	yukseklik, _ := strconv.ParseFloat(scanner.Text(), 64)
	DesiHesapla := (en * boy * yukseklik) / 3000
	fmt.Printf("Kargonuzun desisi : %2.f", DesiHesapla)
}
