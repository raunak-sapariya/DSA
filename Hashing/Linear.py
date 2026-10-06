class Liner:
    def __init__(self, size):
        self.size = size
        self.table = [None] * size

    def h1(self, key):
        return key % self.size

    def insert(self, key):
        index = self.h1(key)

        while self.table[index] is not None:
            index = (index + 1) % self.size

        self.table[index] = key

    def search(self, key):
        index = self.h1(key)

        while self.table[index] is not None:
            if self.table[index] == key:
                return index
        
            index = (index + 1) % self.size

        return -1

    def delete(self, key):
        index = self.search(key)

        if index != -1:
            self.table[index] = None

    def display(self):
        for i in range(self.size):
            print(i, "->", self.table[i])

    def rehash(self, new_size):
        old_table = self.table
        self.size = new_size
        self.table = [None] * new_size

        for key in old_table:
            if key is not None:
                self.insert(key)

h = Liner(11)
# keys = [50, 700, 76, 85, 92, 73, 101, 12, 11, 22, 33]

# for key in keys:
#     h.insert(key)

# h.display()

# print("Search 85:", h.search(85))

# h.delete(85)

# print("After deleting 85:")
# h.display()

# h.insert(85)
# h.display()

h = Liner(5)

for key in [10, 20, 30, 40]:
    h.insert(key)

print("Before rehashing:")
h.display()

h.rehash(11)

print("\nAfter rehashing:")
h.display()