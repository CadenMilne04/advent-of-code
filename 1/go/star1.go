package main

import ("fmt"
        "os"
        "bufio"
        "log"
        "strings"
        "strconv"
        "sort")

func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    r_nums := []int{}
    l_nums := []int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        split := strings.Fields(scanner.Text())
        r_num, err := strconv.Atoi(split[0])
        if err != nil {
            log.Fatal(err)
        }
        l_num, err := strconv.Atoi(split[1])
        if err != nil {
            log.Fatal(err)
        }
        r_nums = append(r_nums, r_num)
        l_nums = append(l_nums, l_num)
    }
    sort.Ints(r_nums)
    sort.Ints(l_nums)
    total := 0
    for i := 0; i < len(r_nums); i++ {
        diff := r_nums[i] - l_nums[i]
        if diff < 0 {
            total += -diff
        } else {
            total += diff
        }
    }

    fmt.Println(r_nums)
    fmt.Println(l_nums)
    fmt.Println(total)

}
