import collections

def firstUniqChar(s: str) -> int:
  hm = {}
  q = collections.deque()
  for i, c in enumerate(s):
    if c not in hm:
      hm[c] = i
      q.append(c)
    elif c in q:
      q.remove(c)
  if len(q) == 0: return -1
  return hm[q.popleft()]

s = "leetcodeldto"
print(firstUniqChar(s))