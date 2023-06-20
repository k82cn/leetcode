struct Solution {}

impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        let mut m: i64 = 1;
        let x = x as i64;
        while m < 46340 && m * m < x {
            m = m * 2;
        }

        m = m / 2;

        while m * m < x {
            m = m + 1;
        }

        if m * m > x {
            return (m - 1) as i32;
        }

        m as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_my_sqrt() {
        struct TestCase {
            input: i32,
            output: i32,
        }

        let cases = vec![
            TestCase {
                input: 4,
                output: 2,
            },
            TestCase {
                input: 2147395599,
                output: 46339,
            },
            TestCase {
                input: 2147483647,
                output: 46340,
            },
            TestCase {
                input: 8,
                output: 2,
            },
        ];

        for c in cases {
            assert_eq!(Solution::my_sqrt(c.input), c.output);
        }
    }
}
