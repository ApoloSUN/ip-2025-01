programa {
  funcao inicio() {

    inteiro T, i, Pe, valor 
    real P, G, A, C, 

    escreva("Digite o número de casos testes: ")
    leia (T)

    para (i = 0; i <= T; i = i + 1) {
      escreva("Digite o número de pessoas: ")
      leia (Pe)

      escreva("Digite a percentagem comprados na categoria popular: ")
      leia (P)

      escreva("Digite a percentagem comprados na categoria geral: ")
      leia (G)

      escreva("Digite a percentagem comprados na categoria arquibancada: ")
      leia (A)

      escreva("Digite a percentagem comprados na categoria cadeiras: ")
      leia (C)

      valor = ((Pe * P) * 1) + ((Pe * G) * 5) + ((Pe * A) * 10) + ((Pe * C) * 20)

      escreva("A RENDA DO JOGO N. ", i, "E = ", valor)

    }

  }
}
