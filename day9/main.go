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

const day = "./day9/"
const demo = "demo"
const input = "input"
const debug = "debug"

var disk []string

func main() {
	var demo = false
	var line string
	if demo {
		line = "2333133121414131402"
	} else {
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
		for scanner.Scan() {
			line = scanner.Text()
		}
	}
	createVisu(line)
	fmt.Println(strings.Join(disk, ""))
	//diskInfo := compact()
	//fmt.Println(compute(diskInfo))
	blocks := defragment()
	computeFragments(blocks)
}

func createVisu(line string) {
	isFile := true
	currentId := 0
	for _, c := range strings.Split(line, "") {
		nb, _ := strconv.Atoi(c)
		for range nb {
			if isFile {
				disk = append(disk, strconv.Itoa(currentId))
			} else {
				disk = append(disk, ".")
			}
		}
		if isFile {
			currentId++
		}
		isFile = !isFile
	}
}

func compact() []string {
	nbs := findNumbersIndex(disk)
	slices.Sort(nbs)
	slices.Reverse(nbs)
	for i := range disk {
		if disk[i] == "." && len(nbs) > 0 && i <= nbs[0] {
			disk[i] = disk[nbs[0]]
			disk[nbs[0]] = "_"
			nbs = nbs[1:]
		}
	}
	return disk
}

func defragment() []Package {
	packages := findPackages(disk)
	emptyBlocks := slices.Clone(packages)
	slices.Reverse(packages)
	for _, p := range packages {
		if p.isFile {
			idx, remainingSize := findEmptyBlock(p.size, emptyBlocks)
			if idx > 0 {
				if remainingSize > 0 {
					endBlocks := slices.Clone(emptyBlocks[idx+1:])
					emptyBlocks = append(emptyBlocks[0:idx], p)
					emptyBlocks = append(emptyBlocks, Package{".", remainingSize, false})
					emptyBlocks = append(emptyBlocks, endBlocks...)
				} else {
					emptyBlocks[idx] = p
				}
				removeOldPackage(p, emptyBlocks)
			}
		}
	}
	var log string
	for i := range emptyBlocks {
		for range emptyBlocks[i].size {
			log += fmt.Sprint(emptyBlocks[i].index)
		}
	}
	fmt.Println(log)
	return emptyBlocks
}

func findNumbersIndex(chars []string) []int {
	var indexes []int
	for i := range chars {
		if chars[i] != "." {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

type Package struct {
	index  string
	size   int
	isFile bool
}

func (p Package) toString() string {
	return p.index
}

func findPackages(chars []string) []Package {
	previous := chars[0]
	currentPackageSize := 1
	var packages []Package
	nextChars := chars[1:]
	for i := range nextChars {
		if nextChars[i] == previous {
			currentPackageSize++
		} else {
			if previous != "." {
				packages = append(packages, Package{previous, currentPackageSize, true})
			} else {
				packages = append(packages, Package{previous, currentPackageSize, false})
			}
			currentPackageSize = 1
		}
		previous = nextChars[i]
	}
	packages = append(packages, Package{previous, currentPackageSize, true})
	return packages
}

func findEmptyBlock(size int, blocks []Package) (int, int) {
	for idx, b := range blocks {
		if b.size >= size && b.index == "." {
			return idx, b.size - size
		}
	}
	return -1, -1
}

func removeOldPackage(oldBlock Package, blocks []Package) {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].index == oldBlock.index {
			blocks[i].index = "."
			return
		}
	}
}

func compute(diskInfo []string) int {
	sum := 0
	for idx, di := range diskInfo {
		tmp, _ := strconv.Atoi(di)
		sum += tmp * idx
	}
	return sum
}

func computeFragments(diskInfo []Package) {
	var fragments []string
	for _, di := range diskInfo {
		for range di.size {
			fragments = append(fragments, di.index)
		}
	}
	fmt.Println(fragments)
	fmt.Println(compute(fragments))
}

func toInt(s string) int {
	nb, _ := strconv.Atoi(s)
	return nb
}
