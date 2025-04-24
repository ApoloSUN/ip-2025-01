package main

import "fmt"

func main() {
	var list1[10]int
	var list2[5]int
	var i, a, c int 
	var vetorP []int
	var vetorI []int

	for i = 0; i < 10; i = i + 1 {
		fmt.Scan(&list1[i])
	}

	somalist2 := 0 

	for a = 0; a < 5; a = a + 1 {
		fmt.Scan(&list2[a])
		somalist2 = somalist2 + list2[a] 
	}

	for c = 0; c < 10; c = c + 1 {
		if list1[c] % 2 == 0 {
			vetorP = append(vetorP, list1[c] + somalist2)
		} else {
			vetorI = append(vetorI, list1[c] + somalist2)
		}
	}

	fmt.Println("Vetor par:", vetorP, "Vetor Ímpar:", vetorI)		
}