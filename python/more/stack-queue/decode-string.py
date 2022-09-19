def decodeString(s: str) -> str:
  def decRec(st):
    return res

  return decRec(s)

def decodeStringStk(s: str) -> str:
  stk = []
  for i in range(len(s)):
    if s[i] == ']':
      temp = ''
      while stk[-1] != '[':
        temp = stk.pop() + temp
      stk.pop()
      num = ''
      while stk and stk[-1].isdigit(): # for multi digit
        num = stk.pop() + num
      stk.append(int(num) * temp)
    else: stk.append(s[i])  
  return "".join(stk)
s = "3[a]2[bc]"
s = "3[a2[c]]"
print(decodeStringStk(s))