package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func readLists() ([]int, []int) {
    file, err := os.Open("./input")
    checkError(err)
    defer file.Close()
   
    var left []int
    var right []int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        numbers := strings.Split(line, "   ")
        leftNumber, err := strconv.Atoi(numbers[0])
        checkError(err)
        rightNumber, err := strconv.Atoi(numbers[1])
        checkError(err)

        left = append(left, leftNumber)
        right = append(right, rightNumber)
    }
    checkError(scanner.Err())

    return left, right
}

func getDistance(left []int, right []int) int {
    sort.Ints(left)
    sort.Ints(right)
    sum := 0

    for i := 0; i < len(left); i++ {
        if (left[i] >= right[i]) {
            sum += left[i] - right[i]
        } else {
            sum += right[i] - left[i]
        }
    }
    
    return sum
}

func getSimilarity(left []int, right []int) int {
    rightOcc := make(map[int]int)
    sum := 0

    for i := 0; i < len(right); i++ {
        rightOcc[right[i]]++
    }

    for i := 0; i < len(left); i++ {
        sum += left[i] * rightOcc[left[i]]
    }

    return sum 
}

func main() {
    left, right := readLists()
    distance := getDistance(left, right)
    similarity := getSimilarity(left, right)

    fmt.Println("Distance:", distance)
    fmt.Println("Similarity:", similarity)
}

