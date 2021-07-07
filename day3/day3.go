package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run(tobogganMap string) {
	tMap := getMap(tobogganMap)
	treeCounter := 0
	slopes := [6][2]int{
		{3, 1},
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for i, slope := range slopes {
		slopeTreeCount := 0
		x := 0
		for y := 0; y < len(tMap); y += slope[1] {
			if tMap[y][x] == "#" {
				slopeTreeCount++
			}
			x += slope[0]
			if x >= len(tMap[y]) {
				x = x % len(tMap[y])
			}
		}
		if i == 0 {
			fmt.Println("Day3 Part I: ", slopeTreeCount)
		} else {
			if treeCounter == 0 {
				treeCounter = slopeTreeCount
			} else {
				treeCounter *= slopeTreeCount
			}
		}
	}
	fmt.Println("Day3 Part II: ", treeCounter)
}

func getMap(iFile string) [][]string {
	var tMap [][]string
	file, err := os.Open(iFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, "")
		var jawn []string
		for _, v := range line_split {
			jawn = append(jawn, string(v))
		}
		tMap = append(tMap, jawn)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tMap
}
