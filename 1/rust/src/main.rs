use std::fs;
use std::num;


fn main() {
    let mut right_list = Vec::new();
    let mut left_list = Vec::new();

    let file = fs::read_to_string("/Users/cadenmilne/programming/advent-of-code/day-1/input.txt").expect("Should have been able to read this file...");
    for line in file.lines() {
        let mut split = line.split_whitespace();
        let r_num = split.next().expect("There should always be a right value");
        let l_num = split.next().expect("There should always be a left value");
        right_list.push(r_num.parse::<i32>().unwrap());
        left_list.push(l_num.parse::<i32>().unwrap());
    }
    right_list.sort();
    left_list.sort();
    println!("Right List: {:?}", right_list);
    println!("Left List: {:?}", left_list);

    let mut total = 0;
    let n: usize = right_list.len();
    for i in 0..n {
        let diff = right_list[i] - left_list[i];
        total += diff.abs();
    }
    println!("Total: {}", total);
}
