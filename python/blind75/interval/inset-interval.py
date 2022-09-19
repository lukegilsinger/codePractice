
def insert(intervals, newInterval):
  new = []
  if len(intervals) == 0: return [newInterval]
  for n, i in enumerate(intervals):
    if newInterval[0] > i[1]:
      new.append(i)
    elif newInterval[1] < i[0]:
      new.append(newInterval)
      return new + intervals[n:]
    else:
      newInterval = [min(newInterval[0], i[0]),max(newInterval[1], i[1])]
      
  new.append(newInterval)
  return new

intervals = [[1,3],[6,9]]
newInterval = [2,5]
# intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]
# newInterval = [4,8]
print(insert(intervals, newInterval))