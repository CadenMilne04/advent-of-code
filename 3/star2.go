package main
import (
    "fmt"
    "os"
    "regexp"
    "strconv"
)
func calcTotal(data string) int{
    text := []byte(data)
    re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
    matches := re.FindAllStringSubmatch(string(text), -1)
    total := 0
    for _, match := range matches {
        fmt.Println(match[0], match[1], match[2])
        num1, _ := strconv.Atoi(match[1])
        num2, _ := strconv.Atoi(match[2])
        total += num1 * num2
    }
    return total
}

func main(){
    data, _ := os.ReadFile("/Users/cadenmilne/programming/advent-of-code/day-3/input.txt")
    text := []byte(data)
    total := 0
    rea := regexp.MustCompile(`([\s\S]*?)don't\(\)`)
    matchesa := rea.FindAllStringSubmatch(string(text), -1)
    fmt.Println((matchesa[0][1]))
    total += calcTotal(matchesa[0][1])
    re := regexp.MustCompile(`do\(\)([\s\S]*?)don't\(\)`)
    matches := re.FindAllStringSubmatch(string(text), -1)
    for _, match := range matches {
        fmt.Println(match[1])
        total += calcTotal(match[1])
    }
    fmt.Println(total)
}
