from curses.ascii import isalpha


def letterCasePermutation(s: str):
  res = []
  def rec(s, perm):
    if len(s) == 0: res.append(perm)
    else:
      i = s[0:1]
      if i.isalpha():
        rec(s[1:],perm[:] + i.lower())
        rec(s[1:],perm[:] + i.upper())
      else:
        rec(s[1:],perm[:] + i)
  rec(s, "")
  return res

s = "a1b2"
print(letterCasePermutation(s))