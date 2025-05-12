package main

import (
		"fmt"
)
	
func main() {
		var numero int64 
		var n1, n2, n3, ME float64
		
		fmt.Println("Digite o número de identificação, as três notas do aluno e a média dos exercícios: ")
		fmt.Scanln(&numero, &n1, &n2, &n3, &ME)
	
		Aproveitamento := (n1 + (n2 * 2) + (n3 * 3) + ME) / 7
	
		if Aproveitamento >= 9.0 && Aproveitamento <= 10.0 {
			fmt.Println("Número do aluno: ", numero)
			fmt.Println("Notas, média dos exercícios e média de aproveitamento: ", n1, ",", n2, ",", n3, ",", Me, ",", Aproveitamento)
			fmt.Println("Aprovado, Conceito A")
		} else if Aproveitamento >= 7.5 && Aproveitamento < 9.0 {
			fmt.Println("Número do aluno: ", numero)
			fmt.Println("Notas, média dos exercícios e média de aproveitamento: ", n1, ",", n2, ",", n3, ",", Me, ",", Aproveitamento)
			fmt.Println("Aprovado, Conceito B")
		} else if Aproveitamento >= 6.0 && Aproveitamento < 7.5 {
			fmt.Println("Número do aluno: ", numero)
			fmt.Println("Notas, média dos exercícios e média de aproveitamento: ", n1, ",", n2, ",", n3, ",", Me, ",", Aproveitamento)
			fmt.Println("Aprovado, Conceito C")
		} else if Aproveitamento >= 4.0 && Aproveitamento < 6.0 {
			fmt.Println("Número do aluno: ", numero)
			fmt.Println("Notas, média dos exercícios e média de aproveitamento: ", n1, ",", n2, ",", n3, ",", Me, ",", Aproveitamento)
			fmt.Println("Reprovado, Conceito D")
		} else {
			fmt.Println("Número do aluno: ", numero)
			fmt.Println("Notas, média dos exercícios e média de aproveitamento: ", n1, ",", n2, ",", n3, ",", Me, ",", Aproveitamento)
			fmt.Println("Reprovado, Conceito E")
		}
}