programa {

  funcao inicio() {

    inteiro n, i
    real soma

    escreva("Digite um número: ")
    leia(n)

    se (n <= 1) {
        escreva("Numero invalido!\n")
    
    } senao {
        soma = 0
        para (i = 1; i <= n; i = i + 1) {
            soma = soma + (1 / i)  
        }
    }
      
    escreva(soma)
    
  }
}
