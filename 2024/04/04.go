package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func readInput() [][]string {
    playing_field := [][]string{}

    file, err := os.Open("./input")
    checkError(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        row := []string{}
        for _, c := range line {
            row = append(row, string(c))
        }
        playing_field = append(playing_field, row)
    }
    checkError(scanner.Err())

    return playing_field
}

func searchInDirection(playing_field [][]string, word string, x_pos int, y_pos int, direction [2]int) bool {
    if (playing_field[y_pos][x_pos] != string(word[0])) {
        return false
    }

    dir_row, dir_col := direction[0], direction[1]

    for i := 0; i < len(word); i++ {
       // check out of bounds
       if x_pos + (i*dir_col) < 0 || x_pos + (i*dir_col) > (len(playing_field[0]) - 1) || y_pos + (i*dir_row) < 0 || y_pos + (i*dir_row) > (len(playing_field) - 1) {
            return false
       }

       // check letter
       if playing_field[y_pos + (i*dir_row)][x_pos + (i*dir_col)] != string(word[i]) {
            return false
       }
    }    

    return true
}

func countOccurences(playing_field [][]string, isA bool) int {
    directions := [8][2]int{
        {-1, 0},  // N 
        {1, 0},   // S 
        {0, -1},  // W 
        {0, 1},   // O 
        {-1, -1}, // NW
        {-1, 1},  // NO
        {1, -1},  // SW
        {1, 1},   // SO
    }

    rows := len(playing_field)
    cols := len(playing_field[0])
    sum := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
                if isA {
                    for _, d := range directions {
                        if searchInDirection(playing_field, "XMAS", c, r, d) {
                            sum++
                        }
                    }
                } else {
                    if searchCrossMAS(playing_field, c, r) {
                        sum++
                    }
                }
            }  
        }
        return sum
    }


func printPattern(playing_field [][]string, x_pos int, y_pos int) {
    fmt.Println("---")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(playing_field[i+y_pos][j+x_pos])
        }
        fmt.Print("\n") 
    }
    fmt.Println("---")
}

func searchCrossMAS(playing_field [][]string, x_pos int, y_pos int) bool {
    if playing_field[y_pos][x_pos] == "M" {
        return  (
                len(playing_field) - 1 >= y_pos + 2 &&
                len(playing_field[0]) - 1 >= x_pos + 2 &&
                playing_field[y_pos][x_pos+2] == "M" &&
                playing_field[y_pos+1][x_pos+1] == "A" &&
                playing_field[y_pos+2][x_pos] == "S" &&
                playing_field[y_pos+2][x_pos+2] == "S") ||
                (
                len(playing_field) - 1 >= y_pos + 2 &&
                len(playing_field[0]) - 1 >= x_pos + 2 &&
                playing_field[y_pos][x_pos+2] == "S" &&
                playing_field[y_pos+1][x_pos+1] == "A" &&
                playing_field[y_pos+2][x_pos] == "M" &&
                playing_field[y_pos+2][x_pos+2] == "S")

    } else if playing_field[y_pos][x_pos] == "S" {
        return (len(playing_field) - 1 >= y_pos + 2 &&
               len(playing_field[0]) - 1 >= x_pos + 1 &&
               len(playing_field[0]) - 1 >= x_pos + 2 &&
               playing_field[y_pos][x_pos+2] == "S" &&
               playing_field[y_pos+1][x_pos+1] == "A" &&
               playing_field[y_pos+2][x_pos] == "M" &&
               playing_field[y_pos+2][x_pos+2] == "M") ||
               (
                len(playing_field) - 1 >= y_pos + 2 &&
                len(playing_field[0]) - 1 >= x_pos + 2 &&
                playing_field[y_pos][x_pos+2] == "M" &&
                playing_field[y_pos+1][x_pos+1] == "A" &&
                playing_field[y_pos+2][x_pos] == "S" &&
                playing_field[y_pos+2][x_pos+2] == "M")
    } else {
        return false
    }   
}

func main() {
    playing_field := readInput()
    fmt.Println("A:", countOccurences(playing_field, true))
    fmt.Println("B:", countOccurences(playing_field, false))
} 
