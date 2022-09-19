def wordBreak(s: str, wordDict) -> bool:
  cache = [False] * (len(s) + 1)
  cache[0] = True
  for l in range(len(s)+1):
    for i in range(l):
      if s[i:l] in wordDict and cache[i] == True: 
        cache[l] = True
        break
  print(cache)
  return cache[-1]
s = "leetcode"
wordDict = ["leet","code"]
# s = "catsandog"
# wordDict = ["cats","dog","sand","and","cat"]
s = "applepenapple"
wordDict = ["apple","pen"]
print(wordBreak(s, wordDict))