DELETED = "DELETED"

class Double:
    def __init__(self, size, R=7):
        self.size = size
        self.R = R
        self.table = [None] * size

    def h1(self, key):
        return key % self.size

    def h2(self, key):
        return self.R - (key % self.R)

    def insert(self, key):
        i =0
        
        while i < self.size:
            index = (self.h1(key) + i * self.h2(key)) % self.size
            if self.table[index] is None or self.table[index] == DELETED:
                self.table[index] = key
                return
        
        i += 1
        
        print("Hash table is full")

    def search(self, key):
        i = 0
        
        while i < self.size:
            index = (self.h1(key) + i * self.h2(key)) % self.size
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

h = Double(11)
keys = [50, 700, 76, 85, 92]

for key in keys:
    h.insert(key)

h.display()

print("Search 85:", h.search(85))

h.delete(85)

print("After deleting 85:")
h.display()
print()
h.insert(85)
h.display()

DELETED = "DELETED"

class Double:
    def __init__(self, size, R=7):
        self.size = size
        self.R = R
        self.table = [None] * size

    def h1(self, key):
        return key % self.size

    def h2(self, key):
        return self.R - (key % self.R)

    def insert(self, key):
        i =0
        
        while i < self.size:
            index = (self.h1(key) + i * self.h2(key)) % self.size
            if self.table[index] is None or self.table[index] == DELETED:
                self.table[index] = key
                return
        
        i += 1
        
        print("Hash table is full")

    def search(self, key):
        i = 0
        
        while i < self.size:
            index = (self.h1(key) + i * self.h2(key)) % self.size
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

h = Double(11)
keys = [50, 700, 76, 85, 92]

for key in keys:
    h.insert(key)

h.display()

print("Search 85:", h.search(85))

h.delete(85)

print("After deleting 85:")
h.display()
print()
h.insert(85)
h.display()