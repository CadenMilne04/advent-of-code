package main

import (
    "os"
    "log"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func checkLevel(levels []string) bool{
    first, err := strconv.Atoi(levels[0])
    if err != nil {
        log.Fatal(err)
    }
    second, err := strconv.Atoi(levels[1])
    if err != nil {
        log.Fatal(err)
    }
    increasing := true
    if first > second {
        increasing = false
    }
    safe := true
    for i := 1; i < len(levels); i++ {
        current, _ := strconv.Atoi(levels[i])
        prev, _ := strconv.Atoi(levels[i-1])
        if current >= prev && increasing == false {
            return false
        }
        if current <= prev && increasing == true {
            return false
        }
        diff := prev - current
        if diff < 0{
            diff = -diff
        }
        if diff > 3 || diff < 1{
            return false
        }
    }
    return safe
}

func main(){
    file, err := os.Open("/Users/cadenmilne/programming/advent-of-code/day-2/day-2-go/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    safe_reports := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        line := scanner.Text()
        levels := strings.Split(line, " ")
        if checkLevel(levels) {
            safe_reports++
        }
    }
    fmt.Println(safe_reports)
}
