package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const day = "./day8/"
const demo = "demo"
const input = "input"
const debug = "debug"

type Point struct {
	x, y int
}

var lineId, lineLen, sumAntinodes int
var antinodes []Point

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
	lineId = 0
	sumAntinodes = 0
	myMap := make(map[string][]Point)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen = len(line)
		pts := strings.Split(line, "")
		for i, v := range pts {
			if v != "." {
				myMap[v] = append(myMap[v], Point{i, lineId})
			}
		}
		lineId++
	}
	fmt.Println(myMap)

	for _, v := range myMap {
		for p, pv := range v {
			computeAntinode(pv, v[0:p])
			computeAntinode(pv, v[p+1:])
		}
	}
	fmt.Println(len(antinodes))
}

func computeAntinode(p Point, pts []Point) {
	for _, v := range pts {
		newX := p.x - (v.x - p.x)
		newY := p.y - (v.y - p.y)
		//fmt.Println("(" + strconv.Itoa(newX) + "," + strconv.Itoa(newY) + ")")
		if newX >= 0 && newX < lineLen && newY >= 0 && newY < lineId && !slices.Contains(antinodes, Point{newX, newY}) {
			antinodes = append(antinodes, Point{newX, newY})
		}
	}
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}
