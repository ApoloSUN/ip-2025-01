package main

import (
	"fmt"
)

func main () {
	var idade int64 
	
	fmt.Println("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade < 16 {
		fmt.Println("Sua classe eleitoral é Não-eleitor ")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Sua classe eleitoral é Eleitor obrigatório")
	} else {
		fmt.Println("Sua classe eleitoral é Eleitor facultativo")
	}
}