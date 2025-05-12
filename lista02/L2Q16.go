package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	
	fmt.Println("Digite o valor do coeficiente A:")
	fmt.Scanln(&a)

	fmt.Println("Digite o valor do coeficiente B:")
	fmt.Scanln(&b)

	fmt.Println("Digite o valor do coeficiente C:")
	fmt.Scanln(&c)

	if a == 0 {
		fmt.Println("Isso não é uma equação de segundo grau")
		return
	}
	
	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("Raízes imaginárias")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		fmt.Println(raiz, "Raiz única")
	} else {
		raiz1 := (-b + math.Sqrt(delta)) / (2 * a)
		raiz2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println(raiz1, ",", raiz2, ",", "Raízes distintas")
	}
}