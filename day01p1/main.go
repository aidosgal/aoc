package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sort(list []int, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}

	ch <- list
}

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

	var list1 []int
	var list2 []int

	for _, line := range lines {
		num1, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		num2, _ := strconv.Atoi(strings.Split(line, "   ")[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	var wg sync.WaitGroup
	ch := make(chan []int, 2)

	wg.Add(2)
	go sort(list1, &wg, ch)
	go sort(list2, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var sortedLists [][]int
	for sortedList := range ch {
		sortedLists = append(sortedLists, sortedList)
	}

	ans := 0
	for i := range sortedLists[0] {
		ans += int(math.Abs(float64(sortedLists[0][i] - sortedLists[1][i])))
	}

	fmt.Println(ans)
}
