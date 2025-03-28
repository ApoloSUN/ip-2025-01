programa {

  funcao inicio() {

    inteiro numero

    escreva("Digite um número: ")
    leia(numero)

    se ((numero % 3 == 0) e (numero % 5 == 0)) {
        escreva("O NUMERO E DIVISIVEL", "\n")
    
    } senao {
        escreva("O NUMERO NAO E DIVISIVEL", "\n")
    }
    
  }
}
