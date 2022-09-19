
def construct(badV):
  def isBadVersion(version: int):
    if version == badV: return True
    return False
  return isBadVersion

# hash map was faster
def firstBadVersionHM(n: int, isBadVersion) -> int:
  hm = {}
  l, r = 0, n
  while r >= l:
    pivot = l + (r - l)  // 2
    bad = isBadVersion(pivot)
    print(pivot, bad)
    if pivot-1 in hm and hm[pivot-1] == False and bad == True: return pivot
    if pivot+1 in hm and hm[pivot+1] == True and bad == False: return pivot + 1
    if(bad):
      r = pivot - 1
    else: 
      l = pivot + 1
    hm[pivot] = bad

  return False

def firstBadVersion(n: int, isBadVersion) -> int:
  l, r = 0, n
  while r >= l:
    pivot = l + (r - l)  // 2
    if(isBadVersion(pivot)):
      r = pivot - 1
    else: 
      l = pivot + 1

  return l
n = 5
bad = 4
isBadVersion = construct(bad)
print(firstBadVersion(n, isBadVersion))