package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

const maiorNumero = 60

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var (
		jogos   int
		dezenas int
	)
	flag.IntVar(&dezenas, "d", 6, "Dezenas")
	flag.IntVar(&jogos, "j", 1, "Jogos")
	flag.Parse()
	for i := 0; i < jogos; i++ {
		numeros, err := gera(dezenas)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(numeros)
	}
}

func gera(n int) ([]int, error) {
	if n < 0 || n > maiorNumero {
		return nil, errors.New(fmt.Sprintf("impossível gerar %v números", n))
	}
	mapa := make(map[int]bool)
	for len(mapa) < n {
		dezena := rand.Intn(maiorNumero) + 1
		mapa[dezena] = true
	}
	numeros := make([]int, 0, n)
	for k := range mapa {
		numeros = append(numeros, k)
	}
	sort.Ints(numeros)
	return numeros, nil
}
