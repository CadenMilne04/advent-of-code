package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func load_the_map() [][]string {
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/6/input.txt")
    scanner := bufio.NewScanner(file)

    the_map := [][]string{}

    for scanner.Scan(){
        line := scanner.Text();
        chars := strings.Split(line, "")
        row := []string{}
        for _, char := range chars{
            row = append(row, char)
        }
        the_map = append(the_map, row)
    }

    return the_map
}

func find_the_guard(the_map [][]string) (int, int) {
    for i := range len(the_map) {
        for j, char := range the_map[i]{
            if(strings.Contains("<^>v", char)){
                return i, j
            }
        }
    }
    return -1, -1
}

func get_new_direction(dx int, dy int) (int, int) {
    if (dx == 1){
        return 0, -1
    } else if (dx == -1) {
        return 0, 1
    } else if (dy == 1) {
        return 1, 0
    } else {
        return -1, 0
    }
}

func create_visited_arrays(the_map [][]string) ([][]int, [][]int){
    
    vdx := make([][]int, len(the_map))
    for i := range vdx {
        vdx[i] = make([]int, len(the_map[0]))
    }
    vdy := make([][]int, len(the_map))
    for i := range vdy {
        vdy[i] = make([]int, len(the_map[0]))
    }
    for i := 0; i < len(vdx); i++ {
        for j := 0; j < len(vdx[0]); j++ {
            vdx[i][j] = -2
            vdy[i][j] = -2
        }
    }
    
    return vdx, vdy
}
func copyMap(original [][]string) [][]string {
    copy := make([][]string, len(original))
    for i := range original {
        copy[i] = append([]string{}, original[i]...) // Create a deep copy of each row
    }
    return copy
}


func check_if_in_loop(x int, y int, dx int, dy int, the_map [][]string) bool {
    vdx, vdy := create_visited_arrays(the_map)
    the_new_map := copyMap(the_map)

    /* Put an obstacle in front of the guard */
    if(x + dx >= 0 && x + dx < len(the_new_map) && y + dy >= 0 && y + dy < len(the_new_map[0]) && the_new_map[x + dx][y + dy] != "#"){
        the_new_map[x + dx][y + dy] = "#"
    }else{
        return false
    }

    /* Reset guard to position at the beginning */
    x, y = find_the_guard(the_map)
    dx, dy = -1, 0

    for(x >= 0 && x < len(the_new_map) && y >= 0 && y < len(the_new_map[0])){
        /* If we've been at this square going the same direction we are in a loop */

        /* Check if the next square will be an obstacle */
        if(x + dx >= 0 && x + dx < len(the_new_map) && y + dy >= 0 && y + dy < len(the_new_map[0]) && the_new_map[x + dx][y + dy] == "#"){
            /* If So Change direction */
            dx, dy = get_new_direction(dx, dy)
            if(vdx[x+dx][y+dy] == dx && vdy[x+dx][y+dy] == dy){
                return true
            }
            vdx[x+dx][y+dy] = dx
            vdy[x+dx][y+dy] = dy
            continue
        }

        x += dx
        y += dy
    }
    return false
}

func count_possible_obstacles (x int, y int, dx int, dy int, the_map [][]string) int {
    total := 0
    total_tries := 0
    visiteds := 0

    visited := make([][]bool, len(the_map))
    for i := range visited {
        visited[i] = make([]bool, len(the_map[0]))
    }

    visited2  := make([][]bool, len(the_map))
    for i := range visited2 {
        visited2[i] = make([]bool, len(the_map[0]))
    }
    for(x >= 0 && x < len(the_map) && y >= 0 && y < len(the_map[0])){
        if(x + dx >= 0 && x + dx < len(the_map) && y + dy >= 0 && y + dy < len(the_map[0]) && the_map[x + dx][y + dy] == "#"){
            /* If Obstacle Change direction */
            dx, dy = get_new_direction(dx, dy)
            continue
        }
        total_tries++
        if (check_if_in_loop(x, y, dx, dy, the_map)){
            if(visited[x + dx][y + dy] == false){
                total++
                visited[x + dx][y + dy] = true
            }
        }
        if !visited2[x][y]{
            visiteds++
            visited2[x][y] = true
        }

        x += dx
        y += dy
    }
    fmt.Println(total_tries)
    fmt.Println(visiteds)
    return total
}

func main()  {
    the_map := load_the_map()
    gX, gY := find_the_guard(the_map)
    /* While guard is inbouds, put object in front and see if he loops */
    total_objects := count_possible_obstacles(gX, gY, -1, 0, the_map)
    fmt.Println("Total objects that created loops: ", total_objects)
    fmt.Println(check_if_in_loop(gX, gY, -1, 0, the_map))
}
