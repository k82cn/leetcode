struct Solution {}

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        let mut a_charts = a.chars().rev();
        let mut b_charts = b.chars().rev();
        let mut result = String::new();
        let mut carry = false;

        let to_value = |v| match v {
            None => 0,
            Some('0') => 0,
            Some('1') => 1,
            Some(x) => panic!("unexpected char {x}"),
        };

        loop {
            match (a_charts.next(), b_charts.next()) {
                (None, None) => {
                    if carry {
                        result.push('1');
                    }
                    break;
                }
                (a_ch, b_ch) => {
                    let av = to_value(a_ch);
                    let bv = to_value(b_ch);

                    let mut sum = av + bv;
                    if carry {
                        sum = sum + 1;
                    }

                    match sum {
                        0 => {
                            result.push('0');
                            carry = false;
                        }
                        1 => {
                            result.push('1');
                            carry = false;
                        }
                        2 => {
                            result.push('0');
                            carry = true;
                        }
                        3 => {
                            result.push('1');
                            carry = true;
                        }

                        _ => panic!("unexpected sum"),
                    }
                }
            }
        }

        result.chars().rev().collect::<String>()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_binary() {
        struct TestCase<'a> {
            a: &'a str,
            b: &'a str,
            o: &'a str,
        }

        let cases = vec![
            TestCase {
                a: "11",
                b: "1",
                o: "100",
            },
            TestCase {
                a: "1010",
                b: "1011",
                o: "10101",
            },
        ];

        for c in cases {
            assert_eq!(Solution::add_binary(c.a.to_string(), c.b.to_string()), c.o);
        }
    }
}
