def longestCommonPrefix(strs) -> str:
  if not strs: return ""
  short = min(strs, key=len)
  for i, sh in enumerate(short):
    for s in strs:
      if s[i] != sh: return short[:i]

  return short


s = ["fl","fly","fl"]
s = ["", "b"]
s = ["ab", "a"]
print(longestCommonPrefix(s))