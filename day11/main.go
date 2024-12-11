package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nbStones = 0
var resultPart2 = 224577979481346

func main() {
	input := "5178527 8525 22 376299 3 69312 0 275"
	//input := "125 17"
	nbTurn := 10000
	inputs := strings.Split(input, " ")
	for _, i := range inputs {
		applyRulesPart1(toInt(i), nbTurn)
	}
	fmt.Println("-- Part 1 --")
	fmt.Println(nbStones)
	fmt.Println("-- Part 2 --")
	var cache = make(map[int]int)
	for _, nb := range inputs {
		cache[toInt(nb)] = 1
	}
	for range nbTurn {
		// fmt.Println(cache)
		cache = applyRulesPart2(cache)
	}
	sum := 0
	for _, v := range cache {
		sum += v
	}
	fmt.Println(sum)
}

func applyRulesPart1(nb, turn int) {
	if turn == 0 {
		nbStones++
		return
	}
	if nb == 0 {
		applyRulesPart1(1, turn-1)
	} else if len(strconv.Itoa(nb))%2 == 0 {
		str := strconv.Itoa(nb)
		mid := len(str) / 2
		nextTurn := turn - 1
		applyRulesPart1(toInt(str[:mid]), nextTurn)
		applyRulesPart1(toInt(str[mid:]), nextTurn)
	} else {
		applyRulesPart1(nb*2024, turn-1)
	}
}

func applyRulesPart2(stones map[int]int) map[int]int {
	var updatedStones = map[int]int{}

	update := func(stone, incr int) {
		if _, v := updatedStones[stone]; !v {
			updatedStones[stone] = 0
		}
		updatedStones[stone] += incr
	}

	for nb, count := range stones {
		if nb == 0 {
			update(1, count)
		} else if len(strconv.Itoa(nb))%2 == 0 {
			str := strconv.Itoa(nb)
			mid := len(str) / 2
			left := toInt(str[:mid])
			right := toInt(str[mid:])
			update(left, count)
			update(right, count)
		} else {
			update(nb*2024, count)
		}
	}
	return updatedStones
}

func toInt(s string) int {
	nb, _ := strconv.Atoi(s)
	return nb
}
