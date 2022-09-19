def getRow(rowIndex: int):
  res = [[1]]
  for r in range(1, rowIndex+1):
    t = []
    for i in range(r+1):
      if i == 0: t.append(res[r-1][i])
      elif i == r: t.append(res[r-1][i-1])
      else: t.append(res[r-1][i] + res[r-1][i-1])
    res.append(t)
  return res[-1]

print(getRow(33))