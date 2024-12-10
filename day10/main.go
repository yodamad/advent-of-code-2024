package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const day = "./day10/"
const demo = "demo"
const input = "input"
const debug = "debug"

var myMap [][]int
var sum = 0

var visited []NinePoint

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

	var zeros []Point
	var currentY = 0
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "")
		for i, p := range points {
			tmpMap := make([]int, len(line))
			myMap = append(myMap, tmpMap)
			myMap[i][currentY] = toInt(p)
			if toInt(p) == 0 {
				zeros = append(zeros, Point{x: i, y: currentY})
			}
		}
		currentY++
	}
	fmt.Println(zeros)

	for _, z := range zeros {
		findPath(z, z)
	}
	fmt.Println(sum)
}

func findPath(p, zero Point) bool {
	//fmt.Println(p)
	if p.value == 9 && !alreadyVisited(p, zero) {
		nine := NinePoint{p, zero}
		visited = append(visited, nine)
		sum++
		return false
	}

	if p.x > 0 && myMap[p.x-1][p.y] == p.value+1 {
		findPath(Point{p.x - 1, p.y, p.value + 1}, zero)
	}
	if p.x < len(myMap)-1 && myMap[p.x+1][p.y] == p.value+1 {
		findPath(Point{p.x + 1, p.y, p.value + 1}, zero)
	}
	if p.y > 0 && myMap[p.x][p.y-1] == p.value+1 {
		findPath(Point{p.x, p.y - 1, p.value + 1}, zero)
	}
	if p.y < len(myMap[p.x])-1 && myMap[p.x][p.y+1] == p.value+1 {
		findPath(Point{p.x, p.y + 1, p.value + 1}, zero)
	}
	return false
}

type Point struct {
	x, y, value int
}

type NinePoint struct {
	Point
	zero Point
}

func alreadyVisited(point, zero Point) bool {
	for _, p := range visited {
		if p.x == point.x && p.y == point.y && p.zero.x == zero.x && p.zero.y == zero.y {
			return true
		}
	}
	return false
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y && p.value == other.value
}

func toInt(s string) int {
	nb, _ := strconv.Atoi(s)
	return nb
}
