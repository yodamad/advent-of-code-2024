package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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

	start := time.Now()
	scanner := bufio.NewScanner(file)
	sum := 0
	sumReorder := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Build rules
		if strings.Contains(line, "|") {
			buildElement(line)
		}
		// Check
		if strings.Contains(line, ",") {

			nbs := strings.Split(line, ",")
			safe := isSafe(nbs)
			if safe {
				newOne, _ := strconv.Atoi(nbs[(len(nbs) / 2)])
				sum += newOne
			} else {
				for !safe {
					for idx, nb := range nbs {
						if safe {
							break
						} else {
							iNb, _ := strconv.Atoi(nb)
							before := nbs[:idx]
							for _, b := range before {
								idxFound := elements[iNb].findMisplaced(nbs, b, false)
								if idxFound != -1 {
									backup := nbs[idxFound]
									nbs[idxFound] = nbs[idx]
									nbs[idx] = backup
									break
								}
							}
							safe = isSafe(nbs)
							if !safe {
								after := nbs[idx+1:]
								for _, b := range after {
									idxFound := elements[iNb].findMisplaced(nbs, b, true)
									if idxFound != -1 {
										backup := nbs[idxFound]
										nbs[idxFound] = nbs[idx]
										nbs[idx] = backup
										break
									}
								}
							}
							safe = isSafe(nbs)
						}
					}
				}
				newOne, _ := strconv.Atoi(nbs[(len(nbs) / 2)])
				sumReorder += newOne
			}
		}
	}
	//fmt.Println(elements)
	fmt.Println(sum)
	fmt.Println(sumReorder)
	fmt.Println(time.Now().UnixMilli() - start.UnixMilli())
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

func isSafe(nbs []string) bool {
	safe := true
	for idx, nb := range nbs {
		before := nbs[:idx]
		iNb, _ := strconv.Atoi(nb)
		safe = safe && elements[iNb].isNotAfter(before)
		after := nbs[idx+1:]
		safe = safe && elements[iNb].isNotBefore(after)
	}
	return safe
}

func (e Element) findMisplaced(nbs []string, element string, checkBefore bool) int {
	idx := -1
	nb, _ := strconv.Atoi(element)
	toAnalyze := e.after
	if checkBefore {
		toAnalyze = e.before
	}
	if slices.Contains(toAnalyze, nb) {
		idx = slices.Index(nbs, element)
	}
	return idx
}
