class Solution:
  def coinChange(self, coins, amount: int) -> int:
    hm = {0: 0}
    size = len(coins)
    for i in range(1, amount + 1):
      if i in coins: 
        hm[i] = 1
      else:
        smallest = amount + 1
        for c in coins:
          if i > c and hm[i-c] != -1 and hm[i-c] < smallest:
            smallest = hm[i-c]
        if smallest > amount: hm[i] = -1    
        else: hm[i] = smallest + 1
    return hm[amount]

sol = Solution()
c, a = [2], 3
ans1 = sol.coinChange(c, a)
print(ans1)