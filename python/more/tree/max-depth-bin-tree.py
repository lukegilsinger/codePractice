
def maxDepth(self, root: Optional[TreeNode]) -> int:
  def dfs(root):
    if root: return max(dfs(root.left),dfs(root.right)) + 1
    else: return 0
  return dfs(root)

root = [3,9,20,null,null,15,7]