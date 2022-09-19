# cant use + or -
class Solution:
    def getSum(self, a, b) -> int:
      mask = 0xffffffff
      while b & mask != 0:
        d = a & b # to getting the carry
        a = a ^ b # xor
        b = d << 1 # shift.  if no carry then done

      return a&mask if b > mask else a
        

n1, n2 = 3, 9

sol = Solution()

ans1 = sol.getSum(n1, n2)
print(ans1)