from audioop import maxpp


def maxProfit(prices):
  l, r, maxi = 0, 0, 0
  while r < len(prices):
    diff = prices[r] - prices[l]
    print(diff)
    if diff >= 0:
      maxi = max(maxi, diff)
      r += 1
    if diff < 0: l += 1
  return maxi

prices = [7,1,5,3,6,4,10]
print(maxProfit(prices))