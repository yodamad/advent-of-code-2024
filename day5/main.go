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

const day = "./day5/"
const demo = "demo"
const input = "input"
const debug = "debug"

type Element struct {
	element       int
	before, after []int
}

var elements [1000]Element

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
		// Build rules
		if strings.Contains(line, "|") {
			buildElement(line)
		}
		// Check
		safe := true
		if strings.Contains(line, ",") {
			nbs := strings.Split(line, ",")
			for idx, nb := range nbs {
				before := nbs[:idx]
				iNb, _ := strconv.Atoi(nb)
				safe = safe && elements[iNb].isNotAfter(before)
				after := nbs[idx+1:]
				safe = safe && elements[iNb].isNotBefore(after)
			}
			if safe {
				newOne, _ := strconv.Atoi(nbs[int(len(nbs)/2)])
				sum += newOne
			}
		}
	}
	//fmt.Println(elements)
	fmt.Println(sum)
}

func buildElement(line string) {
	split := strings.Split(line, "|")
	first, _ := strconv.Atoi(split[0])
	second, _ := strconv.Atoi(split[1])
	elements[first].after = append(elements[first].after, second)
	elements[second].before = append(elements[second].before, first)
}

func (e Element) isNotBefore(extract []string) bool {
	for _, ex := range extract {
		nb, _ := strconv.Atoi(ex)
		if slices.Contains(e.before, nb) {
			return false
		}
	}
	return true
}

func (e Element) isNotAfter(extract []string) bool {
	for _, ex := range extract {
		nb, _ := strconv.Atoi(ex)
		if slices.Contains(e.after, nb) {
			return false
		}
	}
	return true
}
