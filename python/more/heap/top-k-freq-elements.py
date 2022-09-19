# bucket sort

def topKFrequent(nums, k):
  hm = {}
  frq = [[] for i in range(len(nums) + 1)]
  print(frq)
  for n in nums:
    if n in hm: hm[n] = hm[n] + 1
    else: hm[n] = 1
  for n, c in hm.items():
    frq[c].append(n)
  print(hm)
  print(frq)

  res = []
  for i in range(len(frq) -1, 0, -1):
    for n in frq[i]:
      res.append(n)
      if len(res) >= k: return res

nums = [1,1,1,2,2,3,3]
k = 2

print(topKFrequent(nums, k))