programa {

  funcao inicio() {

    inteiro a1, r, n, soma, termo, i

    escreva("Digite o valor inicial da progressão, a razão e o número de elementos: ")
    leia(a1)
    leia(r) 
    leia(n)

    soma = 0

    para (i = 0; i < n; i = i + 1) {
        termo = a1 + (i * r)  
        soma = soma + termo   
    }

    escreva(soma, "\n")
    
  }
}
