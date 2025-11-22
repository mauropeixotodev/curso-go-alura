package main

import (
	"fmt"
	"reflect"
)

func main() {

	nome := "Mauro"
	idade := 21
	versao := 1.1
	fmt.Println("Ola, sr.", nome, "Sua idade e", idade)
	fmt.Println("A versao desse software e", versao)
	fmt.Println("o tipo da variavel e:", reflect.TypeOf(nome))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comendo é", &comando)
	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {
		fmt.Print("Monitorando")
	} else if comando == 2 {
		fmt.Print("Exibindo logs")
	} else if comando == 0 {
		fmt.Print("Saindo do programa")
	} else {
		fmt.Print("Não conheço esse comando")
	}

}
