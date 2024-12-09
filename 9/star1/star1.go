package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_nums_array() []int {
    num_bytes, _ := os.ReadFile("/Users/cadenmilne/programming/advent-of-code/9/input.txt")
    var nums []int
    string_nums := strings.Split(string(num_bytes), "") 
    for _, string_num := range string_nums {
        num, _ := strconv.Atoi(string_num)
        nums = append(nums, num)
    }
    nums = nums[:len(nums)-1]
    return nums
}

func expand_file_system(nums []int) ([]int, []int, []int) {
    /* 1. Expand with holes */
    var expanded_fs []int
    var holeless_fs []int
    var holes_list []int
    cur_index := 0
    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            /* Add num[i], curId's to the expanded_fs */
            for j := 0; j < nums[i]; j++ {
                expanded_fs = append(expanded_fs, cur_index)
                holeless_fs = append(holeless_fs, cur_index)
            }
            cur_index++
        } else {
            /* Add num[i], .'s to the expanded_fs */
            for j := 0; j < nums[i]; j++ {
                holes_list = append(holes_list, len(expanded_fs))
                expanded_fs = append(expanded_fs, -1)
            }
        }
    }
    return expanded_fs, holeless_fs, holes_list
}

func fill_holes(expanded_file_system []int, holeless_file_system []int, holes_list []int) []int {
    for i := len(holes_list) - 1; i >= 0; i-- {
        hole_index := holes_list[i]
        expanded_file_system [hole_index] = holeless_file_system[len(holeless_file_system) - 1 - i]
    }
    return expanded_file_system
}

func consolidate_file_system(nums []int) []int {
    expanded_file_system, holeless_file_system, holes_list := expand_file_system(nums)
    fmt.Println("fs:", expanded_file_system)
    fmt.Println("holeless fs:", holeless_file_system)
    fmt.Println("holes:", holes_list)
    /* 2. Fill in holes list starting from the end */
    filled_fs := fill_holes(expanded_file_system, holeless_file_system, holes_list)
    file_blocks := len(expanded_file_system) - len(holes_list)
    filled_fs = filled_fs[:file_blocks]
    return filled_fs
}

func get_checksum(fs []int) int {
    checksum := 0
    for i := 0; i < len(fs); i++ {
        checksum += fs[i] * i
    }
    return checksum
}

func main() {
    nums := get_nums_array()
    file_system := consolidate_file_system(nums)
    check_sum := get_checksum(file_system)
    fmt.Println("nums:", nums)
    fmt.Println("filled fs:", file_system)
    fmt.Println("cs", check_sum)
}
