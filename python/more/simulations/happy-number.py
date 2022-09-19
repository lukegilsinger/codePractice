def isHappy(num: int) -> bool:
  cache = set()
  def rec(n):
    if n == 1: return True
    if n in cache: return False
    cache.add(n)
    temp = 0
    st = str(n)
    for s in st:
      temp += int(s)**2
    
    return rec(temp)
  return rec(num)

def isHappy2(n: int) -> bool:
  cache = set()
  while n not in cache:
    cache.add(n)
    n = sum([int(s)**2 for s in str(n)])
  print(cache)
  return n == 1
  


n = 19
print(isHappy2(n))