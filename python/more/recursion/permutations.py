def permute(nums):
  # def dfs(nums, path, res):
  #   if not nums:
  #     res.append(path)
  #     # return # backtracking
  #   for i in xrange(len(nums)):
  #     dfs(nums[:i]+nums[i+1:], path+[nums[i]], res)

  res = []
  return dfs(nums, [], res)

  def dfs2(numms):
    res = []
    print(numms)
    if len(numms) == 1:
      return [numms[:]]
    else:
      for n in range(len(numms)):
        n = numms.pop(0)
        perms = dfs2(numms)
        for p in perms:
          p.append(n)
        res.extend(perms)
        numms.append(n)
      return res
      

nums = [1,2,3]
print(permute(nums))