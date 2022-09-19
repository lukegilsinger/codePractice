def intersection(nums1, nums2):
  nums2.sort()
  nums1.sort()
  p1, p2 = 0, 0
  res = []
  while p1 < len(nums1) and p2 < len(nums2):
    # print(res, p1, p2, nums1[p1], nums2[p2])
    if nums1[p1] == nums2[p2]:
      res.append(nums1[p1])
      p1 += 1
      p2 += 1
    elif nums1[p1] < nums2[p2]:
      p1 += 1
    else: 
      p2 += 1
  return res

nums1 = [4,9,5]
nums2 = [9,4,9,8,4]
print(intersection(nums1, nums2))