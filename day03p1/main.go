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

                for idx := startIdx; idx < j; idx++ {
                    for _, d := range directions {
                        ni, nj := i+d[0], idx+d[1]
                        if inBounds(ni, nj, r, c) && isSymbol(a[ni][nj]) {
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
                    sum += num
                }
            }
        }
    }

    fmt.Println(sum)
}

func isSymbol(char byte) bool {
	// A symbol is anything that is not a period (.) and not a digit
	return char != '.' && !unicode.IsDigit(rune(char))
}
