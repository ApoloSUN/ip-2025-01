programa {

  funcao inicio() {

    inteiro n, i
    real f, c

    escreva("Digite o números de temperaturas: ")
    leia(n)

    para (i = 1; i <= n; i = i + 1) {
        escreva("Digite um valor de temperatura em fahrenheit: ")
        leia(f)
        
        c = 5 * (f - 32) / 9

    }
    escreva(f, " FAHRENHEIT EQUIVALE A ", c, " CELSIUS\n")    
  }
}
