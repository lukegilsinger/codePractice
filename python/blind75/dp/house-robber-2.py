def rob(nums):
  return max(nums[0], helpRob(nums[1:]), helpRob(nums[:-1]))
def helpRob(self, nums):
  print(nums)
  r1, r2 = 0, 0
  for n in nums:
    new = max(r2 + n, r1)
    r2 = r1
    r1 = new
  return r1

# houses in a circle
n = [1,2,3]
print(rob(n))