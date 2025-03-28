programa {
  funcao inicio() {

    real salario, Gkw, Ckw, CC, CD

    escreva("Digite o valor do salário mínimo e a quantidade de kW gastos: ")
    leia(salario)
    leia(Gkw)

    Ckw = salario * 0.70 / 100

    CC = Ckw * Gkw

    CD = CC * 0.90

    escreva("Custo por kW: R$ ", Ckw, "\n")
    escreva("Custo do consumo: R$ ", CC, "\n")
    escreva("Custo com desconto: R$ ", CD, "\n")

    
  }
}
