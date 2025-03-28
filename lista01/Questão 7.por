programa {

  funcao inicio() {

    real f, c, p, ml

    escreva("Digite um valor em fahrenheit: ")
    leia(f)

    escreva("Digite um valor em polegadas: ")
    leia(p)

    c = (5 * f - 160) / 9

    ml = p * 25.4

    escreva("O VALOR EM CELSIUS = ", c, "\n")
    escreva("A QUANTIDADE DE CHUVA E = ", ml, "\n")
    
  }
}
