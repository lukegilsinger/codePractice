def reverseString(s) -> None:
  """
  Do not return anything, modify s in-place instead.
  """
  i, j = 0, len(s) - 1
  while i < j:
    s[i], s[j] = s[j], s[i]
    i += 1
    j -= 1
  print(s) 
def reverseString2(s) -> None:
  s[:] = s[::-1]




s = []
reverseString(s)