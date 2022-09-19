def isSymmetric(self, root: Optional[TreeNode]) -> bool:
  l, r = [], []
  def dfsL(root):
    if root: 
      l.append(root.val)
      dfsL(root.left)
      dfsL(root.right)
    else: r.append(None)
  def dfsR(root):
    if root: 
      r.append(root.val)
      dfsR(root.right)
      dfsR(root.left)
    else: r.append(None)
  dfsL(root)
  dfsR(root)
  print(l,r)
  return l == r

def isSymmetric2(root):
  if not root: return True
  def dfs(l,r):
    print(l.val, r.val)
    if l and r:
      return l.val == r.val and dfs(l.left, r.right) and dfs(l.right, r.left)
    return l == r
  dfs(root.left, root.right)
root = [1,2,2,3,4,4,3]