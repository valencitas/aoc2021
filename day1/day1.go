package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {

	dat, err := os.Open("input.txt")
	check(err)

	var counter int
	var prevLine string

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if scanner.Text() > prevLine {
			counter++
		}
		prevLine = scanner.Text()
	}

	fmt.Println(counter)
}

func part2() {
	var input []int

	dat, err := os.Open("input.txt")
	check(err)

	// heck it, put it all in memory
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		input = append(input, val)
	}

	var count int
	for i := 3; i < len(input); i++ {
		// why cant you sum an array in go easily
		if (input[i-3] + input[i-2] + input[i-1]) < (input[i-2] + input[i-1] + input[i]) {
			count++
		}
	}

	fmt.Println(count)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
