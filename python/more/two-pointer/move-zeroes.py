def moveZeroes(nums) -> None:
  """
  Do not return anything, modify nums in-place instead.
  """
  i, j = 0, 1
  while i < len(nums) and j < len(nums):
    # print(i,j)
    # print(nums)
    if nums[i] == 0 and nums[j] != 0:
      nums[i] = nums[j]
      nums[j] = 0
      i += 1
      j += 1
    elif nums[i] == 0 and nums[j] == 0:
      j += 1
    elif nums[i] != 0:
      i += 1
      j += 1
  # print(nums)
   
nums = [0,1,0,0,12]
# nums = [0]
moveZeroes(nums)