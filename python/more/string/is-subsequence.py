def isSubsequence(s: str, t: str) -> bool:
  #two pointer
  sp, tp = 0, 0
  if len(s) == 0: return True
  while sp < len(s) and tp < len(t):
    if s[sp] == t[tp]:
      if sp == len(s) - 1: # at end
        return True
      sp += 1
    tp += 1
  return False

s = ""
t = "ahbgdc"
print(isSubsequence(s,t))