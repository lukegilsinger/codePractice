def runningSum(nums):
  sums = [nums[0]]
  for n in nums[1:]:
    sums.append(n+sums[-1])
  return sums
n = [1,2,3,4]
print(runningSum(n))