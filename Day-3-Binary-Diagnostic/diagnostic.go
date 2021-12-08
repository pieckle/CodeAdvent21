package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func search(slice [][]rune, high bool) []rune {
	return doSearch(slice, 0, high)
}

func doSearch(slice [][]rune, index int, high bool) []rune {
	if len(slice)==1 {
		return slice[0]
	}

	upList := make([][]rune, 0)
	downList := make([][]rune, 0)

	count := 0
	for i := 0; i < len(slice); i++ {
		bit, err := strconv.Atoi(string(slice[i][index]))
		check(err)
		count += bit

		if bit==1 {
			upList = append(upList, slice[i])
		} else {
			downList = append(downList, slice[i])
		}
	}

	match := 'X'
	if len(slice)%2==0 && count==len(slice)/2 {
		if high {
			match = '1'
		} else {
			match = '0'
		}
	} else if (count>len(slice)/2)==high {
		match = '1'
	} else {
		match = '0'
	}

	if match=='1' {
		return doSearch(upList, index+1, high)
	} else {
		return doSearch(downList, index+1, high)
	}
}

func main() {
	dat, err := os.ReadFile("Day-3-Binary-Diagnostic/inputs.txt")
	check(err)

	count := 0
	var fields [12]int

	allEntries := make([][]rune, 0)

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		bits := []rune(scanner.Text())
		allEntries = append(allEntries, bits)
		for i := 0; i < 12; i++ {
			bit, err := strconv.Atoi(string(bits[i]))
			check(err)
			fields[i] += bit
		}
		count++
	}

	gamma := 0
	epsilon := 0
	for i := 0; i < 12; i++ {
		bit := int(math.Pow(2, float64(11-i)))
		if fields[i] > count/2 {
			gamma = gamma | bit
		} else {
			epsilon = epsilon | bit
		}
	}

	o2rateLiteral := search(allEntries, true)
	o2rateValue := 0
	for i := 0; i < 12; i++ {
		if o2rateLiteral[i]=='1' {
			o2rateValue = o2rateValue | int(math.Pow(2, float64(11-i)))
		}
	}
	co2rateLiteral := search(allEntries, false)
	co2rateValue := 0
	for i := 0; i < 12; i++ {
		if co2rateLiteral[i]=='1' {
			co2rateValue = co2rateValue | int(math.Pow(2, float64(11-i)))
		}
	}

	fmt.Println("Part 1")
	fmt.Println("Gamma: ", gamma)
	fmt.Println("Epsilon: ", epsilon)
	fmt.Println("Fuel Consumption: ", gamma*epsilon)
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println("O2 Rate: ", o2rateValue)
	fmt.Println("CO2 Rate: ", co2rateValue)
	fmt.Println("Life Support: ", o2rateValue*co2rateValue)
}
