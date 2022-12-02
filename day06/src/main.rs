use std::collections::HashSet;

fn main() {
    let input = include_str!("input.txt");
    println!("Day 6, Part 1: {}", day6(input, 1));
    println!("Day 6, Part 2: {}", day6(input, 2));
}

fn day6(input: &str, part: u8) -> usize {
    let sequence_len = { if part == 1 { 4 } else { 14 } };
    for i in 0..input.len()-sequence_len {
        let gathered = input
            .chars()
            .skip(i)
            .take(sequence_len)
            .collect::<HashSet<char>>();
        if gathered.len() == sequence_len {
            return i + sequence_len;
        }
    }
    return 0;
}

