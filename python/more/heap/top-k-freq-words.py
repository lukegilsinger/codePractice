# bucket sort
import heapq

def topKFrequent(words, k):
  hm = {}
  frq = [[] for i in range(len(words) + 1)]

  for w in words:
    if w in hm: hm[w] = hm[w] + 1
    else: hm[w] = 1
  for w, c in hm.items():
    frq[c].append(w)
  print(hm)
  print(frq)
  res = []

  for i in range(len(frq) -1, 0, -1):
    frq[i].sort()
    for w in frq[i]:
      res.append(w)
      if len(res) >= k: return res

words = ["i","love","leetcode","i","love","coding"]
k = 3

# words = ["aaa","aa","a"]
# k = 2

print(topKFrequent(words, k))