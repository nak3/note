package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/awk-2/problem

func main() {

	file, err := os.Open("input")
	if err != nil {
		// TODO error
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		tmp := strings.Split(str, " ")
		if len(tmp) != 4 {
			// TODO error
			os.Exit(1)
		}
		p := "Pass"
		for i := 1; i < 4; i++ {
			score, _ := strconv.Atoi(tmp[i]) // TODO: error
			if score < 50 {
				p = "Fail"
				break
			}
		}
		fmt.Printf("%s : %s\n", tmp[0], p) // output for debug
	}
}
