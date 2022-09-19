def longestCommonSubsequence(text1: str, text2: str) -> int:
  res = [[0 for j in range(-1,len(text2))] for i in range(-1,len(text1))]
  for i in range(len(text1)):
    for j in range(len(text2)):
      if text1[i] == text2[j]: # go diagonal
        # print(i,j,text1[i],res[i-1][j-1] + 1)
        res[i][j] = res[i-1][j-1] + 1
      else:
        res[i][j] = max(res[i-1][j], res[i][j-1])

  return res[len(text1)-1][len(text2)-1]
t1 = "bsbininm"
t2 = "jmjkbkjkv"
# t1 ="oxcpqrsvwf"
# t2 ="shmtulqrypy"
print(longestCommonSubsequence(t1,t2))