def maxProduct(nums) -> int:
  print(nums)
  l = len(nums)
  if l == 1:
    return nums[0]
  return maxProduct(nums[:round(l/2)]) * maxProduct(nums[round(l/2):])

def maxProduct2(nums) -> int:
    res = max(nums)
    cmin, cmax = 1,1

    for n in nums:
        if n == 0:
            cmin, cmax = 1,1
            continue
        tmp = cmax*n
        cmax = max(tmp, n*cmin, n)
        cmin = min(tmp, n*cmin, n)
        res = max(res, cmax)
        print(cmin, cmax)

    return res

arr = [2,3,-2, 4]
half = round(4/2)
# print(arr[:half], arr[half:])
# print(maxProduct(arr))
print(maxProduct2(arr))