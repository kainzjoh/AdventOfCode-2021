package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(file string) (filetext []string) {
	dat, _ := os.Open(file)
	defer dat.Close()
	s := bufio.NewScanner(dat)
	for s.Scan() {
		filetext = append(filetext, s.Text())
	}
	return filetext
}

func main() {
	var x, y, move, aim, x2, y2 int
	s := Readfile("move1.txt")
	for i := 0; i < len(s)-1; i++ {
		mm := strings.Split(s[i], " ")
		move, _ = strconv.Atoi(mm[1])

		// part1
		switch mm[0] {
		case "up":
			y = y - move
		case "down":
			y = y + move
		case "forward":
			x = x + move
		}

		// part2
		switch mm[0] {
		case "up":
			aim = aim - move
		case "down":
			aim = aim + move
		case "forward":
			x2 = x2 + move
			y2 = y2 + (aim * move)
		}
	}

	fmt.Println(x * y)
	fmt.Println(x2 * y2)
}
