def preorder(self, root: 'Node') -> List[int]:
  res = []
  def dfs(root):
    if root:
      res.append(root.val)
      for r in root.children:
        dfs(r)
  
  return res


