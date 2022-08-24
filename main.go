package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("#############################################")
	fmt.Println("###          SEQUÊNCIA GERADA:          #####")
	fmt.Println("#############################################")
	fmt.Println("###     ", gerarSequencia(), "     ###")
	fmt.Println("###     ", gerarSequencia(), "     ###")
	fmt.Println("###     ", gerarSequencia(), "     ###")
	fmt.Println("#############################################")
	fmt.Println("###            BOA SORTE!!!:            #####")
	fmt.Println("#############################################")
}

func gerarSequencia() string {
	var lista []string
	for i := 0; i < 6; i++ {

		gerado := sorteiaNumero()
		for existeNaLista(lista, gerado) {
			gerado = sorteiaNumero()
		}

		geradoString := strconv.Itoa(gerado)
		if gerado < 10 {
			geradoString = "0" + geradoString
		}

		lista = append(lista, geradoString)
	}

	return strings.Join(lista, " - ")
}

func existeNaLista(lista []string, numero int) bool {
	numeroString := strconv.Itoa(numero)
	for _, v := range lista {
		if v == numeroString {
			return true
		}
	}

	return false
}

func sorteiaNumero() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	gerado := r1.Intn(60)
	if gerado == 0 {
		gerado = r1.Intn(60)
	}
	return gerado
}
