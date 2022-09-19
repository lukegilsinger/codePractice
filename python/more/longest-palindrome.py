def longestPalindrome(s: str) -> int:
  counts = {}
  length = 0
  anOdd = False
  for i in s:
    if i in counts:
      counts[i] += 1
    else:
      counts[i] = 1
  for c in counts:
    # print(counts[c] // 2)
    if counts[c] % 2 == 1: 
      length += counts[c] - 1
      anOdd = True
    else: length += counts[c]
  return length + 1 if anOdd else length

# s = "abccccdd"
s = "ccc"
print(longestPalindrome(s))