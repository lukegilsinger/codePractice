class Solution:
  def twoSum(self, nums, target):
    print('in twoSum')
    print(nums)
    print(target)
    hm = {}
    
    for i, n in enumerate(nums):
      diff = target - n
      if diff in hm:
        return [hm[diff], i]
      hm[n] = i
    
  def twoSumBad(self, nums, target):
    for i, n in enumerate(nums):
      for j, m in enumerate(nums):
        if n + m == target:
          print(n, m)
          return [i, j]
    return False


nums = [7,8,2,4]
target = 6

sol = Solution()

ans1 = sol.twoSum(nums, target)
print(ans1)
ans = sol.twoSumBad(nums, target)
print(ans)


