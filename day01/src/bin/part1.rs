fn main() {
    let input = include_str!("input1.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> i32 {
    let mut r = 0;
    for line in input.split_whitespace() {
        let digits: Vec<char> = line.trim().chars().filter(|c| c.is_digit(10)).collect();
        let first_and_last: Vec<char> = vec![digits[0], digits[digits.len() - 1]];
        r = r + String::from_iter(first_and_last).parse::<i32>().unwrap();
    }
    return r;
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]

    fn part1_test() {
        let result = part1(
            "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet",
        );
        assert_eq!(result, 142);
    }
}
