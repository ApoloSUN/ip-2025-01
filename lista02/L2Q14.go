package main

import (
	"fmt"
	"strings"
)

func main() {
	var inicial, final float64
	var resposta string
	

	fmt.Println("Digite o preço inicial do carro:")
	fmt.Scanln(&inicial)
	final = inicial 

	fmt.Println("Deseja adicionar ar condicionado? (S/N)")
	fmt.Scanln(&resposta)
	if strings.ToUpper (resposta) == "S" {
		final = final + 1750.0
	} 

	fmt.Println("Deseja adicionar pintura metálica? (S/N)")
	fmt.Scanln(&resposta)
	if strings.ToUpper (resposta) == "S" {
		final = final +  800.0
	} 

	fmt.Println("Deseja adicionar vidro elétrico? (S/N)")
	fmt.Scanln(&resposta)
	if strings.ToUpper (resposta) == "S" {
		final = final +  1200.0
	} 

	fmt.Println("Deseja adicionar direção hidráulica? (S/N)")
	fmt.Scanln(&resposta)
	if strings.ToUpper (resposta) == "S" {
		final = final +  2000.0
	} 

	fmt.Println("O preço final do carro é: R$", final)
}