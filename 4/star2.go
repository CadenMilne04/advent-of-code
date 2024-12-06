package main

import (
	"bufio"
	"fmt"
	"os"
)

func isXmas(i int, j int, crossword [][]string) int {
    /* Check the 3x3 grid for an X-Mas */
    w1, w2 := "", ""
    for k := range 3{
        if !(i+k < 0 || i+k > len(crossword)-1 || j+k < 0 || j+k > len(crossword[0])-1){
            w1 += crossword[i+k][j+k]
        }
        if !(i+2-k < 0 || i+2-k > len(crossword)-1 || j+k < 0 || j+k > len(crossword[0])-1){
            w2 += crossword[i+2-k][j+k]
        }
    }
    if (w1 == "MAS" || w1 == "SAM") && (w2 == "MAS" || w2 == "SAM"){
    println(i, j, w1,w2)
        return 1
    }

    return 0
}

func main(){
    file, _ := os.Open("/Users/cadenmilne/programming/advent-of-code/day-4/input.txt")
    scanner := bufio.NewScanner(file)
    var crossword [][]string
    for scanner.Scan(){
        text := scanner.Text()
        var cw_line []string
        for _, char := range text {
            cw_line = append(cw_line, string(char))
        }
        crossword = append(crossword, cw_line)
    }

    total := 0
    for i := 0; i < len(crossword); i++{
        for j := 0; j < len(crossword[0]); j++{
            total += isXmas(i, j , crossword)
        }
    }
    fmt.Println("Total: ", total)
}
