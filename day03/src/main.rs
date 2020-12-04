use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut vec: Vec<Vec<char>> = Vec::new();
    if let Ok(lines) = read_lines("input/input.txt") {
        for line in lines {
            if let Ok(li) = line {
                vec.push(String::from(li).chars().collect());
            }
        }
    }
    println!("{:?}", solve_part_one(&vec));
    println!("{:?}", solve_part_two(&vec))
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn solve_part_one(grid: &Vec<Vec<char>>) -> i64 {
    return solve(&grid, 3, 1)
}

fn solve(grid: &Vec<Vec<char>>, right:usize, down:usize) -> i64 {
    let mut x = 0;
    let mut y = 0;
    let mut count = 0;
    let nb_lines = grid.len();
    let nb_col = grid[0].len();
    println!("lines {:?}, col {:?}", nb_lines, nb_col);
    while y < nb_lines {
        x = (x+right)%nb_col;
        y += down;
        if y >= nb_lines {
            break
        }
        if grid[y][x] == '#' {
            count += 1;
        }
    }
    return count
}

fn solve_part_two(grid: &Vec<Vec<char>>) -> i64 {
    return solve(&grid, 1, 1) * solve(&grid, 3, 1) * solve(&grid, 5, 1) * solve(&grid, 7, 1) * solve(&grid, 1, 2)
}