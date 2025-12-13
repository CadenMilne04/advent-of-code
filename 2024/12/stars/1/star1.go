package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func load_the_map() [][]string {
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/12/input.txt")
    scanner := bufio.NewScanner(file)
    var the_map [][]string
    for scanner.Scan() {
        text := scanner.Text()
        split := strings.Split(text, "")
        var row []string
        for _, char := range split {
            row = append(row, char)
        }
        the_map = append(the_map, row)
    }
    return the_map
}

func create_visited_map(the_map [][]string) [][]bool {
    var visited [][]bool
    for i := 0; i < len(the_map); i++ {
        var row []bool
        for j := 0; j < len(the_map[0]); j++ {
            row = append(row, false)
        }
        visited = append(visited, row)
    }
    return visited
}

func calculate_area_for_region(i int, j int, cur_patch string, the_map [][]string, visited [][]bool) (int, int) {
    if(i < 0 || i >= len(the_map) || j < 0 || j >= len(the_map[0])){
        return 0, 1
    }
    if(the_map[i][j] != cur_patch) {
        return 0,1
    }
    if(visited[i][j] == true){
        return 0,0
    }

    visited[i][j] = true
    downA, downP := calculate_area_for_region(i-1, j, cur_patch, the_map, visited)
    upA, upP:= calculate_area_for_region(i + 1, j, cur_patch, the_map, visited)
    leftA, leftP := calculate_area_for_region(i, j - 1, cur_patch, the_map, visited)
    rightA, rightP := calculate_area_for_region(i, j + 1, cur_patch, the_map, visited)

    return (1 + leftA + rightA + upA + downA), (downP + upP + leftP + rightP )
}

func main(){
    the_map := load_the_map()
    visited_map := create_visited_map(the_map)

    area_map := map[string]int{}
    perimeter_map := map[string]int{}
    total := 0
    for i := 0; i < len(the_map); i++ {
        for j := 0; j < len(the_map[0]); j++ {
            area, perimeter:= calculate_area_for_region(i, j, the_map[i][j], the_map, visited_map)
            area_map[the_map[i][j]] += area
            perimeter_map[the_map[i][j]] += perimeter
            total += area * perimeter
        }
    }
    fmt.Println(area_map)
    fmt.Println(perimeter_map)
    fmt.Println(total)
}
