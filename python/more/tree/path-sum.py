def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
  if not root: return False
  def dfs(root, curr):
    if root:
      curr += root.val
      if not root.left and not root.right and curr == targetSum: return True
      return dfs(root.left, curr) or dfs(root.right, curr)
    print(curr)
    return False
  return dfs(root, 0)