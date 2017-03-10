package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var charWidthMap [26]int64

func main() {
	var inputCharsWidth string
	var inputSelected string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputCharsWidth = scanner.Text()
	initWidthMapping(inputCharsWidth)

	scanner.Scan()
	inputSelected = scanner.Text()

	var result int64 = 1
	var max int64
	for _, char := range inputSelected {
		width := getWidthFor(char)
		result = result * width
		max = getMax(max, width)
	}

	fmt.Println(int64(len(inputSelected)) * 1 * max)
}

func initWidthMapping(inputWidth string) {
	charsWidth := strings.Split(inputWidth, ` `)

	for idx, item := range charsWidth {
		number, _ := strconv.ParseInt(item, 10, 0)
		charWidthMap[idx] = number
	}
}

func getWidthFor(char rune) int64 {
	return charWidthMap[char-97]
}

func getMax(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
