# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
  def mergeTwoLists(self, list1, list2):
    cur = new = ListNode(0)
    while list1 and list2:
      print(list1.val)
      if list1.val > list2.val:
        cur.next = list2
        list2 = list2.next
      else: 
        cur.next = list1
        list1 = list1.next
      cur = cur.next
    cur.next = list1 or list2
    print(cur.val, new.val)
  
  # recursively    
def mergeTwoLists2(self, l1, l2):
    if not l1 or not l2:
        return l1 or l2
    if l1.val < l2.val:
        l1.next = self.mergeTwoLists(l1.next, l2)
        return l1
    else:
        l2.next = self.mergeTwoLists(l1, l2.next)
        return l2