class Solution:
  def maxProfit(self, prices) -> int:
    largest = 0
    for i, p in enumerate(prices):
      for p2 in prices[i:]:
        d = p2 - p
        # print(d)
        if (d > largest):
          largest = d
    return largest
  def maxProfit2Pointer(self, prices) -> int:
    l, r = 0, 1
    maxi = 0
    while r < len(prices):
      if prices[l] < prices[r]:
        curr = prices[r] - prices[l]
        # print(l, r)
        maxi = max(maxi, curr)    
      else:
        l = r
      r += 1
    return maxi


sol = Solution()
prices = [7,3,5,1,6,4]
# res = sol.maxProfit(prices)
# print(res)
res2 = sol.maxProfit2Pointer(prices)
print(res2)