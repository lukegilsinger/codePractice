class Solution:
    # returns list of list of ints
    def threeSum(self, nums):
      res = []
      nums.sort()
      for i, n in enumerate(nums):
        if i > 0 and n == nums[i-1]:
          continue
        l, r = i + 1, len(nums) - 1
        while l < r:
          sum = n + nums[l] + nums[r]
          # print("--------")
          # print(n, sum)
          if sum == 0:
            # print([n, nums[l], nums[r]])
            res.append([n, nums[l], nums[r]])
            l += 1
            while nums[l] == nums[l-1] and l < r: l += 1
          elif sum < 0: 
            l += 1
          else: 
            r -= 1
      return res



nums = [-1,0,1,2,-1,-4]

sol = Solution()

ans1 = sol.threeSum(nums)
print(ans1)