def characterReplacement(s: str, k: int) -> int:

  l, r = 0, 1
  hm = {  }
  longest = 0

# use a maxF to increase efficitancy
  maxF = 0
  
  while r <= len(s):

    if s[r-1] in hm: hm[s[r-1]] += 1
    else: hm[s[r-1]] = 1
    wSize = r - l
    updated = False
    for i in hm.keys():
      if i in hm and hm[i] >= wSize - k:
        
        longest = max(longest, min(wSize, hm[i] + k))
        updated = True
    if not updated:
      hm[s[l]] -= 1
      l += 1
    r += 1
    # print(l,r,longest)
    # print(hm)

  return longest

s = "ABAA"
k = 0
# s = "AABABBA"
# k = 1
print(characterReplacement(s, k))