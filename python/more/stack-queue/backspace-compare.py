def backspaceCompare(s: str, t: str) -> bool:
  stackS, stackT = [], []
  p = 0
  while p < len(s):
    if s[p] == '#' :
      if len(stackS) > 0: stackS.pop()
    else: stackS.append(s[p])
    p += 1
  p = 0
  while p < len(t):
    if t[p] == '#': 
      if len(stackT) > 0: stackT.pop()
    else: stackT.append(t[p])
    p += 1
  print(stackS, stackT)
  return stackS == stackT 
  

s = "ab#c"
t = "ad#c"
s = "a##c"
t = "#a#c"
print(backspaceCompare(s, t))