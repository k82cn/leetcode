pub struct Solution {}

impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut m = 1;
        let mut result = vec![];

        for d in digits.iter().rev() {
            result.push((*d + m) % 10);
            m = (*d + m) / 10;
        }

        if m != 0 {
            result.push(m);
        }

        result.reverse();

        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        struct TestCase {
            input: Vec<i32>,
            expected: Vec<i32>,
        }

        let cases = vec![
            TestCase {
                input: vec![1, 2, 3],
                expected: vec![1, 2, 4],
            },
            TestCase {
                input: vec![9],
                expected: vec![1, 0],
            },
        ];

        for c in cases {
            assert_eq!(c.expected, Solution::plus_one(c.input));
        }
    }
}
