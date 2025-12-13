package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func load_the_map() [][]string {
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/8/input.txt")
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

type location struct {
    x int
    y int
}

func create_frequency_map(the_map [][]string) map[string][]location {
    frequency_map := make(map[string][]location)
    for i := 0; i < len(the_map); i++ {
        for j := 0; j < len(the_map[0]); j++ {
            char := the_map[i][j]
            if(char != "."){
                var loc location
                loc.x = i
                loc.y = j
                frequency_map[char] = append(frequency_map[char], loc)
            }
        }
    }
    return frequency_map
}

func create_visited_map(the_map [][]string) [][]bool {
    var visited_map [][]bool
    for i := 0; i < len(the_map); i++ {
        row := make([]bool, len(the_map[0]))
        visited_map = append(visited_map, row)
    }
    return visited_map
}

func create_antinodes_at_locations(locA location, locB location, dx int, dy int, visited_map [][]bool) int {
    total := 0
    /* 4 Possibilities */
    var possibilities []location
    a := location{x: locA.x - dx, y: locA.y - dy}
    possibilities = append(possibilities, a)
    a2 := location{x: locA.x + dx, y: locA.y + dy}
    possibilities = append(possibilities, a2)
    b := location{x: locB.x - dx, y: locB.y - dy}
    possibilities = append(possibilities, b)
    b2 := location{x: locB.x + dx, y: locB.y + dy}
    possibilities = append(possibilities, b2)

    for _, loc := range possibilities {
        /* check if it is not the same as one of the locations */
        if!((loc.x == locA.x && loc.y == locA.y) || (loc.x == locB.x && loc.y == locB.y)){
            /* ensure inbounds and not already visited */
            if(loc.x >= 0 && loc.x < len(visited_map) && loc.y >= 0 && loc.y < len(visited_map[0]) && !visited_map[loc.x][loc.y]){
                visited_map[loc.x][loc.y] = true
                total++
            }
        }
    }
    return total
}

func count_antinodes(frequency_map map[string][]location, visited_map [][]bool) int {
    /* loop through the freq map with all combos of two (n^2) satellites, calculating a total and populating the visited map */
    total := 0
    for _, locs := range frequency_map {
        for i := 0; i < len(locs); i++ {
            locA := locs[i]
            for j := i + 1; j < len(locs); j++ {
                locB := locs[j]
                dx := locA.x - locB.x
                dy := locA.y - locB.y
                num_antinodes_created := create_antinodes_at_locations(locA, locB, dx, dy, visited_map)
                total += num_antinodes_created
            }
        }
    }
    return total
}

func main() {
    the_map := load_the_map()
    frequency_map := create_frequency_map(the_map)
    visited_map := create_visited_map(the_map)
    num_antinodes := count_antinodes(frequency_map, visited_map)
    fmt.Println(num_antinodes)
}
