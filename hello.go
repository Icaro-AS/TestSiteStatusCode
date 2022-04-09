package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	showInitial()

	for {
		showMenu()
		comando := readyCmd()

		switch comando {
		case 1:
			startMonitor()
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

func startMonitor() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, v := range sites {
			fmt.Println("Testando site", i, ":", v)

			testSites(v)
		}
		time.Sleep(delay * time.Second)

		fmt.Println("")
	}
	fmt.Println("")

}

func testSites(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
