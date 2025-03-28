programa {

  funcao inicio() {

    real salario, SR

    escreva("Digite o valor do salário: ")
    leia(salario)

    se (salario <= 300) {
        SR = salario * 1.50 
    } senao {
        SR = salario * 1.30  
    }

    escreva("SALARIO COM REAJUSTE = ", SR, "\n")
    
  }
}
