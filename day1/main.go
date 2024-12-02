package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var first, second []int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, "   ")
		tmp, _ := strconv.Atoi(elements[0])
		first = append(first, tmp)
		tmp, _ = strconv.Atoi(elements[1])
		second = append(second, tmp)
	}

	// step 1
	dist(first, second)

	// step 2
	similarity(first, second)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func dist(first, second []int) {
	sort.Ints(first)
	sort.Ints(second)

	sum := 0
	for i, _ := range first {
		sum += Abs(first[i] - second[i])
	}

	println(sum)
}

func similarity(first, second []int) {
	sum := 0
	for i, _ := range first {
		occ := 0
		for j, _ := range second {
			if first[i] == second[j] {
				occ++
			}
		}
		sum += occ * first[i]
	}

	println(sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
