package main

import "fmt"

func main() {
	nome := "Ícaro"
	versao := 1.1

	fmt.Println("Olá senhor, ", nome)
	fmt.Println("Versão do programa: ", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O valor do comando é: ", comando)

}
