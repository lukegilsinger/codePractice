def lengthOfLIS(nums) -> int:
  LIS = [1] * len(nums)
  for i in range(1, len(nums)):
    j = i - 1
    while j >= 0:
      if nums[i] > nums[j]:
        LIS[i] = max(LIS[i], LIS[j] + 1)
      j -= 1
  print(LIS)
  return max(LIS)
nums = [1,2,4,3]
nums = [0,1,0,3,2,3]
# nums = [7,7,7,7,7,7,7]
# nums = [10,9,2,5,3,7,101,18]
print(lengthOfLIS(nums))