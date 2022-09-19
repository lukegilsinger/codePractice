#already finished

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
      curr = 0
      largest = nums[0]
      for n in nums:
        if curr < 0:
          curr = 0
        curr = curr + n
        # print(curr)
        largest = max(largest, curr)
      return largest