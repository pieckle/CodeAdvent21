package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("Day-1-Sonar-Sweep/inputs.txt")
	check(err)

	last := -1
	increase1 := 0

	loading := 0
	var windows [4]int
	increase2 := 0

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)

		//Part 1 calc
		if last >= 0 && depth > last {
			increase1++
		}
		last = depth

		//Part 2 calc
		windows[1] += depth
		windows[2] += depth
		windows[3] += depth

		if loading < 3 {
			loading++
		} else {
			if windows[1] > windows[0] {
				increase2++
			}
		}

		windows[0] = windows[1]
		windows[1] = windows[2]
		windows[2] = windows[3]
		windows[3] = 0
	}

	fmt.Print("Part 1 Solution: ")
	fmt.Println(increase1)

	fmt.Print("Part 2 Solution: ")
	fmt.Println(increase2)
}
