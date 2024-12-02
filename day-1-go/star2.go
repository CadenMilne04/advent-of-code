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

    var r_occurences = make(map[int]int)


    sort.Ints(r_nums)
    sort.Ints(l_nums)
    total := 0
    for i := 0; i < len(r_nums); i++ {
        r_occurences[r_nums[i]] += 1
        value, exists := r_occurences[r_nums[i]]
        if exists {
            r_occurences[r_nums[i]] = value + 1
        } else {
            r_occurences[r_nums[i]] = 1
        }
    }
    for i := 0; i < len(r_nums); i++ {
        value, exists := r_occurences[l_nums[i]]
        if exists {
            total += (value * l_nums[i])
        } 
    }
    fmt.Println(total)
}
