fn main() {
    let input = include_str!("input1.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> i32 {
    let mut r:i32 = 0;
    let numbers = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];
    for line in input.lines() {
        let mut first = 0;
        let mut last = 0;
        'out: for i in 0..line.len() {
            if line.chars().nth(i).unwrap().is_digit(10) {
                first = line.chars().nth(i).unwrap().to_digit(10).unwrap() as i32;
                break 'out;
            }
            for (j, number) in numbers.iter().enumerate() {
                if line.len() >= i+number.len() && line[i..i+number.len()] == **number {
                   first = j as i32 + 1;
                   break 'out;
                }
            }
        }
        'out: for i in (0..line.len()).rev() {
            if line.chars().nth(i).unwrap().is_digit(10) {
                last = line.chars().nth(i).unwrap().to_digit(10).unwrap() as i32;
                break 'out;
            }
            for (j, number) in numbers.iter().enumerate() {
                if line.len() >= i+number.len() && line[i..i+number.len()] == **number {
                   last = j as i32 + 1;
                   break 'out;
                }
            }
        }
        r += (first * 10) + last;
    }
    r
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn part2_test() {
        let result = part2(
            "
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
",
        );

        assert_eq!(result, 281);
    }
}
