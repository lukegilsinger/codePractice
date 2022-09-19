#PRIORITY QUEUE (HEAP)

from collections import deque
import heapq

#Use max heap. python need to uses min heap with negatives
# O(N) to create heap
def lastStoneWeight(stones) -> int:
  stones = [-s for s in stones]
  heapq.heapify(stones)

  while len(stones) > 1:
    print(stones)
    s1 = heapq.heappop(stones)
    s2 = heapq.heappop(stones)
    print(s1 , s2)
    heapq.heappush(stones, s1- s2)
  stones.append(0)
  return -1 * stones[0]


stones = [2,7,4,1,8,1]
print(lastStoneWeight(stones))