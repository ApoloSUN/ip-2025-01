programa {

  funcao inicio() {

    real r, h, area, custo, pi
    
    pi = 3.14159

    escreva("Digite o valor do raio e da altura da lata:")
    leia(r)
    leia(h)

    area = (2 * pi * (r * r)) + (2 * pi * r * h)

    custo = area * 100.00

    escreva("O VALOR DO CUSTO E = ", custo, "\n")
    
  }
}
