package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			printLogs()
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

	sites := readSitesFile()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorre um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFile() []string {

	var sites []string

	//io.util - lê tudo de uma vez
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorrreu um erro", err)
	}

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	read := bufio.NewReader(file)

	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)

	}

	file.WriteString(time.Now().Format("02/01/2006 14:04:05 ") + site + " - " +
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))

}
