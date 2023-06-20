// Definition for a binary tree node.

use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution {}

impl Solution {
    pub fn min_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut dep = 0;
        let mut cur_level = vec![];

        if let Some(r) = root {
            cur_level.push(r.clone());
        }

        while !cur_level.is_empty() {
            let mut next_level = vec![];
            for n in cur_level {
                match (&n.borrow().left, &n.borrow().right) {
                    (None, None) => {
                        return dep + 1;
                    }
                    (left, right) => {
                        if let Some(l) = left {
                            next_level.push(l.clone());
                        }

                        if let Some(r) = right {
                            next_level.push(r.clone());
                        }
                    }
                }
            }

            cur_level = next_level;
            dep = dep + 1;
        }

        dep
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
