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

func initialize_visited(w int, h int) [][]bool{
    visited := [][]bool{}
    for i := 0; i < w; i++{
        visited = append(visited, make([]bool, h))
    }
    return visited
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

func count_guard_visited (x int, y int, dx int, dy int, the_map [][]string, visited [][]bool) int {
    if(x < 0 || x > len(the_map) - 1 || y < 0 || y > len(the_map[0]) - 1){
        return 0;
    } else {
        new_square := 0
        if(visited[x][y] == false){
            new_square = 1
        }
        visited[x][y] = true
        /* Check if the next square will be an obstacle */
        if(x + dx > 0 && x + dx < len(the_map) -1 && y + dy > 0 && y + dy < len(the_map[0]) && the_map[x + dx][y + dy] == "#"){
            /* Change direction */
            dx, dy = get_new_direction(dx, dy)
        }
        the_map[x][y] = "X"
        return new_square + count_possible_obstacles(x + dx, y + dy, dx, dy, the_map, visited)
    }
}

func print_the_map(the_map [][]string){
    for i := range len(the_map){
        for j := range len(the_map[0]){
            fmt.Print(the_map[i][j])
        }
        fmt.Print("\n")
    }
}

func main()  {
    the_map := load_the_map()
    visited := initialize_visited(len(the_map), len(the_map[0]))
    gX, gY := find_the_guard(the_map)
    fmt.Println(gX, gY, the_map[gX][gY])
    /* While guard is inboud, move forward once add to count if not visited, visit, turn right if needed */
    total_visted := count_possible_obstacles(gX, gY, -1, 0, the_map, visited)
    print_the_map(the_map)
    fmt.Println(total_visted)
}
