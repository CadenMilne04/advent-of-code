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
    /* Check safety exluding index i */
    for i := 0; i < len(levels); i++{
        /* Is it increasing or decreasing? */
        increasing := -1 /* -1 for unknown, 1 when increasing, 0 when decreasing */
        prev := -1 /* -1 for undefined */
        safe := true

        /* Check safety */
        for j := 0; j < len(levels); j++{
            /* Skip the index i */
            if i == j { continue }

            current, _ := strconv.Atoi(levels[j])
            /* Dont check anything for the first value */
            if prev == -1 {
                prev = current
                continue
            }

            diff := prev - current
            if diff < 0{
                diff = -diff
            }

            /* We don't know if it is increasing or decreasing yet */
            if increasing == -1 {
                if prev < current{
                    increasing = 1
                } else {
                    increasing = 0
                }
            }


            if (current > prev && increasing == 0) ||
                (current < prev && increasing == 1 ) ||
                (diff > 3 || diff < 1){
                safe = false
            }
            prev = current
        }
        if safe { return true }
    }
    return false
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
