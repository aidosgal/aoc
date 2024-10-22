package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main () {
    a := 1000

    ans := 0

    for a > 0 {
        var s string
        fmt.Scan(&s)
        var lnum string
        for _, c := range s {
            if(unicode.IsDigit(c)){
                lnum += string(c) 
                break
            }
        }
        rs := reverseString(s)
        for _, c := range rs {
            if(unicode.IsDigit(c)){
                lnum += string(c) 
                break
            }
        }
        num, err := strconv.Atoi(lnum)
        if err != nil {
            panic(err)
        }
        ans += num
        a-- 
    }
    fmt.Println(ans)
}

func reverseString(s string) string {
    runes := []rune(s)

    for i, j := 0, len(runes)-1; i< j; i, j = i + 1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}
