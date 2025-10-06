package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var taxas = map[string]float64{
	"USD": 0.18,
	"EUR": 0.17,
	"GBP": 0.14,
	"ARS": 31.50,
	"CLP": 164.00,
	"JPY": 28.00,
	"CNY": 1.30,
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Moedas suportadas:")

	for moeda, taxa := range taxas {
		fmt.Printf("Código: %s\tTaxa:%.2f\n", moeda, taxa)
	}
	fmt.Printf("\n")
	fmt.Printf("Digite o Código da moeda desejado: ")

	reader.Scan()

	codigo := strings.ToUpper(strings.TrimSpace(reader.Text()))

	taxa, ok := taxas[codigo]

	if !ok {
		fmt.Fprintln(os.Stderr, "Moeda não suportada!")
		os.Exit(1)
	}

	fmt.Printf("Digite o valor em R$: ")
	reader.Scan()

	in := strings.ReplaceAll(strings.TrimSpace(reader.Text()), ",", ".")

	valorBRL, err := strconv.ParseFloat(in, 64)

	if err != nil || valorBRL < 0 {
		fmt.Fprintln(os.Stderr, "Valor Inválido!")
		os.Exit(1)
	}

	resultado := valorBRL * taxa

	fmt.Printf("BRL: %.2f => %.2f :%s\n", valorBRL, resultado, codigo)
}
