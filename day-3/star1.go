package main
import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strconv"
)

func main(){
    data, _ := ioutil.ReadFile("/Users/cadenmilne/programming/advent-of-code/day-3/input.txt")
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
    fmt.Println(total)
}
