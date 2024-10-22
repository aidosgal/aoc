package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
    a := [140]string{} 
    for i := 0; i < 140; i++ {
        fmt.Scan(&a[i])
    }

    directions := [8][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},         {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }

    inBounds := func(x, y, rows, cols int) bool {
        return x >= 0 && x < rows && y >= 0 && y < cols
    }

    r := len(a)
    c := len(a[0])
    sum := 0

    t := make(map[int]map[int]map[int]int)
    var x, y int
    f := 0 

    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            snum := ""

            for j < c && unicode.IsDigit(rune(a[i][j])) {
                snum += string(a[i][j])
                j++
            }
            if snum != "" { 
                is_part := false

                numLen := len(snum)
				startIdx := j - numLen
                x = 0
                y = 0

                for idx := startIdx; idx < j; idx++ {
                    for _, d := range directions {
                        ni, nj := i+d[0], idx+d[1]
                        if inBounds(ni, nj, r, c) && a[ni][nj] == '*' {
                            x = ni
                            y = nj
                            is_part = true
                            break
                        } 
                    }
                    if is_part {
                        break
                    }
                }
                if is_part {
                    num, err := strconv.Atoi(snum)
                    if err != nil {
                        fmt.Print(err)
                    }
                    if t[f] == nil {
                        t[f] = make(map[int]map[int]int)
                    }
                    if t[f][x] == nil {
                        t[f][x] = make(map[int]int)
                    }
                    t[f][x][y] = num     
                    f++
                }
            }
        }
    }

    seen := make(map[[2]int][]int)

	for _, xMap := range t {
		for x, yMap := range xMap {
			for y, num := range yMap {
				coord := [2]int{x, y}
				seen[coord] = append(seen[coord], num)
			}
		}
	}

	for _, values := range seen {
		if len(values) == 2 {
			product := values[0] * values[1]
			sum += product
		}
	}

    fmt.Println(sum)
}
