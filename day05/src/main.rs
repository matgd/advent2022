fn main() {
    let input = include_str!("input.txt");
    println!("[Part 1]: {}", day5(input, 1));
    println!("[Part 2]: {}", day5(input, 2));
}

fn day5(input: &str, part: u8) -> String {
    let lines = input.lines();
    let mut total_stacks = 0;
    let mut stacks: Vec<Vec<char>> = Vec::new();
    let mut parse_stack_definition = true;

    lines.for_each(|line| {
        // No more stacks definition.
        if parse_stack_definition && !line.contains("[") { 
            parse_stack_definition = false;

            // Reverse for proper order.
            stacks.iter_mut().for_each(|stack| {
                stack.reverse();
            });
        }

        if parse_stack_definition {
            if total_stacks == 0 {
                total_stacks = (line.len() + 1) / 4;
                for _ in 0..total_stacks {
                    stacks.push(Vec::new());
                }
            }

            for i in 0..total_stacks {
                let stack_box = line.chars().skip(i * 4 + 1).nth(0).unwrap();
                if stack_box != ' ' {
                    stacks[i].push(stack_box);
                }
            }
        } else if line.starts_with("move") {
            let split = line.split_whitespace().collect::<Vec<&str>>();

            let mv = split[1].parse::<usize>().unwrap();
            let fr = split[3].parse::<usize>().unwrap() - 1;
            let to = split[5].parse::<usize>().unwrap() - 1;
            
            if part == 1 {
                for _ in 0..mv {
                    let popped = stacks[fr].pop().unwrap();
                    stacks[to].push(popped);
                }
            } else {
                let mut temp = Vec::new();
                for _ in 0..mv {
                    temp.push(stacks[fr].pop().unwrap());
                }
                temp.reverse();
                for i in 0..mv {
                    stacks[to].push(temp[i]);
                }
            }
        }
    });
        
    let mut result = String::new();
    stacks.iter().for_each(|stack| {
        result.push(*stack.last().unwrap());
    });

    result
}
