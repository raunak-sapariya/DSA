class stack:
    def __init__(self,size):
        self.stack = []
        self.size = size

    def isempty(self):
        return len(self.stack) == 0

    def isfull(self):
        return len(self.stack) == self.size

    def push(self, data):
        if self.isfull():
            print("over flow")
            return

        self.stack.append(data)

    def pop(self):
        if self.isempty():
            print("under flow")
            return

        return self.stack.pop()

    def top(self):
        if self.isempty():
            return -1

        return self.stack[-1]

    def display(self):
        if self.isempty():
            print("stack is empty")
            return

        for i in range(len(self.stack)-1, -1, -1):
            print(self.stack[i], end=" ")
        print()

# s = stack(5)

# s.push(1)
# s.push(2)
# s.push(3)
# s.display()
# top = s.top()
# print("\nTop element is:", top)

# poped=s.pop()
# print("\nPopped element is:", poped)
# top = s.top()
# print("\nTop element is:", top)

# s.display()

# print("Moving elements from one stack to another stack")
# source = stack(5)
# destination = stack(5)

# source.push(1)
# source.push(2)
# source.push(3)
# source.push(4)
# source.push(5)
# print("Elements in source stack:")
# source.display()

# while not source.isempty():
#     element = source.pop()
#     destination.push(element)

# print("\nElements in destination stack:")
# destination.display()

print("Positive numbers:")

s = stack(10)

s.push(-7)
s.push(10)
s.push(-3)
s.push(8)
s.push(-5)

temp = stack(10)

while not s.isempty():
    element = s.pop()
    if element > 0:
        print(element, end=" ")
    else:
        temp.push(element)

print()

# Restore negative numbers back to the original stack without changing their order
while not temp.isempty():
    element = temp.pop()
    s.push(element)

print("Remaining stack:")
s.display()


