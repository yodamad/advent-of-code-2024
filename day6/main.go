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

const day = "./day6/"
const demo = "demo"
const input = "input"
const debug = "debug"

type Point struct {
	x, y int
	way  string
}

var debugPoint = Point{2, 7, ""}
var initPoint, currentPoint Point
var visitedPoints, myMap []Point
var hitWay = ""

const NORTH = "^"
const WEST = ">"
const EAST = "<"
const SOUTH = "v"

var currentWay = NORTH
var journeyLength = 0
var maxY = 0
var maxX = 0

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
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") {
			points := strings.Split(line, "")
			for ptIdx, ptValue := range points {
				if ptValue == "#" {
					myMap = append(myMap, Point{x: ptIdx, y: maxY})
				}
				if ptValue == "^" {
					currentPoint = Point{x: ptIdx, y: maxY}
					initPoint = currentPoint
				}
			}
		}
		maxX = len(line)
		maxY++
	}
	fmt.Println(myMap)
	fmt.Println(currentPoint)

	for nextTurn() {
	}
	fmt.Println(len(visitedPoints))
	fmt.Println(journeyLength)
	fmt.Println(strconv.Itoa(int(time.Now().UnixMilli()-start.UnixMilli())) + "ms")

	// part 2 : mode brut force
	nbBlockingPoints := 0
	for y2 := 0; y2 < maxY; y2++ {
		for x2 := 0; x2 < maxX; x2++ {
			if !slices.Contains(myMap, Point{x2, y2, ""}) {
				currentPoint = initPoint
				extendedMap := slices.Clone(myMap)
				extendedMap = append(extendedMap, Point{x2, y2, ""})
				nbBlockingPoints += isBlockingOne(extendedMap)
			}
		}
	}
	fmt.Println(nbBlockingPoints)
}

func nextTurn() bool {
	switch currentWay {
	case NORTH:
		for i := currentPoint.y; i >= 0; i-- {
			if slices.Contains(myMap, Point{x: currentPoint.x, y: i}) {
				journeyLength += currentPoint.y - i - 1
				currentPoint = Point{x: currentPoint.x, y: i + 1}
				currentWay = WEST
				return true
			} else if !slices.Contains(visitedPoints, Point{x: currentPoint.x, y: i}) {
				visitedPoints = append(visitedPoints, Point{x: currentPoint.x, y: i})
			}
		}
		journeyLength += currentPoint.y
		return false
	case SOUTH:
		for i := currentPoint.y; i < maxY; i++ {
			if slices.Contains(myMap, Point{x: currentPoint.x, y: i}) {
				journeyLength += (i - 1) - currentPoint.y
				currentPoint = Point{x: currentPoint.x, y: i - 1}
				currentWay = EAST
				return true
			} else if !slices.Contains(visitedPoints, Point{x: currentPoint.x, y: i}) {
				visitedPoints = append(visitedPoints, Point{x: currentPoint.x, y: i})
			}
		}
		journeyLength += maxY - currentPoint.y - 1
		return false
	case WEST:
		for i := currentPoint.x; i < maxY; i++ {
			if slices.Contains(myMap, Point{x: i, y: currentPoint.y}) {
				journeyLength += i - 1 - currentPoint.x
				currentPoint = Point{x: i - 1, y: currentPoint.y}
				currentWay = SOUTH
				return true
			} else if !slices.Contains(visitedPoints, Point{x: i, y: currentPoint.y}) {
				visitedPoints = append(visitedPoints, Point{x: i, y: currentPoint.y})
			}
		}
		journeyLength += maxX - currentPoint.x - 1
		return false
	case EAST:
		for i := currentPoint.x; i >= 0; i-- {
			if slices.Contains(myMap, Point{x: i, y: currentPoint.y}) {
				journeyLength += currentPoint.x - i - 1
				currentPoint = Point{x: i + 1, y: currentPoint.y}
				currentWay = NORTH
				return true
			} else if !slices.Contains(visitedPoints, Point{x: i, y: currentPoint.y}) {
				visitedPoints = append(visitedPoints, Point{x: i, y: currentPoint.y})
			}
		}
		journeyLength += currentPoint.x
		return false
	default:
		return false
	}
}

func isBlockingOne(extendedMap []Point) int {
	currentPoint = initPoint
	currentWay = NORTH
	visitedPoints = visitedPoints[:0]
	isBlocking := notBlocking(extendedMap)
	for isBlocking == -1 {
		isBlocking = notBlocking(extendedMap)
	}
	return isBlocking
}

func notBlocking(extendedMap []Point) int {
	switch currentWay {
	case NORTH:
		for i := currentPoint.y - 1; i >= 0; i-- {
			if slices.Contains(extendedMap, Point{currentPoint.x, i, ""}) {
				if slices.Contains(visitedPoints, Point{currentPoint.x, i, NORTH}) {
					return 1
				} else {
					visitedPoints = append(visitedPoints, Point{currentPoint.x, i, NORTH})
					currentPoint = Point{x: currentPoint.x, y: i + 1}
					currentWay = WEST
					return -1
				}
			}
		}
		return 0
	case SOUTH:
		for i := currentPoint.y + 1; i < maxY; i++ {
			if slices.Contains(extendedMap, Point{currentPoint.x, i, ""}) {
				if slices.Contains(visitedPoints, Point{currentPoint.x, i, SOUTH}) {
					return 1
				} else {
					visitedPoints = append(visitedPoints, Point{currentPoint.x, i, SOUTH})
					currentPoint = Point{x: currentPoint.x, y: i - 1}
					currentWay = EAST
					return -1
				}
			}
		}
		return 0
	case WEST:
		for i := currentPoint.x + 1; i < maxX; i++ {
			if slices.Contains(extendedMap, Point{i, currentPoint.y, ""}) {
				if slices.Contains(visitedPoints, Point{i, currentPoint.y, WEST}) {
					return 1
				} else {
					visitedPoints = append(visitedPoints, Point{i, currentPoint.y, WEST})
					currentPoint = Point{x: i - 1, y: currentPoint.y}
					currentWay = SOUTH
					return -1
				}
			}
		}
		return 0
	case EAST:
		for i := currentPoint.x - 1; i >= 0; i-- {
			if slices.Contains(extendedMap, Point{i, currentPoint.y, ""}) {
				if slices.Contains(visitedPoints, Point{i, currentPoint.y, EAST}) {
					return 1
				} else {
					visitedPoints = append(visitedPoints, Point{i, currentPoint.y, EAST})
					currentPoint = Point{x: i + 1, y: currentPoint.y}
					currentWay = NORTH
					return -1
				}
			}
		}
		return 0
	default:
		return 0
	}
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y && p.way == other.way
}
