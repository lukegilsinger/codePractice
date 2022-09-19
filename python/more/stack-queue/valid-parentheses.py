def isValid(s: str) -> bool:
  d = {
    ")": "(",
    "}": "{",
    "]": "[",
  }
  stk = []
  for ss in s:
    if ss not in d:
      stk.append(ss)
    else:
      if len(stk) > 0 and d[ss] == stk[-1]: stk.pop()
      else: return False
  if len(stk) == 0: return True
  return False

s = "()[]{}{"
print(isValid(s))