def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
  hm = set()
  def dfs(root):
    if not root: return False
    if root.val in hm: return True
    hm.add(k-root.val)
    return dfs(root.left) or dfs(root.right)

  return dfs(root)