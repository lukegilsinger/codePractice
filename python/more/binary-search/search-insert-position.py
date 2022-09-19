def searchInsert(nums, target: int) -> int:
  l,r = 0, len(nums)-1
  while r >= l:
    pivot = l + (r-l) // 2
    n = nums[pivot]
    if n == target: return pivot
    if n < target: l = pivot + 1
    else: r =  pivot - 1
  print(l,r,pivot)
  return l

nums = [1,3,5,6]
target = 7
print(searchInsert(nums, target))