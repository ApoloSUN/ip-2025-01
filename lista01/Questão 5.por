programa {

  funcao inicio() {
    
    inteiro conta
    real consumo, valorconta
    caracter tipo

    escreva("Digite o número da conta, o consumo e o tipo de consumidor: ")
    leia(conta)
    leia(consumo) 
    leia(tipo)

    valorconta = 0

    se (tipo == 'R') {
        valorconta = 5.00 + 0.05 * consumo
    }

    se (tipo == 'C') {
        se (consumo <= 80) { 
            valorconta = 500.00

        } senao {
            valorconta = 500.00 + 0.25 * (consumo - 80)
        }
    }
      
    se (tipo == 'I') {
        se (consumo <= 100) {
            valorconta = 800.00
        } senao {
            valorconta = 800.00 + 0.04 * (consumo - 100)
        }
    }

    escreva("CONTA = ", conta, "\n")
    escreva("VALOR DA CONTA = ", valorconta, "\n")
  }
}
