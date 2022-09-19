# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# level order needs to be breath first search
# Use a queue and interative for BFS

from collections import deque


def levelOrder(self, root: 'Node') -> List[int]:
  queue, res = deque([root] if root else []), []

  while len(queue):
    qlen, row = len(queue), []
    for i in range(qlen):
      curr = queue.popleft()
      row.append(curr.val)
      if curr.left: queue.append(curr.left)
      if curr.right: queue.append(curr.right)
    res.append(row)
  return res

