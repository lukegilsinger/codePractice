def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
  p = head
  while p:
    while p.next and p.next.val == p.val:
      p.next = p.next.next
    p = p.next
  # if head and head.val == val: head = head.next
  return head