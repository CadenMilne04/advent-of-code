package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load_the_map() [][]int {
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/10/input.txt")
    scanner := bufio.NewScanner(file)
    var the_map [][]int
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, "")
        var row []int
        for _, char := range split {
            num, _ := strconv.Atoi(char)
            row = append(row, num)
        }
        the_map = append(the_map, row)
    }
    return the_map
}

func calculate_score(i int, j int, cur_num int, the_map [][]int, reached_nines [][]bool) int {
    /* base case */
    if (i < 0 || i > len(the_map) || j < 0 || j > len(the_map[0])) {
        return 0;
    }
    if (cur_num == 9 && reached_nines[i][j] == false) {
        reached_nines[i][j] = true
        return 1;
    }

    /* recursive case */
    up := 0
    if(i - 1 >= 0 && the_map[i - 1][j] == cur_num + 1) {
        up = calculate_score(i - 1, j, cur_num + 1, the_map, reached_nines)
    }
    down := 0
    if(i + 1 < len(the_map) && the_map[i + 1][j] == cur_num + 1) {
        down = calculate_score(i + 1, j, cur_num + 1, the_map, reached_nines)
    }
    left := 0
    if(j - 1 >= 0 && the_map[i][j - 1] == cur_num + 1) {
        left = calculate_score(i, j - 1, cur_num + 1, the_map, reached_nines)
    }
    right := 0
    if(j + 1 < len(the_map[0]) && the_map[i][j + 1] == cur_num + 1) {
        right = calculate_score(i, j + 1, cur_num + 1, the_map, reached_nines)
    }
    return up + down + left + right
}

func create_reached_nines(n int, m int) [][]bool {
    var false_map [][]bool
    for range n {
        var row []bool
        for range m {
            row = append(row, false)
        }
        false_map = append(false_map, row)
    }
    return false_map
}


func calculate_trailhead_scores(the_map [][]int) int {
    total := 0
    for i := 0; i < len(the_map); i++ {
        for j := 0; j < len(the_map[0]); j++ {
            if the_map[i][j] == 0 {
                reached_nines := create_reached_nines(len(the_map), len(the_map[0]))
                score := calculate_score(i, j, 0, the_map, reached_nines)
                total += score
            }
        }
    }
    return total
}

func main() {
    the_map := load_the_map()
    trail_head_score := calculate_trailhead_scores(the_map)
    // fmt.Println(the_map)
    fmt.Println(trail_head_score)
}
