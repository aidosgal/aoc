package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    readFile, err := os.Open("./input.txt")
    if err != nil {
        fmt.Println(err)
    }

    fs := bufio.NewScanner(readFile)
    fs.Split(bufio.ScanLines)
    var lines []string

    for fs.Scan() {
        lines = append(lines, fs.Text())
    }
    ans := 0
    for _, line := range lines {
        game_lines := strings.Split(line, ": ")

        segs := strings.Split(game_lines[1], "; ")

        max_r := 0
        max_g := 0
        max_b := 0
        for _, seg := range segs {
            colors := strings.Split(seg, ", ")
            for _, color := range colors {
                c := strings.Split(color, " ")
                color_number, err := strconv.Atoi(c[0])
                if err != nil {
                    fmt.Print(err)
                }
                if (c[1] == "blue" && color_number > max_b) {
                    max_b = color_number 
                }
                if (c[1] == "red" && color_number > max_r) {
                    max_r = color_number
                }
                if (c[1] == "green" && color_number > max_g) {
                    max_g = color_number
                }
            } 
        }
        prd := max_r * max_b * max_g
        ans += prd
    }
    fmt.Println(ans)
}
