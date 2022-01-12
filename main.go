package main

import (
	"fmt"
	"time"

	day1 "github.com/MikMuellerDev/aoc_2021_go/day_1"
	day2 "github.com/MikMuellerDev/aoc_2021_go/day_2"
)

func main() {
	start := time.Now()
	day1.Run()
	day2.Run()
	fmt.Println("Execution time:", time.Since(start))
}
