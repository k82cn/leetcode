use std::collections::HashMap;

#[derive(Debug)]
struct Solution {
}

impl Solution {
    pub fn sum(n: i32) -> i32 {
        let mut m = n;
        let mut s = 0;

        loop {
            let g = m%10;
            s = s + g*g;
            m = m/10;
            if m == 0 {
                return s;
            }
        }
    }

    pub fn is_happy(n: i32) -> bool {
        let mut c = HashMap::new();
        let mut m = n;

        loop {
            m = Self::sum(m);

            if m == 1 {
                return true;
            } 

            if c.get(&m).is_none() {
                c.insert(m, m);
            } else {
                return false;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_is_happy() {
        let cases = vec![(19, true), (2, false)];

        for (n, t) in cases.iter() {
            let result = Solution::is_happy(*n);    
            assert_eq!(result, *t);
        }
    }
}
