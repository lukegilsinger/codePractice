def generate(numRows: int):
  res = [[1]]
  for n in range(1, numRows):
    t = []
    for i in range(n+1):
      if i == 0: t.append(res[n-1][i])
      elif i == n: t.append(res[n-1][i-1])
      else: t.append(res[n-1][i] + res[n-1][i-1])
    res.append(t)
  return res


print(generate(3))