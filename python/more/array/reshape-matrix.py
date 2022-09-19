def matrixReshape(mat, r: int, c: int):
  res = [[]]
  if len(mat) * len(mat[0]) != r * c: return mat
  rr = 0
  for i in range(len(mat)):
    for j in range(len(mat[i])):
      if len(res[rr]) == c: 
        res.append([])
        rr += 1
      res[rr].append(mat[i][j])
  return res

mat = [[1,2],[3,4]]
r = 4
c = 1
print(matrixReshape(mat, r, c))

