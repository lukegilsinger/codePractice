

def floodFill(image, sr: int, sc: int, color: int):
  def dfs(im, r, c, col, orig):
    if col == orig: return im
    if (r < 0 or c < 0 or r >= len(im) or c >= len(im[r])): 
      return
    if im[r][c] != orig: return
    if im[r][c] == orig: im[r][c] = col

    # if r < len(im) - 1 and im[r+1][c] == val: dfs(im, r+1, c, col, val)
    # if c < len(im[r]) - 1 and im[r][c+1] == val: dfs(im, r, c+1, col, val)
    # if r > 0 and im[r-1][c] == val: dfs(im, r-1, c, col, val)
    # if c > 0 and im[r][c-1] == val: dfs(im, r, c-1, col, val)

    for (x, y) in [(1,0),(-1,0),(0,1),(0,-1)]:
      dfs(im,r+x,c+y,col,orig)

    return im
  return dfs(image, sr, sc, color, image[sr][sc])

image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1
sc = 1
color = 2
image = [[0,0,0],[0,0,0]]
sr = 0
sc = 0
color = 0
print(floodFill(image, sr, sc, color))