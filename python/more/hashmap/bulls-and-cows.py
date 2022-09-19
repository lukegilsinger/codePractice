def getHint(secret: str, guess: str) -> str:
  b, c = 0, 0
  hms, hmg = {},{}
  for i in range(len(guess)):
    if secret[i] == guess[i]: b += 1
    else:
      if secret[i] in hms: hms[secret[i]] += 1
      else: hms[secret[i]] = 1
      if guess[i] in hmg: hmg[guess[i]] += 1
      else: hmg[guess[i]] = 1
  for x in hmg:
    if x in hms:
      c += min(hmg[x], hms[x])
  return f'{b}A{c}B'

s = "1807"
g = "7810"

s = "1123"
g = "0111"
print(getHint(s,g))