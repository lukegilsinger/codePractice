
# TODO
# binary tree
def rob(nums):
  def robRec(nums):
    if len(nums) == 0: return
    print(nums)
    robRec(nums[1::2])
    robRec(nums[::2])

  robRec(nums)


  return r1


n = [3,2,3,None,3,None,1]
print(rob(n))