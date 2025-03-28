programa {

  funcao inicio() {

    inteiro horas, valor, H3

    escreva("Digite a quantidades de horas em que o veículo foi utilizado: ")
    leia(horas)

    H3 = horas / 3

    se (horas % 3 == 0) {
        valor = 10 * H3
    } 
    senao se (horas % 3 == 2) {
        valor =  (10 * H3) + 10 
    }
    senao {
         valor = 10 * H3 + 5
    }

    escreva("O VALOR A PAGAR E = ", valor, "\n")
    
  }
}
