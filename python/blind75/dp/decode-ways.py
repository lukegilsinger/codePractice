def numDecodings(s: str) -> int:
  hm = {len(s): 1}
  for i in range(len(s) -1, -1, -1):
    if s[i] == '0': hm[i] = 0
    else: hm[i] = hm[i+1]
    if (i+1 < len(s) and (s[i] == '1' or (s[i] == '2' and s[i+1] in '0123456'))):
      hm[i] += hm[i+2]
  return hm[0]

def numDecodings2(s: str) -> int:
  hm = {len(s): 1}
  def dfs(i):
    if i in hm:
      return hm[i]
    if s[i] == '0': return 0
    res = dfs(i+1)
    if (i + 1 < len(s) and (s[i] == '1' or (s[i] == '2' and s[i+1] in '0123456'))):
      res += dfs(i+2)
    hm[i] = res
    return res
  return dfs(0)

s = "10"
# s = "27"
s = "12"
s = "226109"
s = "06"
s = "121212"
print(numDecodings(s))
print(numDecodings2(s))