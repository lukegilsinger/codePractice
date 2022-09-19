def canConstruct(ransomNote: str, magazine: str) -> bool:
  hm = {}
  for s in magazine:
    if s in hm: hm[s] += 1
    else: hm[s] = 1
  for r in ransomNote:
    if r in hm and hm[r] > 0: hm[r] -= 1
    else: return False
  return True

  
def canConstruct2(ransomNote: str, magazine: str) -> bool:
  return all(ransomNote.count(c)<=magazine.count(c) for c in set(ransomNote))

ransomNote = "a"
magazine = "b"
print(canConstruct(ransomNote, magazine))