/* Implement a sorting algorithm by using dependencies as the comparator */

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/* map[right number][left number] == true*/
/* ^^^ Means that the left number must be before the right number  */

func find_non_depended_upon_index(nums []string, dependencies map[int]map[int]bool) int{
    /* Find the one with the no dependencies */
    if len(nums) == 1{
        return 0
    }
    for i := 0; i < len(nums); i++{
        num1, _ := strconv.Atoi(nums[i])
        has_dependency := false
        for j := 0; j < len(nums); j++ {
            num2, _ := strconv.Atoi(nums[j])
            if (i != j){
                /* if num2 must be before the right number, the right number is NOT the earliest number */
                if (dependencies[num1][num2] == true){
                    has_dependency = true
                }
            }
        }
        if !has_dependency {return i}
    }
    return -1
}

func selction_sort(nums []string, dependencies map[int]map[int]bool) []int{
    result := []int{}
    length := len(nums)
    for i := 0; i < length; i++{
        /* Find the index of nums that should be at the start. */
        minI := find_non_depended_upon_index(nums, dependencies)
        /* Remove it from the list */
        newNum, _ := strconv.Atoi(nums[minI])
        result = append(result, newNum)
        nums = slices.Delete(nums, minI, minI + 1)
    }
    return result
}

func main()  {
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/day-5/input.txt")
    scanner := bufio.NewScanner(file)

    var set map[int]map[int]bool
    set = make(map[int]map[int]bool)

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {break}
        split := strings.Split(line, "|")
        num, _:= strconv.Atoi(split[1])
        dependency, _:= strconv.Atoi(split[0])

        /* Add dependency to num's set */
        if set[num] == nil {
            set[num] = make(map[int]bool)
        }
        set[num][dependency] = true
    }
    total := 0
    for scanner.Scan() {
        safe := true
        line := scanner.Text()
        nums := strings.Split(line, ",")

        cant_have := map[int]bool{};
        for _, snum := range nums {
            num, _ := strconv.Atoi(snum)

            if(cant_have[num] == true){
                safe = false
                break
            }

            dependencies := set[num]
            for key, _ := range dependencies {
                cant_have[key] = true
            }
        }
        if safe {
        } else {
            sorted_list := selction_sort(nums, set)
            mid := len(sorted_list)/2
            mval := sorted_list[mid] 
            total += mval
        }
    }
    fmt.Println(total)
}


/* map[right number][left number] == true*/
/* ^^^ Means that the left number must be before the right number  */
