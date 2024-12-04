package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const day = "./day4/"
const demo = "demo"
const input = "input"
const debug = "debug"

const DOWN = "down"
const UP = "up"
const RIGHT = "right"
const LEFT = "left"
const DOWN_RIGHT = "down_right"
const DOWN_LEFT = "down_left"
const UP_RIGHT = "up_right"
const UP_LEFT = "up_left"

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

	var myMap [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		myMap = append(myMap, chars)
	}
	fmt.Println(myMap)

	found := 0
	for l := range myMap {
		// lines
		for c := range myMap[l] {
			// columns
			if myMap[l][c] == "X" {
				for _, way := range findPossibleWays(c, l, "M", myMap) {
					ok, _ := checkNeighbors(c, l, 2, myMap, "A", way)
					if ok {
						ok, _ := checkNeighbors(c, l, 3, myMap, "S", way)
						if ok {
							fmt.Println("Found XMAS at " + "(" + strconv.Itoa(c) + "," + strconv.Itoa(l) + ") in " + way)
							found++
						}
					}
				}
			}
		}
	}
	fmt.Println(found)
}

func findPossibleWays(abs, ord int, neighbor string, fullMap [][]string) []string {
	var ways []string
	if ord > 0 && fullMap[ord-1][abs] == neighbor {
		ways = append(ways, UP)
	}
	if ord < len(fullMap)-1 && fullMap[ord+1][abs] == neighbor {
		ways = append(ways, DOWN)
	}
	if abs < len(fullMap[0])-1 && fullMap[ord][abs+1] == neighbor {
		ways = append(ways, RIGHT)
	}
	if abs > 0 && fullMap[ord][abs-1] == neighbor {
		ways = append(ways, LEFT)
	}
	if ord < len(fullMap)-1 && abs < len(fullMap[0])-1 && fullMap[ord+1][abs+1] == neighbor {
		ways = append(ways, DOWN_RIGHT)
	}
	if ord < len(fullMap)-1 && abs > 0 && fullMap[ord+1][abs-1] == neighbor {
		ways = append(ways, DOWN_LEFT)
	}
	if ord > 0 && abs < len(fullMap[0])-1 && fullMap[ord-1][abs+1] == neighbor {
		ways = append(ways, UP_RIGHT)
	}
	if ord > 0 && abs > 0 && fullMap[ord-1][abs-1] == neighbor {
		ways = append(ways, UP_LEFT)
	}
	return ways
}

func checkNeighbors(abs, ord, step int, fullMap [][]string, neighbor, orientation string) (bool, string) {
	switch orientation {
	case UP:
		return ord >= 0+step && fullMap[ord-step][abs] == neighbor, UP
	case DOWN:
		return ord < len(fullMap)-step && fullMap[ord+step][abs] == neighbor, DOWN
	case RIGHT:
		return abs < len(fullMap[0])-step && fullMap[ord][abs+step] == neighbor, RIGHT
	case LEFT:
		return abs >= 0+step && fullMap[ord][abs-step] == neighbor, LEFT
	case UP_RIGHT:
		return ord >= 0+step && abs < len(fullMap[0])-step && fullMap[ord-step][abs+step] == neighbor, UP_RIGHT
	case UP_LEFT:
		return ord >= 0+step && abs >= 0+step && fullMap[ord-step][abs-step] == neighbor, UP_LEFT
	case DOWN_RIGHT:
		return ord < len(fullMap)-step && abs < len(fullMap[0])-step && fullMap[ord+step][abs+step] == neighbor, DOWN_RIGHT
	case DOWN_LEFT:
		return ord < len(fullMap)-step && abs >= 0+step && fullMap[ord+step][abs-step] == neighbor, DOWN_RIGHT
	default:
		if ord > 0+step && fullMap[ord-step][abs] == neighbor {
			return true, UP
		}
		if ord < len(fullMap)-step && fullMap[ord+step][abs] == neighbor {
			return true, DOWN
		}
		if abs < len(fullMap[0])-step && fullMap[ord][abs+step] == neighbor {
			return true, RIGHT
		}
		if abs > 0+step && fullMap[ord][abs-step] == neighbor {
			return true, LEFT
		}
		if ord > 0+step && abs < len(fullMap[0])-step && fullMap[ord-step][abs+step] == neighbor {
			return true, UP_RIGHT
		}
		if ord > 0+step && abs > 0+step && fullMap[ord-step][abs-step] == neighbor {
			return true, UP_LEFT
		}
		if ord < len(fullMap)-step && abs < len(fullMap[0])-step && fullMap[ord+step][abs+step] == neighbor {
			return true, DOWN_RIGHT
		}
		if ord < len(fullMap)-step && abs > 0+step && fullMap[ord+step][abs-step] == neighbor {
			return true, DOWN_LEFT
		}
		return false, ""
	}
}
