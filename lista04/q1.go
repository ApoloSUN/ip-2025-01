package main

import "fmt"

func main() {
	var list[10]int
	var i int 
	var menores = 0 

	for i = 0; i < 10; i = i + 1 {
		fmt.Scanf("%d", &list[i])
		if list[i] > 50 {
			fmt.Println("Número:", list[i], "no índice", i)
		} else {
			menores = menores + 1
		} 
		
	}
	
	if menores == 10 {
		fmt.Println("Não possui nenhum número superior a 50")
    }
}
