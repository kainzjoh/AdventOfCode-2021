package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var value int = 0
	var i int = 0
	var counter int = 0
	var preValue int = 0
	for scanner.Scan() {

		if i > 0 {
			preValue, _ = strconv.Atoi(scanner.Text())
			if preValue > value {
				counter += 1
			}
		}

		value, _ = strconv.Atoi(scanner.Text())
		i++
	}
	fmt.Println(counter)

}
