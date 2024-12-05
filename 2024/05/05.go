package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func readInput() ([][2]string, [][]string) {
    file, err := os.Open("./input")
    checkError(err)
    defer file.Close()

    rules := [][2]string{}
    updates := [][]string{}

    b, err := os.ReadFile("input")
    checkError(err)
    parts := strings.Split(string(b), "\n\n")

    for _, line := range strings.Split(parts[0], "\n") {
        rule := [2]string{line[:2], line[3:5]}
        rules = append(rules, rule)
    }

    for _, line := range strings.Split(parts[1], "\n") {
        updates = append(updates, strings.Split(line, ","))
    }
    return rules, updates 
}

func checkUpdate(rules [][2]string, update []string) bool {
    for i := 0; i < len(update) -1; i++ {
        to_check := [2]string{update[i], update[i+1]}
        
        if !slices.Contains(rules, to_check) {
            return false
        }
    } 
    return true
}

func calcValidUpdates(rules [][2]string, updates [][]string) int {
    sum := 0
	for _, update := range updates {
		if checkUpdate(rules, update) {
			var middle, _ = strconv.Atoi(update[len(update)/2])
			sum += middle
		}
	}
	return sum 
}

func sortUpdate(rules [][2]string, update []string) []string {
    var cmp = func(a string, b string) int {
        if slices.Contains(rules, [2]string{a,b}) {
            return -1
        }
        if slices.Contains(rules, [2]string{b,a}) {
            return 1
        }
        return 0
    }

    slices.SortFunc(update, cmp)

    return update 
}

func correctUpdates(rules [][2]string, updates [][]string) int {
    sum := 0
    for _, update := range updates {
        if !checkUpdate(rules, update) {
            sorted_update := sortUpdate(rules, update)
			var middle, _ = strconv.Atoi(update[len(sorted_update)/2])
			sum += middle
        }
    }

    return sum 
}


func main() {
    rules, updates := readInput()
    a := calcValidUpdates(rules, updates)
    b := correctUpdates(rules, updates)

    fmt.Println("A:",  a)
    fmt.Println("B:",  b)
} 
