package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile() string {
    b, err := os.ReadFile("input")
    checkError(err)
    return string(b)
}

func filterMultiplies(input string) []string {
    pattern := `mul\([1-9][0-9]*,[1-9][0-9]*\)`
    re := regexp.MustCompile(pattern)
    return re.FindAllString(input, -1)
}

func filterConditionalMultiplies(input string) []string {
    do_chunks := strings.Split(input, "do()")
    valid_str := ""

    for _, chunk := range do_chunks {
        valid_str += strings.Split(chunk, "don't()")[0]
    }

    return filterMultiplies(valid_str)
}

func calcMultiplies(mults []string) int {
    sum := 0
    for _, s := range mults {
        pattern := `[1-9]+[0-9]*`
        re := regexp.MustCompile(pattern)
        s_factors := re.FindAllString(s, -1)

        first, err := strconv.Atoi(string(s_factors[0]))
        checkError(err)
        second, err:= strconv.Atoi(string(s_factors[1]))
        checkError(err)
        sum += (first*second)
    }

    return sum
}

func main() {
    blurb := readFile()
    multiplies := filterMultiplies(blurb)
    sum := calcMultiplies(multiplies)
   
    b_multiplies := filterConditionalMultiplies(blurb)
    b_sum := calcMultiplies(b_multiplies)

    fmt.Println("A:", sum)
    fmt.Println("B:", b_sum)
}

