package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int64
	

	fmt.Println("Digite um dia:")
	fmt.Scanln(&dia)

	fmt.Println("Digite um mês:")
	fmt.Scanln(&mes)

	fmt.Println("Digite um ano:")
	fmt.Scanln(&ano)

	if mes == 1 {
		fmt.Printf("%d de janeiro de %d\n", dia, ano)
	
	} else if mes == 2 { 
		fmt.Printf("%d de fevereiro de %d\n", dia, ano)
	
	} else if mes == 3 { 
		fmt.Printf("%d de março de %d\n", dia, ano)

	} else if mes == 4 { 
		fmt.Printf("%d de abril de %d\n", dia, ano)

	} else if mes == 5 { 
		fmt.Printf("%d de maior de %d\n", dia, ano)

	} else if mes == 6 { 
		fmt.Printf("%d de junho de %d\n", dia, ano)

	} else if mes == 7 { 
		fmt.Printf("%d de julho de %d\n", dia, ano)

	} else if mes == 8 { 
		fmt.Printf("%d de agosto de %d\n", dia, ano)

	} else if mes == 9 { 
		fmt.Printf("%d de setembro de %d\n", dia, ano)

	} else if mes == 10 { 
		fmt.Printf("%d de outubro de %d\n", dia, ano)

	} else if mes == 11 { 
		fmt.Printf("%d de novembro de %d\n", dia, ano)
	
	} else { 
		fmt.Printf("%d de dezembro de %d\n", dia, ano)
	}
}