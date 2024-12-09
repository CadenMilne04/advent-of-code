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
    return nums
}

type hole struct {
    index int
    length int
}

type file struct {
    id int
    index int
    length int
}

func expand_file_system(nums []int) ([]int, []file, []hole) {
    /* 1. Expand with holes */
    var expanded_fs []int
    var file_list []file
    var holes_list []hole
    cur_index := 0
    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            /* Add num[i], curId's to the expanded_fs */
            file_list = append(file_list, file{id: cur_index, index: len(expanded_fs), length: nums[i]})
            for j := 0; j < nums[i]; j++ {
                expanded_fs = append(expanded_fs, cur_index)
            }
            cur_index++
        } else {
            /* Add num[i], .'s to the expanded_fs */
            holes_list = append(holes_list, hole{index: len(expanded_fs), length: nums[i]})
            for j := 0; j < nums[i]; j++ {
                expanded_fs = append(expanded_fs, -1)
            }
        }
    }
    return expanded_fs, file_list, holes_list
}

func fill_holes(expanded_file_system []int, file_list []file, holes_list []hole) []int {
    for i := len(file_list) - 1; i >= 0; i-- {
        file_id := file_list[i].id
        file_index := file_list[i].index
        file_length := file_list[i].length
        /* Find first hole, fill it by updating the hole's size! */
        for j := 0; j < len(holes_list); j++ {
            hole_index := holes_list[j].index
            hole_length := holes_list[j].length
            if file_length <= hole_length && hole_index < file_index {
                for k := 0; k < file_length; k++ {
                    expanded_file_system[file_index + k] = -1
                    expanded_file_system[hole_index + k] = file_id
                }
                holes_list[j].index = hole_index + file_length
                holes_list[j].length = hole_length - file_length
                // fmt.Println(holes_list[j])
                break
            }
        }
    }
    // fmt.Println(holes_list)
    return expanded_file_system
}

/* Answer is too high */

func consolidate_file_system(nums []int) []int {
    expanded_file_system, file_list, holes_list := expand_file_system(nums)
    fmt.Println("fs:", expanded_file_system)
    fmt.Println("file list:", file_list)
    fmt.Println("holes:", holes_list)
    /* 2. Fill in holes list starting from the end */
    filled_fs := fill_holes(expanded_file_system, file_list, holes_list)
    // file_blocks := len(expanded_file_system) - len(holes_list)
    // filled_fs = filled_fs[:file_blocks]
    return filled_fs
}

func get_checksum(fs []int) int {
    checksum := 0
    for i := 0; i < len(fs); i++ {
        if(fs[i] != -1){
            checksum += fs[i] * i
        }
    }
    return checksum
}

func main() {
    nums := get_nums_array()
    file_system := consolidate_file_system(nums)
    check_sum := get_checksum(file_system)
    fmt.Println("filled fs:", file_system)
    fmt.Println("cs", check_sum)
}
