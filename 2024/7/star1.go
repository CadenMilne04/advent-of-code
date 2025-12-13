package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func string_array_to_int_array(string_nums []string) []int {
    nums := []int{}
    for i := 1; i < len(string_nums); i++{
        num, _ := strconv.Atoi(string_nums[i])
        nums = append(nums, num)
    }
    return nums
}

func recursion(goal int, nums []int, i int) bool {
    if(i < 0){
        if(goal == 0) { return true } else { return false }
    }
    
    minus :=  recursion(goal - nums[i], nums, i - 1) 
    if(goal % nums[i] == 0){
        return minus ||recursion(goal / nums[i], nums, i - 1)
    }
    return minus
 }

func main(){
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/7/input.txt")
    scanner := bufio.NewScanner(file)

    total := 0
    for scanner.Scan(){
        line := scanner.Text()
        split := strings.Split(line, ":")
        goal, _ := strconv.Atoi(split[0])
        string_nums := strings.Split(split[1], " ")
        nums := string_array_to_int_array(string_nums)
        if(recursion(goal, nums, len(nums) - 1)){
            total += goal
        }
        fmt.Println(goal, nums)
    }
    fmt.Println(total)
}
