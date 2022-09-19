def canJump(nums) -> bool:
  #recursive with dp still not eff enough
  hs = set()
  def dfs(pos):
    # print('--', pos)
    if pos < len(nums):
      hs.add(pos)
      jumps = nums[pos]
      for j in range(jumps,0,-1):
        # print(pos, j, pos+j)
        if pos + j not in hs:
          dfs(pos+j)

  dfs(0)
  print(hs)
  return len(nums) - 1 in hs

def canJumpGreedy(nums):
  goal = len(nums) - 1

  for n in range(len(nums)-1, -1, -1):
    if nums[n] + n >= goal: 
      goal = n
  return goal == 0

nums = [2,3,1,5,5,5,5,5,4,3,2,1,4]
nums = [3,2,1,0,4]
print(canJumpGreedy(nums))