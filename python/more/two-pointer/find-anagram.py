# sliding window

from xml.dom.pulldom import PROCESSING_INSTRUCTION


def findAnagram(s, p):
  print(s)
  temp = list(p)
  l,r = 0,0
  res = []

  while l < len(s):
    print(l,r)
    if r >= len(s): 
      l += 1
      r = l
      temp = list(p)
    elif s[r] not in temp:
      l += 1
      r = l
      temp = list(p)
    elif s[r] not in p:
      l = r + 1
      r = l
      temp = list(p)
    elif s[r] in temp:
      temp.remove(s[r])
      r += 1
    if len(temp) == 0:
      res.append(l)
      l += 1
      r = l
      temp = list(p)

  return res
    
def findAnagram2(s, p):
  pCount = {}
  for i in p:
    if i in pCount: pCount[i] += 1
    else: pCount[i] = 1

  sCount = {}
  l, r = 0, len(p)
  for i in s[0:r]:
    if i in sCount: sCount[i] += 1
    else: sCount[i] = 1
    if i not in pCount: pCount[i] = 0

  res = []

  while r <= len(s):
    print(pCount, sCount)
    if pCount == sCount:
      res.append(l)
    if r >= len(s): break
    sCount[s[l]] -= 1
    l += 1
    if s[r] in sCount: sCount[s[r]] += 1
    else: sCount[s[r]] = 1
    if s[r] not in pCount: pCount[s[r]] = 0
    r += 1

  return res

s = "cbaebabacd"
p = "abc"
# s = "abab"
# p = "ab"
s = "baa"
p = "aa"
print(findAnagram2(s,p))