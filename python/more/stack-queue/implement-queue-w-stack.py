class MyQueue:

  def __init__(self):
    self.stk = []
    self.q = []

  def push(self, x: int) -> None:
    self.stk = []
    for i in self.q:
      self.stk.append(i)
    self.q = []
    self.q.append(x)
    for i in self.stk:
      self.q.append(i)
    # print(self.q)

  def pop(self) -> int:
    return self.q.pop()

  def peek(self) -> int:
    return self.q[-1]

  def empty(self) -> bool:
    return len(self.q) == 0
      
q = MyQueue()
print(q.push(1))
print(q.push(2))
print(q.peek())
print(q.pop())
print(q.empty())

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()