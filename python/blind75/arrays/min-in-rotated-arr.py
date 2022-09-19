import math

def findMin(nums) -> int:
  print(nums)
  s = len(nums)
  if s == 1:
    return False
  p = round(s / 2) - 1
  if s == 2: 
    p = 1
  print(s, p)
  if nums[p] > nums[p+1]:
    return nums[p+1]
  return findMin(nums[:p]) or findMin(nums[p:])
  
def findMin2(nums):
  l, r = 0, len(nums) - 1
  p = round(r/2)
  res = nums[0]
  while (l <= r):
    print(l,r,p)

    #sorted
    if (nums[l] < nums[r]):
      res = min(res, nums[l])

    p = (l+r) // 2
    res = min(res, nums[p])
    #serch right
    if(nums[p] > nums[l]):
      l = p + 1
    
    # search left
    else:
      r = p - 1
  return res


arr = [3,4,5,1,2]
arr2 = [2,1]
print(findMin2(arr2))