
def canFinish(numCourses: int, prerequisites) -> bool:
  #preMap
  # hm = {}
  # for n in range(numCourses): hm[n] = []
  hm = {i:[] for i in range(numCourses)}

  for p in prerequisites:
    hm[p[0]].append(p[1])

  def dfs(node, visited):
    print(node)
    print(visited)
    if len(hm[node]) == 0:
      return True
    if node in visited: return False
    visited.append(node)
    for p in hm[node]: 
      if not dfs(p, visited): return False
    hm[node] = [] # because know if can be taken 
    return True
  
  for node in range(numCourses):
    if not dfs(node, []): return False
  return True

numCourses = 20
prerequisites = [[0,10],[3,18],[5,5],[6,11],[11,14],[13,1],[15,1],[17,4]]
print(canFinish(numCourses, prerequisites))