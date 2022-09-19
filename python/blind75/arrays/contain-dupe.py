from operator import truediv


class Solution:
    def containsDuplicate(self, nums) -> bool:
        hs = set()
        for n in nums:
          if n in hs:
            return True
          hs.add(n)
        return False
    def containsDuplicate2(self, nums) -> bool:
      return len(nums) > len(set(nums))

s = Solution()
l = [1,6,8,9,3,4,11,1,14,5]
print(s.containsDuplicate(l))
print(s.containsDuplicate2(l))