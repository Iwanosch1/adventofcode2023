package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CubeSet struct {
	Blue int
	Green int
	Red int
}


func readFile() []string{
	file, err := os.Open("./input2.txt")
    if (err != nil) {
		log.Fatal(err)
    }
    defer file.Close()
	
    scanner := bufio.NewScanner(file)
	var res []string
    for scanner.Scan() {
		res = append(res, scanner.Text())
        // fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return res
}

func getHighestColors(games []string) [][3]int {
	// always go r-g-b
	var highestColorPerGame [][3]int
	for i, game := range games {
		var gameString = strings.ReplaceAll(game, "Game " + strconv.Itoa(i + 1) + ": ", "")
		var pulls = strings.Split(gameString, ";")
		// get colors of 
		var highestColors [3]int
		for _, pull := range pulls {
			var colors = strings.Split(pull, ",")
			for _, color := range colors {
				var re = regexp.MustCompile("[0-9]+")
				var num, err = strconv.Atoi(re.FindAllString(color, 1)[0])
					if err != nil {
						// ... handle error
						panic(err)
					}
				if (strings.Contains(color, "red") && num > highestColors[0]) {
					highestColors[0] = num
				} else if (strings.Contains(color, "green") && num > highestColors[1]) {
					highestColors[1] = num
				} else if (strings.Contains(color, "blue") && num > highestColors[2]) {
					highestColors[2] = num
				}
			}
		}
		highestColorPerGame = append(highestColorPerGame, highestColors)
	}
	return highestColorPerGame
} 

func main() {
	var games = readFile()
	
	// 01
	// var maxValues = [3]int{12,13,14}
	// var sum1 = 0
	// for i, gameColors := range getHighestColors(games) {
	// 	if (maxValues[0] < gameColors[0] || maxValues[1] < gameColors[1] || maxValues[2] < gameColors[2]) {
	// 		fmt.Println("impossible", i + 1, gameColors)
	// 	} else {
	// 		sum1 += i+1
	// 	}
	// }

	// 02
	var sum2 = 0
	for _, gameColors := range getHighestColors(games) {
		var power = gameColors[0] * gameColors[1] * gameColors[2]
		fmt.Println(power)
		sum2 += power
	}

	fmt.Println(sum2)
}