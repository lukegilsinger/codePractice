from typing import List
def eraseOverlapIntervals(intervals: List[List[int]]) -> int:
  intervals.sort(key = lambda x: x[0])
  last = intervals[0][1]
  cnt = 0
  for s,e in intervals[1:]:
    if s < last:
      cnt += 1
      last = min(e,last)
    else: 
      last = e
  return cnt

intervals = [[1,2],[2,3],[3,4],[1,3]]
print(eraseOverlapIntervals(intervals))