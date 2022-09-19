def spiralOrder(matrix):
  m = len(matrix[0])
  n = len(matrix)
  i, j = 0, 0
  i_min, j_min, i_max, j_max, = 0,0,m-1,n-1
  res = []
  move = 'r' if m > 1 else 'd'
  while len(res) < m * n:
    print(j,i,move)
    print(matrix[j][i])
    res.append(matrix[j][i])
    if move == 'r':
      i += 1
      if i == i_max:
        move = 'd'
        j_min += 1
        continue
    if move == 'd':
      j += 1
      if j == j_max:
        move = 'l'
        i_max -= 1
        continue
    if move == 'l':
      i -= 1
      if i == i_min:
        move = 'u'
        j_max -= 1
        continue
    if move == 'u':
      j -= 1
      if j == j_min:
        move = 'r'
        i_min += 1
        continue

  return res


matrix = [[1,2,3],[4,5,6],[7,8,9]]
matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
matrix = [[1]]
print(spiralOrder(matrix))