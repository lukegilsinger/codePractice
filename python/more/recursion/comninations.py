def combine(n: int, k: int):
  res = []
  def rec(st, com):
    if len(com) == k:
      res.append(com)
    else:
      for i in range(st+1, n+1):
        cop = com.copy()
        cop.append(i)
        print(cop)
        rec(i,cop)

  # for i in range(1,n+1):
  #   for j in range(i+1,n+1):
  #     for k in range(j+1,n+1):
  #       print(i,j,k)
  rec(0, [])
  return res
  

n = 4
k = 2
print(combine(n,k))