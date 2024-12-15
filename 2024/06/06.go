package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

func (c coord) Equal(other coord) bool {
	return c.x == other.x && c.y == other.y
}

func contains(slice []coord, target coord) bool {
	for _, item := range slice {
		if item.Equal(target) {
			return true
		}
	}
	return false
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() ([][]string, coord) {
	playing_field := [][]string{}

	file, err := os.Open("./input")
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r_idx := 0
	player_start := coord{x: -1, y: -1}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row := []string{}
		for c_idx, c := range line {
			row = append(row, string(c))
			if string(c) == "^" {
				player_start = coord{x: c_idx, y: r_idx}
			}
		}
		playing_field = append(playing_field, row)
		r_idx++
	}
	checkError(scanner.Err())

	return playing_field, player_start
}

// -> ( 1, 0)
//  v ( 0, 1)
// <- (-1, 0)
//  ^ ( 0,-1)

func explore_next_step(playing_field [][]string, cur_pos coord, direction coord) (string, coord, coord) {
	// out of bound
	next_pos := coord{cur_pos.x + direction.x, cur_pos.y + direction.y}
	if next_pos.y < 0 || next_pos.y >= len(playing_field) || next_pos.x < 0 || next_pos.x >= len(playing_field[next_pos.y]) {
		return "end", cur_pos, direction
	}

	if playing_field[next_pos.y][next_pos.x] == "." || playing_field[next_pos.y][next_pos.x] == "^" {
		return ".", coord{next_pos.x, next_pos.y}, direction
	} else if playing_field[next_pos.y][next_pos.x] == "#" {
		new_direction := coord{-direction.y, direction.x}
		return "#", cur_pos, new_direction
	}

	return "error", cur_pos, direction
}

func a() (int, []coord) {
	playing_field, player_pos := readInput()
	out_of_bound := false
	visited := []coord{}
	// Starting Pos
	visited = append(visited, player_pos)

	cell := "^"
	pos := player_pos
	dir := coord{0, -1}

	for !out_of_bound {
		cell, pos, dir = explore_next_step(playing_field, pos, dir)
		if cell == "." {
			if !contains(visited, pos) {
				visited = append(visited, pos)
			}
		} else if cell == "end" {
			out_of_bound = true
		}
	}

	return len(visited), visited
}

func testObstacles(path []coord) int {
	playing_field, player_pos := readInput()
	path = path[1:]
	loops := 0

	for _, c := range path {
		test_field := make([][]string, len(playing_field))
		for i := range playing_field {
			test_field[i] = make([]string, len(playing_field[i]))
			copy(test_field[i], playing_field[i])
		}

		test_field[c.y][c.x] = "#"
		if checkForLoop(test_field, player_pos) {
			loops++
		}
	}
	return loops
}

func checkForLoop(playing_field [][]string, player_pos coord) bool {
	out_of_bound := false
	// Starting Pos
	visited_count := make(map[coord]int)

	cell := "^"
	pos := player_pos
	dir := coord{0, -1}

	for !out_of_bound {
		cell, pos, dir = explore_next_step(playing_field, pos, dir)
		if cell == "." {
			visited_count[pos]++
			if visited_count[pos] > 4 {
				return true
			}
		} else if cell == "end" {
			out_of_bound = true
		}
	}

	return false
}

func main() {
	cells, visited := a()
	fmt.Println("A:", cells)

	loops := testObstacles(visited)
	fmt.Println("B:", loops)
}
