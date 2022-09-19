def pivotIndex(nums):
  print(nums)
  l = 0
  r = sum(nums)
  s = len(nums)
  for p in range(s):
    print(l,r)
    l = l + nums[p]
    if l == r: return p
    r = r - nums[p]
  return -1

nums = [1,7,3,6,5,6]
print(pivotIndex(nums))