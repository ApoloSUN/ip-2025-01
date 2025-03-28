programa {

  funcao inicio() {

    inteiro N, i, Q

    escreva("Digite um valor maior que 5 e menor que 2000: ")
    leia(N)

    para (i = 2; i <= N; i = i + 2) {
        Q = i * i
        escreva(i, "^2 = ", Q, "\n")
    }
  }
}
