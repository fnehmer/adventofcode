package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Run on input
func readInput() int {
	file, err := os.Open("./input")
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sum += add_line(line)
	}
	return sum
}

// Custom concat operation for part 2
func concatNumbers(x int, y int) int {
	result, err := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	checkError(err)
	return result
}

// Recursively check all combinations of operators
func eval(nums []int, res int, acc int, idx int) bool {
	if idx == len(nums) {
		return acc == res
	}
	// add
	if eval(nums, res, acc+nums[idx], idx+1) {
		return true
	}
	// mul
	if eval(nums, res, acc*nums[idx], idx+1) {
		return true
	}
	// part 2: concat
	if eval(nums, res, concatNumbers(acc, nums[idx]), idx+1) {
		return true
	}
	return false
}

// Either adds result if it can be evaluated or adds 0 otherwise
func add_line(line string) int {
	// Parse result to check
	res, err := strconv.Atoi(strings.Split(line, ":")[0])
	checkError(err)
	nums := []int{}

	// Parse numbers to check
	for _, s_num := range strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ") {
		num, err := strconv.Atoi(s_num)
		checkError(err)
		nums = append(nums, num)
	}

	// Brute Force
	if eval(nums, res, nums[0], 1) {
		return res
	} else {
		return 0
	}

}

func main() {
	fmt.Println(readInput())
}
