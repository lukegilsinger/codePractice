def minimumTotal(triangle):
  def dfs(level, pos):
    # print(level, pos)
    if level >= len(triangle) - 1 : 
      return triangle[level][pos]
    
    return min(dfs(level + 1, pos) + triangle[level][pos], dfs(level + 1, pos + 1) + triangle[level][pos])
  
  return dfs(0, 0)

def minimumTotal2(triangle):
  dp = [0] * (len(triangle) + 1)
  for row in triangle[::-1]:
    for i, n in enumerate(row):
      dp[i] = n +  min(dp[i], dp[i+1])
  return dp[0]
triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]

print(minimumTotal(triangle))
print(minimumTotal2(triangle))