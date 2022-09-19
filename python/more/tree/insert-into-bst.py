
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
  if not root: return TreeNode(val)
  def dfs(root):
    if(root):
      print(root.val)
      if root.val > val:
        if root.left: dfs(root.left)
        else: root.left = TreeNode(val)
      elif root.val < val:
        if root.right: dfs(root.right)
        else: root.right = TreeNode(val)
  dfs(root)
  return root