package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const debug = false

func main() {
	file, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var elements []string
	var input [][]int
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		elements = strings.Split(line, " ")
		var intElements []int
		for i, _ := range elements {
			convInt, _ := strconv.Atoi(elements[i])
			intElements = append(intElements, convInt)
		}
		input = append(input, intElements)
	}

	sum := 0
	for k := range input {
		dists := computeDists(input[k])
		safe := false
		if checkValid(dists) {
			sum++
			safe = true
		} else {
			for m := range len(input[k]) {
				subInput := removeLevel(slices.Clone(input[k]), m)
				if checkValid(computeDists(subInput)) {
					sum++
					safe = true
					break
				}
			}
		}
		println(strconv.Itoa(k+1) + " is " + strconv.FormatBool(safe))
	}
	println(sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeDists(input []int) []int {
	var dists []int
	for m := range input {
		if m < len(input)-1 {
			dists = append(dists, input[m]-input[m+1])
		}
	}
	return dists
}

func checkValid(dists []int) bool {
	if slices.Min(dists) > 0 && slices.Max(dists) < 4 {
		return true
	} else if slices.Min(dists) > -4 && slices.Max(dists) < 0 {
		return true
	} else {
		return false
	}
}

func removeLevel(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
