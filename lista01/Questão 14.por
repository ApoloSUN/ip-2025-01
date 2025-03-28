programa {

  funcao inicio() {

    real h, a, ab, v

    escreva("Digite a altura da pirâmide: ")
    leia(h)
    escreva("Digite o valor de uma aresta do hexágono da base da pirâmide: ")
    leia(a)

    ab = (3 * (a * a) * (1.732 / 2))

    v = (1 / 3) * ab * h

    escreva("O VOLUME DA PIRAMIDE E = ", v, " METROS CUBICOS", "\n")
    
  }
}
