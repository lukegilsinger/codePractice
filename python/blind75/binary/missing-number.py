from tkinter import N


class Solution:
  # bit operation solution
    def missingNumber(self, nums) -> int:
      res = 0
      for i in range(0,len(nums)):
        print(i)
        res = i ^ res ^ nums[i]
      return res ^ len(nums)
      # subtract sums solution
    def missingNumber2(self, nums) -> int:
      res = len(nums)
      for i in range(0,len(nums)):
        res += (i - nums[i])
      return res

# nums = [9,6,4,2,3,5,7,0,1]
nums = [3,0,1]

sol = Solution()

ans1 = sol.missingNumber2(nums)
print(ans1)