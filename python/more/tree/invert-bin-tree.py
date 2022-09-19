def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
  def dfs(root):
    if(root):
      temp = dfs(root.left)
      root.left = dfs(root.right)
      root.right = temp
    return root
  return dfs(root)
root = [4,2,7,1,3,6,9]
expected = [4,7,2,9,6,3,1]