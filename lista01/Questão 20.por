programa {

  funcao inicio() {

    inteiro horas, minutos, segundos, TS

    escreva("Digite as horas: ")
    leia(horas)

    escreva("Digite os minutos: ")
    leia(minutos)

    escreva("Digite os segundos: ")
    leia(segundos)

    TS = (horas * 3600) + (minutos * 60) + segundos

    escreva("O TEMPO EM SEGUNDOS E = ", TS, "\n")
    
  }
}
