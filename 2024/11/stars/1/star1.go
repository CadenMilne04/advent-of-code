package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
    var inputArray []string
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/11/input.txt")
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        text := scanner.Text()
        split := strings.Split(text, " ")
        for _, word := range split {
            inputArray = append(inputArray, word)
        }
    }
    return inputArray
}

func stonesAfterBlinking(stones []string) []string {
    var newStones []string
    for _, stone := range stones {

        stoneNum, _ := strconv.Atoi(stone)
        if stoneNum == 0 {
            newStones = append(newStones, "1")
        } else if len(stone) % 2 == 0 {
            firstHalfNum, _ := strconv.Atoi(stone[:len(stone)/2])
            firstHalfString := strconv.Itoa(firstHalfNum)
            secondHalfNum, _ := strconv.Atoi(stone[len(stone)/2:])
            secondHalfString := strconv.Itoa(secondHalfNum)
            newStones = append(newStones, firstHalfString)
            newStones = append(newStones, secondHalfString)
        } else {
            newStoneNum := stoneNum * 2024
            newStoneString := strconv.Itoa(newStoneNum)
            newStones = append(newStones, newStoneString)
        }
    }
    return newStones
}

func main() {
    stones := readInput()
    fmt.Println(stones)
    for i := 1; i <= 25; i++ {
        stones = stonesAfterBlinking(stones)
        fmt.Println("Stones after blinking", i, "times: \n", stones)
    }

    fmt.Println("Total # of stones:", len(stones))
}
