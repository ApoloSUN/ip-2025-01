programa {

  funcao inicio() {

    real a, b, c, d, determinante

    escreva("Digite os quatro elementos (A, B, C e D) da matriz: ")
    leia(a)
    leia(b)
    leia(c)
    leia(d)

    determinante = a * d - b * c

    escreva("O VALOR DO DETERMINANTE E = ", determinante, "\n")
    
  }
}
