package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseMultiplication(str string) int {
	parts := strings.Split(str, ",")
	// strip first 4 off of [0] to remove mul(
	parts[0] = parts[0][4:]
	// string last off of [1] to remove )
	parts[1] = parts[1][:len(parts[1])-1]
	firstNum, _ := strconv.Atoi(parts[0])
	secondNum, _ := strconv.Atoi(parts[1])
	return firstNum * secondNum
}

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don?[']?t\(\))`)

	// Get all matches
	allMatches := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		allMatches = append(allMatches, matches...)
	}

	totalValue := 0
	shouldAdd := true
	// Parse all multiplication values
	for _, value := range allMatches {

		if value == "do()" {
			shouldAdd = true
		} else if value == "don't()" {
			shouldAdd = false
		} else if shouldAdd {
			product := parseMultiplication(value)
			totalValue += product
		}
	}
	fmt.Printf("Total Product is %v\nFunction took %v\n", totalValue, time.Since(start))
}
