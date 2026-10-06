DELETED = "DELETED"
class Quadratic:
    def __init__(self, size):
        self.size = size
        self.table = [None] * size

    def h1(self, key):
        return key % self.size

    def insert(self, key):
        i =0

        while i < self.size:
            index = (self.h1(key) + i * i) % self.size
            if self.table[index] is None:
                self.table[index] = key
                return

            i += 1

        print("Hash table is full")

    def search(self, key):
        i =0

        while i < self.size:
            index = (self.h1(key) + i * i) % self.size

            if self.table[index] is None:
                return -1
            
            if self.table[index] == key:
                return index

            i += 1

        return -1

    def delete(self, key):
        index = self.search(key)

        if index != -1:
            self.table[index] = DELETED
            print("Deleted", key)
        else:
            print("not found")

    def display(self):
        for i in range(self.size):
            print(i, "->", self.table[i])

h = Quadratic(11)

keys = [50, 700, 76, 85, 92, 73, 101, 12, 11, 22, 33]

for key in keys:
    h.insert(key)

h.display()

print("Search 85:", h.search(85))

h.delete(85)

print("After deleting 85:")
h.display()

h.insert(85)
h.display()