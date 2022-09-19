def rob(nums):
  r1, r2 = 0, 0
  for n in nums:
    new = max(r2 + n, r1)
    r2 = r1
    r1 = new

  return r1


n = [1,2,3,1]
n = [2,7,9,3,1]
print(rob(n))