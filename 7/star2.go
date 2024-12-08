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

func can_equal(test int, total int, nums []int, i int) bool {
    if(i == len(nums)){
        if(total == test) { 
            fmt.Println("correct: ", nums)
            return true 
        } else { return false }
    }
    
    // fmt.Println(nums, i, total)
    minus :=  can_equal(test, total + nums[i], nums, i + 1) 
    divide := can_equal(test, total * nums[i], nums, i + 1)
    concat := false
    if(i < len(nums) - 1){
        var new_nums []int
        for j := 0; j < len(nums); j++ {
            new_nums = append(new_nums, nums[j])
        }
        // fmt.Println("old stuff", nums, new_nums, i, total)
        concatted_nums := strconv.Itoa(total*nums[i]) + strconv.Itoa(nums[i+1])
        // fmt.Println("new num for i + 1", concatted_nums)
        new_num, _ := strconv.Atoi(concatted_nums)
        new_nums[i+1] = new_num
        // fmt.Println("new array", new_nums)
    
        concat = can_equal(test, 0, new_nums, i + 1)
    }
    concat2 := false
    if(i < len(nums) - 1){
        var new_nums2 []int
        for j := 0; j < len(nums); j++ {
            new_nums2 = append(new_nums2, nums[j])
        }
        concatted_nums2 := strconv.Itoa(total+nums[i]) + strconv.Itoa(nums[i+1])
        new_num2, _ := strconv.Atoi(concatted_nums2)
        new_nums2[i+1] = new_num2

        concat2 = can_equal(test, 0, new_nums2, i + 1)
    }
    return minus || divide || concat || concat2
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
        if(can_equal(goal, 0, nums, 0)){
            total += goal
        }
    }
    fmt.Println(total)
}

/* The new operator is NOT associative */
