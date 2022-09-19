class Solution:
    def climbStairs(self, n: int) -> int:
        hm = {
          1: 1,
          2: 2,
        }
        for i in range(3, n + 1):
          hm[i] = hm[i-1] + hm[i-2]
        return hm[n]
    def climbStairsLessMem(self, n: int) -> int:
      hm = {
        2: 1, # minus 2
        1: 2, # minus 1
      }
      for i in range(3, n + 1):
        temp = hm[1]
        hm[1] = hm[2] + temp
        hm[2] = temp
      return hm[1]

sol = Solution()
n = 36
ans1 = sol.climbStairs(n)
print(ans1)
ans2 = sol.climbStairsLessMem(n)
print(ans2)