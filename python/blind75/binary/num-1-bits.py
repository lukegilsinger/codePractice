class Solution:
    def hammingWeight(self, n) -> int:
      r = 0
      while n:
        r += n % 2
        n = n >> 1
      return r
    def hammingWeight2(self, n) -> int:
      r = 0
      while n:
        n = n & (n - 1)
        r += 1
      return r
        

n = 7
sol = Solution()

ans1 = sol.hammingWeight(n)
print(ans1)

ans1 = sol.hammingWeight2(n)
print(ans1)