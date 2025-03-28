programa {

  funcao inicio() {

    inteiro x, y, i

    escreva("Digite dois valores: ")
    leia(x)
    leia(y)

    se (x % 2 == 0) {
        para (i = 0; i < y; i = i + 1) {
            escreva(x + (i * 2))
            se (i < y - 1) {
                escreva(" ")
            }
        }
      
    } senao {
        escreva("O PRIMEIRO NUMERO NAO E PAR\n")
    }
    
  }
}
