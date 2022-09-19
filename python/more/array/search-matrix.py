from operator import truediv


def searchMatrix(matrix, target: int) -> bool:
  m, n = len(matrix), len(matrix[0])
  row = 0

  l, h = 0, m
  while l <= h and l < m:
    mid = (l + h) // 2
    print(l,h,mid)
    if target < matrix[mid][0]: h = mid - 1
    elif target > matrix[mid][n-1]: l = mid + 1
    else: 
      row = mid
      break
  print(row)
  l, h = 0, n
  while l <= h and l < n:
    mid = (l + h) // 2
    if target == matrix[row][mid]: return True
    elif target < matrix[row][mid]: h = mid - 1
    elif target > matrix[row][mid]: l = mid + 1
  return False

matrix = [[1,4,5,7],[10,11,16,20],[23,30,34,60]]
target = 3
matrix = [[1]]
target = 10
print(searchMatrix(matrix, target))