def search(nums, target):
  low, high = 0, len(nums)-1

  while (high - low >= 0):
    # p = (high+low)/2
    p = (low + (high - low)//2)
    print(p)
    if nums[p] == target: return p
    if nums[p] > target: 
      high = p-1
    if nums[p] < target:
      low = p+1
  return -1

nums = [-1,0,3,5,9,12]
target = 9
print(search(nums, target))