def minCostClimbingStairs(cost) -> int:
  print(cost)
  hm = {}
  def rec(step):
    print(step)
    if step in hm: return hm[step]
    if step in [0,1]: return 0
    hm[step] = min(rec(step-1) + cost[step-1], rec(step-2) + cost[step-2])
    return hm[step]
  return rec(len(cost))
cost = [10,15,20]
cost = [1,100,1,1,1,100,1,1,100,1]
print((minCostClimbingStairs(cost)))