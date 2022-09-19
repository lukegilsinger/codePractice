# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
  def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
    c = 0
    hm = {}
    while head:
      hm[c] = head
      head = head.next
      c += 1
      print(c, head)
    return hm[c // 2]
    
  def middleNode2(self, head: Optional[ListNode]) -> Optional[ListNode]:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow

head = [1,2,3,4,5]
sol = Solution()
sol.middleNode(head)