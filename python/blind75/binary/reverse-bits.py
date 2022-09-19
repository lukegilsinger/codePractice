class Solution:
    def reverseBits(self, n: int) -> int:
      print(n)
      res = 0
      for i in range(32):
        bit = (n >> i) & 1
        res = res | (bit << (31 - i))
      return res
sol = Solution()
n = 43261596 #'00000010100101000001111010011100'
ans1 = sol.reverseBits(n)
print(ans1)