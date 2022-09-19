#faster if not rec
def uniquePaths(m: int, n: int) -> int:
  hm={}
  def rec(x,y):
    if x == 0 or y == 0: return 1
    if (x,y) in hm: return hm[(x,y)]
    hm[(x,y)] = rec(x-1, y) + rec(x, y-1)
    return hm[(x,y)]
  # res = rec(m-1,n-1)
  # print(hm)
  return rec(m-1,n-1)
m = 3
n = 7
print(uniquePaths(m, n))