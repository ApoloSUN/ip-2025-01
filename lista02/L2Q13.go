package main

import (
	"fmt"
)

func main() {
	var numero int64
	

	fmt.Println("Digite um número inteiro positivo de 3 dígitos:")
	fmt.Scanln(&numero)

	if numero >= 100 && numero <= 999 {
	        dezena := (numero / 10) % 10     
		    fmt.Println("O algarismo da casa das dezenas é", dezena)
	
	} else { 
	         fmt.Println("Número inválido")
	}
}