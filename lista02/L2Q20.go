package main

import (
	"fmt"
)

func main() {
	var pagamento string 
	var precoN float64
	
	fmt.Println("Digite o preço inicial do produto e a forma de pagamento:  ")
	fmt.Scanln(&precoN, &pagamento)

	if pagamento == "1" {
		precoF := precoN - (precoN * 0.1)
		fmt.Println("Preço final: R$", precoF)
	} else if pagamento == "2" {
		precoF := precoN - (precoN * 0.05)
		fmt.Println("Preço final: R$", precoF)
	} else if pagamento == "3" {
		precoF := precoN
		parcela := precoF / 2 
		fmt.Println("Preço final: R$", precoF, "em 2 parcelas de R$", parcela)		
	} else {
		precoF := precoN + (precoN * 0.1)
		parcela := precoF / 3 
		fmt.Println("Preço final: R$", precoF, "em 3 parcelas de R$", parcela)		
	}
}