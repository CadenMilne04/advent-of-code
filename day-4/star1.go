package main

import (
	"bufio"
	"fmt"
	"os"
)
func findWordVertical(i int, j int, word string, crossword [][]string) int{
    total := 0
    w1, w2, w3, w4, w5, w6, w7, w8 := "", "", "", "", "", "", "", ""
    for k := range 4 {
        if !(i+k < 0 || i+k > len(crossword)-1 || j < 0 || j > len(crossword[0])-1){
            w1 += crossword[i + k][j]
        }
        if !(i-k < 0 || i-k > len(crossword)-1 || j < 0 || j > len(crossword[0])-1){
            w2 += crossword[i - k][j]
        }
        if !(i < 0 || i > len(crossword)-1 || j+k < 0 || j+k > len(crossword[0])-1){
            w3 += crossword[i][j + k]
        }
        if !(i < 0 || i > len(crossword)-1 || j-k < 0 || j-k > len(crossword[0])-1){
            w4 += crossword[i][j - k]
        }
        if !(i+k < 0 || i+k > len(crossword)-1 || j+k < 0 || j+k > len(crossword[0])-1){
            w5 += crossword[i+k][j+k]
        }
        if !(i+k < 0 || i+k > len(crossword)-1 || j-k < 0 || j-k > len(crossword[0])-1){
            w6 += crossword[i+k][j-k]
        }
        if !(i-k < 0 || i-k > len(crossword)-1 || j+k < 0 || j+k > len(crossword[0])-1){
            w7 += crossword[i-k][j+k]
        }
        if !(i-k < 0 || i-k > len(crossword)-1 || j-k < 0 || j-k > len(crossword[0])-1){
            w8 += crossword[i-k][j-k]
        }
    }
    if(w1 == "XMAS") {total += 1}
    if(w2 == "XMAS") {total += 1}
    if(w3 == "XMAS") {total += 1}
    if(w4 == "XMAS") {total += 1}
    if(w5 == "XMAS") {total += 1}
    if(w6 == "XMAS") {total += 1}
    if(w7 == "XMAS") {total += 1}
    if(w8 == "XMAS") {total += 1}
    return total
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

    visited := make([][]bool, len(crossword))
    for i := range visited {
        visited[i] = make([]bool, len(crossword[0]))
    }

    total := 0
    for i := 0; i < len(crossword); i++{
        for j := 0; j < len(crossword[0]); j++{
            total += findWordVertical(i, j, "", crossword)
        }
    }
    fmt.Println("Total: ", total)
}
