package main

import "fmt"

func main() {
	var list[10]int
	var i int 
	var listP []int
	var somalistP int = 0
	var listI []int
	var QuantI int = 0

	for i = 0; i < 10; i = i + 1 {
		fmt.Scan(&list[i])
		if list[i] % 2 == 0 {
			listP = append(listP, list[i])
			somalistP = somalistP + list[i]
		} else {
			listI = append(listI, list[i])
			QuantI = QuantI + 1
		}
	}

	fmt.Println("Números pares digitados:", listP)
	fmt.Println("Soma dos números pares:", somalistP)
	fmt.Println("Números ímpares digitados:", listI)
	fmt.Println("Quantidade de números ímpares digitados:", QuantI)
	
}