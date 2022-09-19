def searchBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
  def dfs(root):
    if(root):
      print(root.val)
      if root.val == val: return root
      elif root.val > val:
        return dfs(root.left)
      elif root.val < val:
        return dfs(root.right)
    return None
  return dfs(root)