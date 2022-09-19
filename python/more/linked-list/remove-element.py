def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
  p = head
  while p:
    while p.next and p.next.val == val:
      p.next = p.next.next
    p = p.next
  if head and head.val == val: head = head.next
  return head