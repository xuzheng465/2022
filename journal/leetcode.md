```python
class ListNode:
  def __init__(self, val=0, next=None):
    self.val = val
    self.next = next
```

[206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)

```python
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None:
            return head
        
        pre = None
        cur = head
        while cur:
            next_ = cur.next
            cur.next = pre
            pre = cur
            cur = next_

        return pre
```
