package main

import (
	"fmt"
	"os"
)

func main() {

	showInitial()
	showMenu()

	comando := readyCmd()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func showInitial() {
	nome := "Ícaro"
	versao := 1.1
	fmt.Println("Olá senhor, ", nome)
	fmt.Println("Versão do programa: ", versao)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readyCmd() int {
	var cmd int
	fmt.Scan(&cmd)

	fmt.Println("O valor do comando é: ", cmd)

	return cmd
}
