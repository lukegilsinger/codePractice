def findBall(grid):
  cols = len(grid[0])
  rows = len(grid)
  # print(cols, rows)

  def dropBall(col,row):
    # print(col,row)
    if col < 0 or col == cols: return -1
    if row >= rows : return col
    if grid[row][col] == 1:
      #check right
      if col+1 > cols - 1 or grid[row][col+1] == -1: return -1
      else: #go down right
        return dropBall(col+1, row+1)
    if grid[row][col] == -1:
      #check left
      if col-1 < 0 or grid[row][col-1] == 1: return -1
      else: #go down left
        return dropBall(col-1, row+1)
  # return dropBall(0, 0)
  return [(lambda c: dropBall(c,0))(c) for c in range(cols)]


grid = [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]
grid = [[-1]]
grid = [[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]]
print(findBall(grid))