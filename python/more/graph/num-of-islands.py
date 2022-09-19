def numIslands(grid):
  rows = len(grid)
  cols = len(grid[0])
  cnt = 0
  visited = set()

  def dfs(r,c,depth):
    print(r,c,depth)
    if (r < 0 or c < 0 or r >= rows or c >= cols or (r,c) in visited): return
    visited.add((r,c))
    if grid[r][c] == "1":
      for (x,y) in [(1,0),(-1,0),(0,1),(0,-1)]: dfs(r+x, c+y, depth+1)
      if depth == 0: 
        return True

  for i in range(rows):
    for j in range(cols): 
      if dfs(i,j,0) == True: cnt += 1
  return cnt

grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
print(numIslands(grid))