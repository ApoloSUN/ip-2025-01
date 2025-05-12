package main

import (
	"fmt"
)

func main () {
	var matricula, horas int64 
	const Sminimo = 788.0 
	const VHE = 10.0
	
	fmt.Println("Digite a matrícula do funcionário e quantidade de horas extras trabalhadas: ")
	fmt.Scanln(&matricula, &horas)

	SHE := VHE * horas 
	Sbruto := 3 * Sminimo + SHE
	DescontoINSS := 0.0
	DescontoIR := 0.0

	if Sbruto > 1500.0 {
		DescontoINSS = 0.12 * Sbruto
	} 

	if Sbruto > 2000.0 {
		DescontoIR = 0.2 * Sbruto
	}

	totalDescontos := DescontoINSS + DescontoIR
	Sliquido := Sbruto- totalDescontos 

	fmt.Println("Matrícula: ", matricula)
	fmt.Println("Sálario bruto: ", Sbruto)
	fmt.Println("Desconto INSS e Imposto de Renda: ", DescontoINSS, ",", DescontoIR)
	fmt.Println("Sálario líquido: ", Sliquido)	
}