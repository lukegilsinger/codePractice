# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# The left subtree of a node contains only nodes with keys less than the node's key.
# The right subtree of a node contains only nodes with keys greater than the node's key.
# Both the left and right subtrees must also be binary search trees.

# class Solution:
def isValidBST(self, root: Optional[TreeNode]) -> bool:
  def dfs(r, mini, maxi):
    print(r)
    if (r.val <= mini or r.val >= maxi): return False
    if r.left:
      if (dfs(r.left, mini, min(maxi, r.val)) == False): return False
    if r.right:
      if (dfs(r.right, max(mini, r.val), maxi) == False): return False
    return True
  return dfs(root, -9999, 9999)
  
