package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type list_node struct {
    value int;
    next *list_node;
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
            mid := len(nums)/2
            mval, _ := strconv.Atoi(nums[mid])
            total += mval
        }
    }
    fmt.Println(total)
}

/* Notes */
// dependencies[curNum] = set(x, y, z)
// seenAlready = set(x, y, z)
// if there is a dependency that we haven't seen already, it's invalid
