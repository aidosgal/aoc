package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func findAndCalculate(x int, list []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == x {
			count++
		}
	}

	ch <- count * x
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
	ch := make(chan int, len(list1))

	for i := 0; i < len(list1); i++ {
		wg.Add(1)
		go findAndCalculate(list1[i], list2, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	ans := 0
	for sum := range ch {
		ans += sum
	}

	fmt.Println(ans)
}
