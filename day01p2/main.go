package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
    a := 1000
    m := map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
    ans := 0
    for a > 0 {
        var s string
        fmt.Scan(&s)
        var lnum string
        var fd, ld string
        var word string
        for _, c := range s {
            if(unicode.IsDigit(c)) {
                if fd == "" {
                    fd = string(c)
                }
                ld = string(c)
            }else{
                word += string(c)

                for w, digit := range m {
                    if strings.HasSuffix(word, w) {
                        if fd == "" {
                            fd = digit
                        }
                        ld = digit
                        if len(word) > 0 {
							word = string(word[len(word)-1]) 
						}
                    }
                }
            } 
        }
        lnum = fd + ld
        num, err := strconv.Atoi(lnum)
        if err != nil {
            panic(err)
        }
        ans += num
        a--
    }
    fmt.Println(ans)
}
