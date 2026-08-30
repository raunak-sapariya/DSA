class Queue:

    def __init__(self, size):
        self.size = size
        self.queue = [None] * size
        self.head = self.tail = -1

    def enqueue(self, data):
        if (self.tail == self.size-1):
            print("Full")
            return
        elif (self.head == -1):
            self.head = self.tail = 0
            self.queue[self.tail] = data
        else:
            self.tail += 1
            self.queue[self.tail] = data

    def dequeue(self):
        if (self.head == -1):
            print("Empty")
            return
        elif (self.head == self.tail):
            temp = self.queue[self.head]
            self.head = -1
            self.tail = -1
            return temp
        else:
            temp = self.queue[self.head]
            self.head = self.head+1
            return temp

    def display(self):
        if (self.head == -1):
            print("Empty")
        else:
            for i in range(self.head, self.tail+1):
                print(self.queue[i], end=" ")
            print()

q = Queue(5)
q.enqueue(1)
q.enqueue(2)
q.enqueue(3)
q.display()

q.dequeue()
q.dequeue()
q.display()
