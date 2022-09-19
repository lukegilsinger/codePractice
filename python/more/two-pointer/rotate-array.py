def rotate(nums, k):
  # rotate once three times
  for _ in range(k):
    prev = nums[-1] #initiate a first previous 
    for i in range(len(nums)):
        temp = nums[i] #hodl nums[i]
        nums[i] = prev #overwrite the current index 
        prev = temp #swap the value 

def rotate2(nums, k):
  #use extra space and mod
  new = [0] * len(nums)
  for i in range(len(nums)):
      new[(i+k)%len(nums)] = nums[i]

  for i in range(len(nums)):
      nums[i] = new[i]
  print(nums)

def rotate3(nums, k):
  #cyclic replacement
  c = 0
  p = 0
  prev = nums[p]
  curr = nums[0]
  while c < len(nums):
    nextp = (k+p) % len(nums)
    temp = nums[nextp]
    prev = nums[p % len(nums)]
    nums[(k+p) % len(nums)] = temp
    curr = 
    p += k
    c += 1
  print(nums)

nums = [1,2,3,4,5,6,7]
k = 3
# nums = [-1,-100,3,99]
# k = 2
# print(rotate(nums, k))
print(rotate2(nums, k))
print(rotate3(nums, k))