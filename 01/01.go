package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getSum(data string) int {
	lines := strings.Split(data, "\n")

	sum := 0
	for _, line := range lines {
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)

		if len(numbers) > 0 {
			firstNum, _ := strconv.Atoi(numbers[0])
			lastNum, _ := strconv.Atoi(numbers[len(numbers)-1])

			lineNumber := firstNum*10 + lastNum
			sum += lineNumber
		}
	}
	return sum
}

func main() {
	file, err := os.Open("./input1.txt")
	if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	output := getSum(data)
	fmt.Println("final sum: ", output)
}
