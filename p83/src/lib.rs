// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub struct Solution{}

impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head.clone();
        let mut curr_opt = head.as_mut();

        while let Some(cur) = curr_opt {
            let mut next_opt = cur.next.take();

            while let Some(next) = next_opt.as_mut() {
                if cur.val == next.val {
                    next_opt = next.next.take();
                } else {
                    cur.next = next_opt;
                    break;
                }
            }

            curr_opt = cur.next.as_mut();
        }

        head
    }
}
