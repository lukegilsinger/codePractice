
#convert to hashset

def longestConsecutive(nums):
  hashset = set(nums)
  longest = 0

  for n in hashset:
    if n - 1 in hashset: continue
    j = 1
    while n + j in hashset:
      j += 1
    longest = max(longest, j)
  return longest
nums = [100,4,200,1,3,2]
nums2 = [0,3,7,2,5,8,4,6,0,1]
print(longestConsecutive(nums))
print(longestConsecutive(nums2))