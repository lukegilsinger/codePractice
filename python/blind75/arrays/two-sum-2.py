# sorted.  solve with O(1) space
class Solution:
    def twoSum(self, numbers, target):
      #two pointers
      l, r = 0, len(numbers) - 1
      while l < r:
        sum = numbers[l] + numbers[r]
        if sum == target:
          return [l + 1, r + 1]
        elif sum < target:
          l += 1
        else:
          r -= 1

    def twoSumDic(self, numbers, target):
      hm = {}
      
      for i, n in enumerate(nums):
        diff = target - n
        if diff in hm:
          return [hm[diff]+1, i+1]
        hm[n] = i

nums = [2,7,11,15]
target = 9

sol = Solution()

ans1 = sol.twoSum(nums, target)
print(ans1)