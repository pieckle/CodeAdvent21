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
	dat, err := os.ReadFile("Day-2-Dive/inputs.txt")
	check(err)

	depth1 := 0
	horizontal1 := 0

	depth2 := 0
	horizontal2 := 0
	aim2 := 0

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		mag, err := strconv.Atoi(parts[1])
		check(err)

		switch parts[0] {
		case "forward":
			//Part 1 Calc
			horizontal1 += mag

			//Part 2 Calc
			horizontal2 += mag
			depth2 += aim2 * mag
			break
		case "up":
			//Part 1 Calc
			depth1 -= mag

			//Part 2 Calc
			aim2 -= mag
			break
		case "down":
			//Part 1 Calc
			depth1 += mag

			//Part 2 Calc
			aim2 += mag
			break
		}
	}

	fmt.Println("Part 1")
	fmt.Println("Horizontal position: ", horizontal1)
	fmt.Println("Vertical position: ", depth1)
	fmt.Println("Magnitude: ", horizontal1*depth1)

	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("Horizontal position: ", horizontal2)
	fmt.Println("Vertical position: ", depth2)
	fmt.Println("Magnitude: ", horizontal2*depth2)
}
