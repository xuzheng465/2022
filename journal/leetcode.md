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



07-07

[19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

```python
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        if not head:
            return head
        dummy = ListNode(-1)
        dummy.next = head
        pre = dummy
        count = 0
        cur = head
        while cur:
            cur = cur.next
            count += 1
        c = 0
        while c < (count-n):
            pre = pre.next
            c+=1
        pre.next = pre.next.next
        return dummy.next
```



```python
# 双指针解法

```





