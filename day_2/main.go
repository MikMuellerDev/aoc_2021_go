package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Run() {

	rawInput, err := ioutil.ReadFile("./day_2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------\n[+] Day 2")
	part1(string(rawInput))
	part2(string(rawInput))
}

func part1(rawInput string) {
	input := strings.Split(rawInput, "\n")
	var posX int
	var posY int

	for _, v := range input {
		instruction := strings.Split(v, " ")[0]
		value, _ := strconv.Atoi(strings.Split(v, " ")[1])
		switch instruction {
		case "up":
			posY -= value
		case "down":
			posY += value
		case "forward":
			posX += value
		}
	}
	fmt.Println("	[1] Answer:\n		[*] Product:", posX*posY)
}

func part2(rawInput string) {
	input := strings.Split(rawInput, "\n")
	var posX int
	var posY int
	var aim int

	for _, v := range input {
		instruction := strings.Split(v, " ")[0]
		value, _ := strconv.Atoi(strings.Split(v, " ")[1])
		switch instruction {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			posX += value
			posY += value * aim
		}
	}
	fmt.Println("	[2] Answer:\n		[*] Product:", posX*posY)
}
