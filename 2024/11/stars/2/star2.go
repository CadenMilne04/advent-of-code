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

type key struct {
    stone string
    blinks int
}

func cStonesDP(stone string, n int, dp map[key]int) int {
    value, ok := dp[key{stone: stone, blinks: n}]
    if ok {
        return value
    }

    if(n < 1) { return 1 }

    stoneNum, _ := strconv.Atoi(stone)

    if stoneNum == 0{
        if0 := cStonesDP("1", n-1, dp)
        dp[key{stone: "1", blinks: n-1}] = if0
        return if0
    } else if len(stone) % 2 == 0 {
        firstHalfNum, _ := strconv.Atoi(stone[:len(stone)/2])
        firstHalfString := strconv.Itoa(firstHalfNum)
        secondHalfNum, _ := strconv.Atoi(stone[len(stone)/2:])
        secondHalfString := strconv.Itoa(secondHalfNum)
        left := cStonesDP(firstHalfString, n - 1, dp)
        right :=cStonesDP(secondHalfString, n - 1, dp)
        dp[key{stone: firstHalfString, blinks: n-1}] = left
        dp[key{stone: secondHalfString, blinks: n-1}] = right
        return left + right
    } else {
        newStoneNum := stoneNum * 2024
        newStoneString := strconv.Itoa(newStoneNum)
        sc := cStonesDP(newStoneString, n-1, dp)
        dp[key{stone: newStoneString, blinks: n-1}] = sc
        return sc
    }
}

func cStones(stone string, n int, dp map[string]int) int {
    if (n < 1) {
        return 1
    }
    stoneNum, _ := strconv.Atoi(stone)

    if stoneNum == 0{
        return cStones("1", n-1, dp)
    } else if len(stone) % 2 == 0 {
        firstHalfNum, _ := strconv.Atoi(stone[:len(stone)/2])
        firstHalfString := strconv.Itoa(firstHalfNum)
        secondHalfNum, _ := strconv.Atoi(stone[len(stone)/2:])
        secondHalfString := strconv.Itoa(secondHalfNum)
        return cStones(firstHalfString, n - 1, dp) + cStones(secondHalfString, n - 1, dp)
    } else {
        newStoneNum := stoneNum * 2024
        newStoneString := strconv.Itoa(newStoneNum)
        return cStones(newStoneString, n-1, dp)
    }
}

func main() {
    stones := readInput()
    fmt.Println(stones)

    total := 0
    dp := map[key]int{}
    for _, stone := range stones {
        total += cStonesDP(stone, 75, dp)
    }

    fmt.Println("Total # of stones:", total)
}
