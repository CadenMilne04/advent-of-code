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

func main()  {
    the_map := load_the_map()
    visited := initialize_visited(len(the_map), len(the_map[0]))
    gX, gY := find_the_guard(the_map)
    fmt.Println(gX, gY, the_map[gX][gY])
    fmt.Println(visited)
    /* While guard is inboud, move forward once add to count if not visited, visit, turn right if needed */
}
