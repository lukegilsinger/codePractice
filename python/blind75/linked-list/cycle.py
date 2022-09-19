# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from opcode import hascompare


def hasCycle(self, head) -> bool:
  hs = set()
  while head:
    if head in hs: return True
    hs.add(head)
    head = head.next
  return False

# O(1)
def hasCycle2(self, head) -> bool:
  hs = set()
  while head:
    if head in hs: return True
    hs.add(head)
    head = head.next
  return False

head = [3,2,0,-4]
pos = 1
print(hascompare(head))