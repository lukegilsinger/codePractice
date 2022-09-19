#33
import time

class Solution:
    def search(self, nums, target) -> int:
      l = 0
      r = len(nums) - 1

      while l <= r:
        m = (l + r) // 2
        print(nums[m])
        if target == nums[m]: return m
        # in left sorted
        if nums[l] <= nums[m]:
          if (target < nums[l] or target > nums[m]): l = m + 1 # go right
          else: r = m - 1 # go left
        # in right sorted
        else: 
          if (target > nums[r] or target < nums[m]): r = m - 1 # go left
          else: l = m + 1 # go right
      return -1 # not found

#bin search
# has left, mid, right pointer
# left <= right

nums = [4,5,6,7,0,1,2]
target = 3

sol = Solution()

ans1 = sol.search(nums, target)
print(ans1)