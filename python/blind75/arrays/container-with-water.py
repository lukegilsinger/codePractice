class Solution:
    def maxArea(self, height) -> int:
      l, r = 0, len(height) - 1
      mx = 0
      while l < r:
        mx = max((r - l) * min(height[l], height[r]), mx)
        if height[l] < height[r]:
          l += 1
        else: 
          r -= 1
      return mx
    def maxArea2(self, height) -> int:
        l, r = 0, len(height) - 1
        mx = 0
        while l < r:
          mx = max((r - l) * min(height[l], height[r]), mx)
          if height[l] < height[r]:
            while height[l] > height[l+1]:
              l += 1
            l += 1
          else: 
            while height[r] > height[r-1]:
              r -= 1
            r -= 1
        return mx

height = [1,8,6,2,5,4,8,3,7]
sol = Solution()

ans1 = sol.maxArea(height)
print(ans1)