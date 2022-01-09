package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Run() {

	rawInput, err := ioutil.ReadFile("./day_1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------\n[+] Day 1")
	part1(string(rawInput))
	part2(string(rawInput))
}

func part1(rawInput string) {
	input := strings.Split(rawInput, "\n")

	increases := 1
	for i, v := range input {
		if i != 0 && i-1 < len(input) {
			if v > input[i-1] {
				increases++
			}
		}
	}
	fmt.Println("	[1] Answer:\n		[^] Increases:", increases)
}

func part2(rawInput string) {
	input := strings.Split(rawInput, "\n")
	// increases := 1
	processed := make([]int, 0, 1)

	for i := 0; i < len(input); i++ {
		if i+1 < len(rawInput) && i+2 < len(input) {
			sum := 0
			for j := 0; j < 3; j++ {
				iteration, _ := strconv.Atoi(input[i+j])
				sum += iteration
			}
			processed = append(processed, sum)
		}
	}
	increases := 0
	for i, v := range processed {
		if i != 0 && i-1 < len(processed) {
			if v > processed[i-1] {
				increases++
			}
		}
	}
	fmt.Println("	[2] Answer:\n		[^] Increases:", increases)
}
