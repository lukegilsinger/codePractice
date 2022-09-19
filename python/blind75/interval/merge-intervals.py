from typing import List

def merge(intervals: List[List[int]]) -> List[List[int]]:
  print(intervals)
  intervals.sort(key = lambda x: x[0])
  res = [intervals[0]]
  for s, e in intervals[1:]:
    lastEnd = res[-1][1]
    if s <= lastEnd: res[-1][1] = max(e, lastEnd)
    else: res.append([s,e])
  return res


intervals = [[1,3],[2,6],[8,10],[15,18]]
intervals = [[1,4],[4,5]]
print(merge(intervals))