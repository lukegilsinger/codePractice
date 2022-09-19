# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
  def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
    def dfs(root, p, q):
      print(root.val)
      if not root: return None
      if p.val > root.val and q.val > root.val:
        dfs(root.right, p, q)
      elif p.val < root.val and q.val < root.val:
        dfs(root.left, p, q)
      else: return root
    return dfs(root, p, q)