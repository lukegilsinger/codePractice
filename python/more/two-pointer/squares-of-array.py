import collections
def sortedSquares(nums):
  if len(nums) == 0: return []
  l, r = 0, len(nums) - 1
  nl, nr = nums[l]*nums[l], nums[r]*nums[r]
  sorted = []
  while l < r:
    if nl > nr:
      sorted.append(nl)
      l += 1
      nl = nums[l]*nums[l]
    else:
      sorted.append(nr)
      r -= 1
      nr = nums[r]*nums[r]

  sorted.append(min(nr, nl))

  return sorted[::-1]
nums = [-7,-3,2,3,11]
nums = []
print(sortedSquares(nums))