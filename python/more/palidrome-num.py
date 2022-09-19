def isPalindrome(x: int) -> bool:
  s = str(x)
  for i in range((len(s)+1)//2):
    if s[i] != s[-i-1]:
      return False
  return True

x = 12211
print(isPalindrome(x))