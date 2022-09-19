def isAnagram(s: str, t: str) -> bool:
  if len(s) != len(t): return False
  hm = {}
  for ss in s:
    if ss in hm: hm[ss] += 1
    else: hm[ss] = 1

  for tt in t:
    if tt in hm and hm[tt] > 0: hm[tt] -= 1
    if tt in hm and hm[tt] == 0: hm.pop(tt)

  if len(hm) == 0: return True
  return False

  
def isAnagram2(s: str, t: str) -> bool:
  if len(s) != len(t): return False
  return all(t.count(ss) == s.count(ss) for ss in s)

s = "anagram"
t = "nagaram"
# s = "rat"
# t = "tar"
print(isAnagram(s,t))