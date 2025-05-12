package main

import (
	"fmt"
	"math"
)

func main() {
	var figura string 
	var raio, altura float64
	
	fmt.Println("Digite a figura, o raio e a altura desejada: ")
	fmt.Scanln(&figura, &raio, &altura)

	pi := 3.14

	if figura == "1" {
		Volume := (pi * (raio * raio) * altura) / 3
		Area := pi * raio * math.Sqrt((raio * raio) + (altura * altura))
		fmt.Println("Volume =", Volume, "e Area =", Area)
	} else if figura == "2" {
		Volume := (pi * (raio * raio) * altura)
		Area := 2 * pi * raio * altura
		fmt.Println("Volume =", Volume, "e Area =", Area)
	} else {
		Volume := (4 / 3) * pi * (raio * raio * raio) 
		Area := 4 * pi * (raio * raio)
		fmt.Println("Volume =", Volume, "e Area =", Area)
	}
}