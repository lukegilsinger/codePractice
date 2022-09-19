class Solution:
    def maxSubArray(self, nums) -> int:
      curr = 0
      largest = nums[0]
      l, r = 0, 0
      while r < len(nums):
        curr = curr + nums[r]
        # print(curr)
        largest = max(largest, curr)
        r += 1
        if curr < 0:
          curr = 0
      return largest
    
    def maxSubArray2(self, nums) -> int:
      curr = 0
      largest = nums[0]
      for n in nums:
        if curr < 0:
          curr = 0
        curr = curr + n
        # print(curr)
        largest = max(largest, curr)
      return largest
          

s = Solution()
l = [-2,1,-3,4,-1,2,1,-5,4]
print (s.maxSubArray(l))
print (s.maxSubArray2(l))