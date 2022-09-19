# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
  def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
    res = []
    def dfs(root):
      if(root):
        res.append(root.val)
        dfs(root.left)
        dfs(root.right)
    dfs(root)
    return res
  
  def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
    res = []
    def dfs(root):
      if(root):
        dfs(root.left)
        res.append(root.val)
        dfs(root.right)
    dfs(root)
    return res
  
  def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
    res = []
    def dfs(root):
      if(root):
        dfs(root.left)
        dfs(root.right)
        res.append(root.val)
    dfs(root)
    return res