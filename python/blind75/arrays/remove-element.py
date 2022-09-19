# 27
class Solution:
    def removeElement(self, nums, val) -> int:
      i = 0
      for x in nums:
          if x != val:
              nums[i] = x
              i += 1
      return i


nums = [3,2,2,3]
val = 3

sol = Solution()

ans1 = sol.removeElement(nums, val)
print(ans1)