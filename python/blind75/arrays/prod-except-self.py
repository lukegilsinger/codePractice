class Solution:
    def productExceptSelf(self, nums):
      i = 0
      pre, post = 1, 1
      r = [1] * len(nums)
      for i in range(len(nums)):
        r[i] = pre
        pre *= nums[i]
        # print(pre)
      for i in range(len(nums)-1,-1,-1):
        r[i] *= post
        post *= nums[i]
        # print(post)
      return r

s = Solution()
l = [1,2,3,4]
print(s.productExceptSelf(l))