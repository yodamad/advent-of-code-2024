package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const demo = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const demoPart2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))\n"

const regex_1 = "mul\\([0-9]*,[0-9]*\\)"
const regex_do = "do\\(\\)"
const regex_dont = "don't\\(\\)"

func main() {
	testDemo := false
	pattern := regexp.MustCompile(regex_1)
	patterDo := regexp.MustCompile(regex_do)
	patterDont := regexp.MustCompile(regex_dont)

	if testDemo {

		idxDo := patterDo.FindAllStringIndex(demoPart2, -1)
		idxDont := patterDont.FindAllStringIndex(demoPart2, -1)

		fmt.Println(idxDo)
		fmt.Println(idxDont)

		var dos, donts []int
		for d := range idxDo {
			dos = append(dos, idxDo[d][0])
		}
		for d := range idxDont {
			donts = append(donts, idxDont[d][0])
		}

		sum := 0
		matches := pattern.FindAllString(demo, -1)
		matchesIdx := pattern.FindAllStringIndex(demo, -1)
		for idx, match := range matchesIdx {
			start := match[0]
			do := findClosest(dos, start, 0)
			dont := findClosest(donts, start, -1)

			if dont < do {
				extractNbs := regexp.MustCompile(`\d+`)
				nbs := extractNbs.FindAllString(matches[idx], -1)
				a, _ := strconv.Atoi(nbs[0])
				b, _ := strconv.Atoi(nbs[1])
				sum += a * b
			}
		}
		fmt.Println(sum)
	} else {
		// Read input
		file, err := os.Open("./day3/input")
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		sum := 0
		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			/* Part1
			matches := pattern.FindAllString(scanner.Text(), -1)
			for _, mul := range matches {
				extractNbs := regexp.MustCompile(`\d+`)
				nbs := extractNbs.FindAllString(mul, -1)
				a, _ := strconv.Atoi(nbs[0])
				b, _ := strconv.Atoi(nbs[1])
				sum += a * b
			}
			*/
			line := scanner.Text()
			idxDo := patterDo.FindAllStringIndex(line, -1)
			idxDont := patterDont.FindAllStringIndex(line, -1)

			fmt.Println(idxDo)
			fmt.Println(idxDont)

			var dos, donts []int
			for d := range idxDo {
				dos = append(dos, idxDo[d][0])
			}
			for d := range idxDont {
				donts = append(donts, idxDont[d][0])
			}
			matches := pattern.FindAllString(line, -1)
			matchesIdx := pattern.FindAllStringIndex(line, -1)
			for idx, match := range matchesIdx {
				start := match[0]
				do := findClosest(dos, start, 0)
				dont := findClosest(donts, start, -1)

				if dont < do {
					fmt.Println(matches[idx])
					extractNbs := regexp.MustCompile(`\d+`)
					nbs := extractNbs.FindAllString(matches[idx], -1)
					a, _ := strconv.Atoi(nbs[0])
					b, _ := strconv.Atoi(nbs[1])
					sum += a * b
				}
			}
		}
		fmt.Println(sum)
	}
}

func findClosest(input []int, search, defaultValue int) int {
	delta := 99999
	res := defaultValue

	for _, elem := range input {
		if (search - elem) < 0 {
			// too far
			break
		} else if (search - elem) < delta {
			delta = search - elem
			res = elem
		}
	}
	return res
}
