package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// reads the file into an bi-dimensional array with 1 char per entry
func readFile() [][]string{
	file, err := os.Open("./input3.txt")
    if (err != nil) {
		log.Fatal(err)
    }
    defer file.Close()
	
    scanner := bufio.NewScanner(file)
	var res [][]string
    for scanner.Scan() {
        var lineString = scanner.Text()
        var lineArray []string
        for i := 0; i < len(lineString); i++ {   //run a loop and iterate through each character
            lineArray = append(lineArray, string(lineString[i]))  //print characters along with the space character
        }
		res = append(res, lineArray)
        // fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return res
}

// checks if to given entry an adjacent value is not a number and != "."
func hasSymbolAdjacent(inputArray [][]string, i int, j int) bool {
    var res = false
    for m := max(0,i-1); m < min(i+2,len(inputArray)); m++ {
        for n := max(0, j-1); n < min(j+2, len(inputArray[i])); n++ {
            if _, err := strconv.Atoi(inputArray[m][n]); err != nil {
                if (inputArray[m][n] != ".") {
                    return true
                } 
            }
        }
    } 
    return res
}

func main() {
    var inputArray = readFile()
    var sum = 0
    for i := 0; i < len(inputArray); i++ {
        var current int = 0
        var currentValid = false
        for j := 0; j < len(inputArray[i]); j++ {
            if !currentValid {
                currentValid = hasSymbolAdjacent(inputArray, i, j)
            } 
            if num, err := strconv.Atoi(inputArray[i][j]); err == nil {
                if current == 0 {
                    current = num
                } else {
                    current = current * 10
                    current += num
                }
            } else {
                if currentValid {
                    sum += current
                }
                currentValid = false
                current = 0
            }
        }
    }

    fmt.Println(sum)
    // fmt.Println(len(inputArray))
    // fmt.Println(len(inputArray[0]))
}