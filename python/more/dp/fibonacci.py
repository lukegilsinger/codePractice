def fib(n: int) -> int:
  hm = {
    0: 0,
    1: 1,
  }

  def rec(n):
    if n in hm: return hm[n]
    
    hm[n] = rec(n-2) + rec(n-1)

    return hm[n]

  return rec(n)

n = 6
print(fib(n))