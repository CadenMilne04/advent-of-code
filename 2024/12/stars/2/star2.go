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

func checkIfExternalCorner(i int, j int, di int, dj int, cur_patch string, the_map [][]string, corners [][]bool) int {
    if(j+dj < 0 || j+dj >= len(the_map[0]) || the_map[i][j+dj] != cur_patch){
        if(i+di < 0 || i+di >= len(the_map) || the_map[i+di][j] != cur_patch){
            if(i+di < 0 || i+di >= len(the_map) || j+dj < 0 || j+dj >= len(the_map[0])){
                return 1
            }
            if (the_map[i+di][j+dj] != cur_patch && corners[i+di][j+dj] == false){
                corners[i+di][j+dj] = true
                return 1
            }
        }
    }
    return 0
}

func checkIfInternalCorner(i int, j int, di int, dj int, cur_patch string, the_map [][]string, corners [][]bool) int {

    if(i+di >= 0 && i+di < len(the_map) && j+dj >= 0 && j+dj < len(the_map[0])){
        if(the_map[i+di][j] == cur_patch && the_map[i][j+dj] == cur_patch){
            if(i+di < 0 || i+di >= len(the_map) || j+dj < 0 || j+dj >= len(the_map[0]) || the_map[i+di][j+dj] != cur_patch && corners[i+di][j+dj] == false){
                corners[i+di][j+dj] = true
                return 1
            }
        }
    }
    return 0
}

func calculate_area_for_region(i int, j int, cur_patch string, the_map [][]string, visited [][]bool, corners [][]bool) (int, int) {
    if(i < 0 || i >= len(the_map) || j < 0 || j >= len(the_map[0]) || the_map[i][j] != cur_patch){
        return 0, 0
    }
    if(visited[i][j] == true){
        return 0,0
    }

    visited[i][j] = true
    /* Using the perimeters from each dir, we can see if we are at an edge and then we will know if we are at a corner*/
    cornersHere := 0
    cornersHere += checkIfExternalCorner(i, j, 1, 1, cur_patch, the_map, corners)
    cornersHere += checkIfExternalCorner(i, j, 1, -1, cur_patch, the_map, corners)
    cornersHere += checkIfExternalCorner(i, j, -1, 1, cur_patch, the_map, corners)
    cornersHere += checkIfExternalCorner(i, j, -1, -1, cur_patch, the_map, corners)
    cornersHere += checkIfInternalCorner(i, j, 1, 1, cur_patch, the_map, corners)
    cornersHere += checkIfInternalCorner(i, j, 1, -1, cur_patch, the_map, corners)
    cornersHere += checkIfInternalCorner(i, j, -1, 1, cur_patch, the_map, corners)
    cornersHere += checkIfInternalCorner(i, j, -1, -1, cur_patch, the_map, corners)
    downA, downP := calculate_area_for_region(i-1, j, cur_patch, the_map, visited, corners)
    upA, upP:= calculate_area_for_region(i + 1, j, cur_patch, the_map, visited, corners)
    leftA, leftP := calculate_area_for_region(i, j - 1, cur_patch, the_map, visited, corners)
    rightA, rightP := calculate_area_for_region(i, j + 1, cur_patch, the_map, visited, corners)

    return (1 + leftA + rightA + upA + downA), (cornersHere + downP + upP + leftP + rightP)
}

func main(){
    the_map := load_the_map()
    visited_map := create_visited_map(the_map)
    corners := create_visited_map(the_map)

    area_map := map[string]int{}
    perimeter_map := map[string]int{}
    total := 0
    for i := 0; i < len(the_map); i++ {
        for j := 0; j < len(the_map[0]); j++ {
            area, perimeter:= calculate_area_for_region(i, j, the_map[i][j], the_map, visited_map, corners)
            area_map[the_map[i][j]] += area
            perimeter_map[the_map[i][j]] += perimeter
            total += area * perimeter
        }
    }
    fmt.Println(area_map)
    fmt.Println(perimeter_map)
    fmt.Println(total)
}
