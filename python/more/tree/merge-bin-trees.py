# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

def mergeTrees(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> Optional[TreeNode]:
  def dfs(r1, r2):
    if r1 and r2:
      print(r1.val, r2.val, r1.val+r2.val)
      root = TreeNode(r1.val+r2.val)
      root.left = dfs(r1.left, r2.left)
      root.right = dfs(r1.right, r2.right)
      return root
    else: return r1 or r2
  return dfs(root1,root2)

root1 = [1,3,2,5]
root2 = [2,1,3,null,4,null,7]
mergeTrees(root1, root2)