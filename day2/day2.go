package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

func part1() {
	dat, err := os.Open("input.txt")
	check(err)

	var horizontal int
	var vertical int

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			horizontal = horizontal + val
		case "down":
			vertical = vertical + val
		case "up":
			vertical = vertical - val
		}
	}

	//fmt.Println("horizontal: ", horizontal, "vertical: ", vertical)
	//fmt.Println(horizontal * vertical)
}

func part2() {
	dat, err := os.Open("input.txt")
	check(err)

	var horizontal int
	var vertical int
	var aim int

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		i := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(i[1])
		switch i[0] {
		case "forward":
			horizontal = horizontal + val
			vertical = vertical + (aim * val)
		case "down":
			aim = aim + val
		case "up":
			aim = aim - val
		}
	}

	//fmt.Println("horizontal: ", horizontal, "vertical: ", vertical)
	fmt.Println(horizontal * vertical)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
