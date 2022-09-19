# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
def reverseList(head: ListNode) -> ListNode:
  prev, curr = None, head
  while curr:
    print(curr.val)
    nxt = curr.next
    curr.next = prev
    prev = curr
    curr = nxt
  return prev

def reverseListRec(head: ListNode) -> ListNode:
  def rec(node, prev=None):
    if not node:
      return prev
    print(node.val)
    n = node.next
    node.next = prev
    return rec(n, node)
  return rec(head)

head = [1,2,3,4,5]
print(reverseList(head))
print(reverseListRec(head))
