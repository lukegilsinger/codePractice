def combinationSum(nums, target: int) -> int:
  cache = {}
  for n in nums:
    for t in range(1,target+1):
      if t not in cache : cache[t] = []
      if n == t: cache[t].append([n])
      if n < t:
        for d in cache[t-n]:
          cache[t].append(d + [n])

  return cache[target]

def combinationSumRecursiveBin(nums, traget):
  res = []
  def dfs(i, cur, total):
    if total == target:
      res.append(cur.copy())
      return
    if i >= len(nums) or total > target: return
    cur.append(nums[i])
    dfs(i, cur, total + nums[i])
    cur.pop()
    dfs(i+1, cur, total)
  
  dfs(0, [], 0)
  return res


n = [2,3,6,7]
target = 7
print(combinationSum(n, target))
print(combinationSumRecursiveBin(n, target))