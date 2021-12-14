package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	ParseInput(input)
}

func ParseInput(input []string) int {
	numIncreasing := 0
	a, err := strconv.Atoi(input[0])
	b, err := strconv.Atoi(input[1])
	c, err := strconv.Atoi(input[2])
	d, err := strconv.Atoi(input[3])
	if err != nil {
		panic(err)
	}
	firstWindow := a + b + c
	secondWindow := firstWindow + d - a

	for currentIndex := 0; currentIndex < len(input)-3; currentIndex++ {
		a, err = strconv.Atoi(input[currentIndex])
		d, err = strconv.Atoi(input[currentIndex+3])
		secondWindow = firstWindow + d - a
		if currentIsIncrease(secondWindow, firstWindow) {
			numIncreasing++
		}
		firstWindow = secondWindow
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(numIncreasing)
	return numIncreasing

}

func currentIsIncrease(currentVal int, previousVal int) bool {
	return currentVal > previousVal
}

func getInput() []string {
	res := []string{}
	file, err := os.Open("/home/windingroad100hf/github/aoc/12-01-2021/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		res = append(res, scanner.Text())
	}
	return res

}
