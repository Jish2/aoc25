package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var p int64 = 50
	var zeroes = 0

	file, err := os.Open("01/input-1.txt")
	// file, err := os.Open("01/test.txt")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		dir := string(line[0])
		num, err := strconv.ParseInt(line[1:], 0, 64)

		if err != nil {
			fmt.Println(err)
		}

		if dir == "L" {
			p = (p + 100 - num) % 100
		} else {
			p = (p + 100 + num) % 100
		}

		if p == 0 {
			zeroes++
		}

	}

	fmt.Println(p, zeroes)

}
