def isIsomorphic(s: str, t: str) -> bool:
  hms, hmt = {}, {}
  for i in range(len(s)):
    if s[i] not in hms:
      hms[s[i]] = t[i]
    else:
      if hms[s[i]] != t[i]: return False

    if t[i] not in hmt:
      hmt[t[i]] = s[i]
    else:
      if hmt[t[i]] != s[i]: return False
    
  return True
    
s = "egg"
t = "add"

s = "foo"
t = "bar"

s = "paper"
t = "title"
print(isIsomorphic(s,t))