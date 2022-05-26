package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Lütfen en değerini giriniz (Örn : 2.15) : \n")
	scanner.Scan()
	en, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Lütfen alanın boy değerini giriniz (Örn : 5.20) : \n")
	scanner.Scan()
	boy, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("Girdiğiniz en : %f. Girdiğiniz boy : %f. Toplam metrekare %f", en, boy, en*boy)

}
