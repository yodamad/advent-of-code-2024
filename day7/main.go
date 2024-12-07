package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const day = "./day7/"
const demo = "demo"
const input = "input"
const debug = "debug"

func main() {
	file, err := os.Open(day + input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		total, _ := strconv.Atoi(strings.Split(line, ":")[0])
		sNbs := strings.Split(strings.TrimPrefix(strings.Split(line, ":")[1], " "), " ")
		var nbs []int

		for _, sNb := range sNbs {
			nbs = append(nbs, toInt(sNb))
		}

		// part 1
		var results []int
		computeCombinations(nbs, 1, nbs[0], &results)
		if slices.Contains(results, total) {
			sum += total
		}
	}
	fmt.Println("part1 is ")
	fmt.Println(sum)
	fmt.Println("==============")
}

func computeCombinations(arr []int, index int, acc int, results *[]int) {
	if index == len(arr) {
		*results = append(*results, acc)
		return
	}
	// Sum
	computeCombinations(arr, index+1, acc+arr[index], results)
	// Multiplication
	computeCombinations(arr, index+1, acc*arr[index], results)
	// Concatenation
	computeCombinations(arr, index+1, toInt(strconv.Itoa(acc)+strconv.Itoa(arr[index])), results)

}

func toInt(s string) int {
	nb, _ := strconv.Atoi(s)
	return nb
}
