package main

import (
	"fmt"
)

func main() {
	var (
		idade int64
	)

	fmt.Println("Digite a idade:")
	fmt.Scanln(&idade)

	if idade >= 0 && idade <= 2 {
	         fmt.Println("Récem-nascido")
	
	} else if idade >= 3 && idade <= 11 {
	         fmt.Println("Criança")
	
	} else if idade >= 12 && idade <= 19 { 
	         fmt.Println("Adolescente")
	
	} else if idade >= 20 && idade <= 55 { 
	         fmt.Println("Adulto")
	
	} else { 
	         fmt.Println("Idoso")
	}
}