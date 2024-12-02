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

func isSafe(report []int) bool {
    if len(report) < 2 {
        return true
    } else {
        direction := 0
        next_direction := 0
        for i := 0; i < len(report)-1; i++ {
            if report[i+1] > report[i] {
                next_direction = 1
            } else if report[i+1] < report[i] {
                next_direction = -1
            } else {
                return false
            }

            if (next_direction - direction) > 1 || (next_direction - direction) < -1 {
                return false
            } else {
                if (report[i+1] - report[i] > 3 || report[i+1] - report[i] < -3) {
                    return false
                }
                direction = next_direction
            }
        }
    }
    return true
}

func testDampened(report []int) bool {
    for i := 0; i < len(report); i++ {
        temp := make([]int, len(report))
        copy(temp, report)

        d_report := append(temp[:i], temp[i+1:]...)

        if isSafe2(d_report, true) {
            return true
        }
    }
    return false
}


func isSafe2(report []int, dampened bool) bool {
    if len(report) < 2 {
        return true
    } else {
        direction := 0
        next_direction := 0
        for i := 0; i < len(report)-1; i++ {
            if report[i+1] > report[i] {
                next_direction = 1
            } else if report[i+1] < report[i] {
                next_direction = -1
            } else {
                if dampened {
                    return false
                } else {
                    return testDampened(report)
                }
            }

            if (next_direction - direction) > 1 || (next_direction - direction) < -1 {
                if dampened {
                    return false
                } else {
                    return testDampened(report)
                }
            } else {
                if (report[i+1] - report[i] > 3 || report[i+1] - report[i] < -3) {
                if dampened {
                    return false
                } else {
                    return testDampened(report)
                }
                }
                direction = next_direction
            }
        }
    }

    return true
}



func readLists2() int {
    file, err := os.Open("./input")
    checkError(err)
    defer file.Close()
   
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        s_report := strings.Split(line, " ")
        report := []int{}

        for _, i := range s_report {
            j, err := strconv.Atoi(i)
            checkError(err)
            report = append(report, j)
        }

        if isSafe2(report, false) {
            sum++
        }
    }
    checkError(scanner.Err())

    return sum 
}

func readLists() int {
    file, err := os.Open("./input")
    checkError(err)
    defer file.Close()
   
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        s_report := strings.Split(line, " ")
        report := []int{}

        for _, i := range s_report {
            j, err := strconv.Atoi(i)
            checkError(err)
            report = append(report, j)
        }

        if isSafe(report) {
            sum++
        }
    }
    checkError(scanner.Err())

    return sum 
}
func main() {
    fmt.Println("Safe reports:", readLists())
    fmt.Println("Safe dampened reports:", readLists2())
}

