def isValidSudoku(board) -> bool:
  hm = {}
  for i in range(len(board)):
    hm[(i,9)] = set()
    for j in range(len(board[0])):
      # print(i,j, board[i][j])
      if board[i][j] != '.':
        if board[i][j] in hm[(i,9)]: return False
        hm[(i,9)].add(board[i][j])
        if (9,j) not in hm: hm[(9,j)] = set()
        if board[i][j] in hm[(9,j)]: return False
        hm[(9,j)].add(board[i][j])
        # hm3
        x = i // 3
        y = j // 3
        if (x, y) not in hm: hm[(x,y)] = set()
        if board[i][j] in hm[(x, y)]: 
          return False
        hm[(x,y)].add(board[i][j])

  return True

board = [["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
print(isValidSudoku(board))