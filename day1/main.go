package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	leftLocationIds := []int{}
	rightLocationIds := []int{}

	for scanner.Scan() {
		numsArr := strings.Split(scanner.Text(), "   ")
		leftNum, _ := strconv.Atoi(numsArr[0])
		rightNum, _ := strconv.Atoi(numsArr[1])

		leftLocationIds = append(leftLocationIds, leftNum)
		rightLocationIds = append(rightLocationIds, rightNum)
	}
	sort.Ints(leftLocationIds)
	sort.Ints(rightLocationIds)

	totalDiff := 0.0
	totalSimilarity := 0
	dict := make(map[int]int)

	for _, value := range rightLocationIds {
		dict[value] += 1
	}

	for index, value := range leftLocationIds {
		totalDiff += math.Abs(float64(value - rightLocationIds[index]))
		totalSimilarity += value * dict[value]
	}
	fmt.Printf("Total Difference %f\nTotal Similarity is %v\nFunction took %v time\n", totalDiff, totalSimilarity, time.Since(start))
}
