def numIslands(grid):
  rows = len(grid)
  cols = len(grid[0])
  cnt = 0
  visited = set()

  def dfs(r,c, depth):
    print(r,c,depth)
    if r >= rows or c >= cols or r < 0 or c < 0 or (r,c) in visited:
      return
    visited.add((r,c))
    if grid[r][c] == '0': return
    if grid[r][c] == '1':
      dfs(r+1, c, depth+1)
      dfs(r, c+1, depth+1)
      dfs(r-1, c, depth+1)
      dfs(r, c-1, depth+1)
      if depth == 0:
        return True
    return

  for r in range(rows):
    for c in range(cols):
      if dfs(r,c, 0) == True: cnt += 1
  return cnt

from collections import deque

def numIslands2(grid):
  rows = len(grid)
  cols = len(grid[0])
  cnt = 0
  visited = set()

  # bfs is iteritave
  def bfs(r,c):
    q = deque() #just a queue
    visited.add((r,c))
    q.append((r,c))
    cnt = 0
    while q:
      # pop left pops first element we added
      row, col = q.popleft() #can change to popright for iteritive dfs
      directions = [[1,0],[-1,0],[0,1],[0,-1]]
      for dr, dc in directions:
        r = row+dr
        c = col+dc
        if ((r) in range(rows) and (c) in range(cols)
        and grid[r][c] == "1" and (r,c) not in visited):
          q.append((r,c))
          visited.add((r,c))


  for r in range(rows):
    for c in range(cols):
      if grid[r][c] == "1" and (r,c) not in visited:
        bfs(r,c)
        cnt += 1
  return cnt


grid1 = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
grid = [["1","1","1"],
["0","1","0"],
["1","1","1"]]
print(numIslands(grid1))
print(numIslands2(grid1))