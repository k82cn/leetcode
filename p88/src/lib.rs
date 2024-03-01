struct Solution {}

impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut ptr = nums1.len();

        let mut m_ptr = m as usize;
        let mut n_ptr = n as usize;

        while m_ptr > 0 && n_ptr > 0 {
            if nums1[m_ptr - 1] > nums2[n_ptr - 1] {
                nums1[ptr - 1] = nums1[m_ptr - 1];
                m_ptr -= 1;
                ptr -= 1;
            } else if nums1[m_ptr - 1] < nums2[n_ptr - 1] {
                nums1[ptr - 1] = nums2[n_ptr - 1];
                n_ptr -= 1;
                ptr -= 1;
            } else {
                nums1[ptr - 1] = nums2[n_ptr - 1];
                n_ptr -= 1;
                ptr -= 1;

                nums1[ptr - 1] = nums1[m_ptr - 1];
                m_ptr -= 1;
                ptr -= 1;
            }
        }

        while m_ptr > 0 {
            nums1[ptr - 1] = nums1[m_ptr - 1];
            m_ptr -= 1;
            ptr -= 1;
        }

        while n_ptr > 0 {
            nums1[ptr - 1] = nums2[n_ptr - 1];
            n_ptr -= 1;
            ptr -= 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        struct TestCase {
            pub nums1: Vec<i32>,
            pub m: i32,
            pub nums2: Vec<i32>,
            pub n: i32,
            pub expected: Vec<i32>,
        }

        let cases = vec![TestCase {
            nums1: vec![1, 2, 3, 0, 0, 0],
            m: 3,
            nums2: vec![2, 5, 6],
            n: 3,
            expected: vec![1, 2, 2, 3, 5, 6],
        }];

        for mut c in cases {
            Solution::merge(&mut c.nums1, c.m, &mut c.nums2, c.n);
            assert_eq!(c.nums1, c.expected);
        }
    }
}
