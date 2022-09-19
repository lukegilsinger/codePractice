# key is to go through lists backwards

def merge(nums1, m, nums2, n):
  print(nums1, nums2)
  i = m + n - 1
  m -= 1
  n -= 1
  while n >= 0:
    if m >= 0 and nums1[m] > nums2[n]:
      nums1[i] = nums1[m]
      m -= 1
    else:
      nums1[i] = nums2[n]
      n -= 1
    i -= 1
    
  print(nums1)


nums1 = [1,2,3,0,0,0]
m = 3
nums2 = [0,5,6]
n = 3
print(merge(nums1, n, nums2, m))