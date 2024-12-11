package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nbStones = 0

func main() {
	input := "5178527 8525 22 376299 3 69312 0 275"
	// input := "125 17"
	nbTurn := 75
	inputs := strings.Split(input, " ")
	for _, i := range inputs {
		applyRules(toInt(i), nbTurn)
	}
	fmt.Println("----")
	fmt.Println(nbStones)
}

func applyRules(nb, turn int) {
	if turn == 0 {
		nbStones++
		return
	}
	if nb == 0 {
		applyRules(1, turn-1)
	} else if len(strconv.Itoa(nb))%2 == 0 {
		str := strconv.Itoa(nb)
		mid := len(str) / 2
		nextTurn := turn - 1
		applyRules(toInt(str[:mid]), nextTurn)
		applyRules(toInt(str[mid:]), nextTurn)
	} else {
		applyRules(nb*2024, turn-1)
	}
}

func toInt(s string) int {
	nb, _ := strconv.Atoi(s)
	return nb
}
