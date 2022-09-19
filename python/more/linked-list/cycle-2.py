def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
  hs = set()
  while head:
    if head in hs: return head
    print(head)
    hs.add(head)
    head = head.next
  return None

# 2 pointer
def detectCycle2(self, head: Optional[ListNode]) -> Optional[ListNode]:
  slow = fast = head
  while fast and fast.next:
      slow, fast = slow.next, fast.next.next
      if slow == fast: break
  else: return None  # if not (fast and fast.next): return None
  while head != slow:
      head, slow = head.next, slow.next
  return head