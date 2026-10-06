class BinaryTree:

    def __init__(self, size):
        self.capacity = size
        self.tree = [None] * size

    def insert_root(self, value):
        if self.tree[0] is not None:
            print("Root already exists!")
            return
        self.tree[0] = value

    def _insert_child(self, parent_index, child_index, value, side):
        if not 0 <= parent_index < self.capacity or self.tree[parent_index] is None:
            print(f"Cannot insert {side} child, parent missing!")
            return
        if child_index >= self.capacity:
            print("Index out of range!")
            return
        self.tree[child_index] = value

    def insert_left(self, parent_index, value):
        self._insert_child(parent_index, 2 * parent_index + 1, value, "left")

    def insert_right(self, parent_index, value):
        self._insert_child(parent_index, 2 * parent_index + 2, value, "right")

    def display(self):
        print("Binary Tree (array representation):")
        for index, value in enumerate(self.tree):
            print(f"[{index}] {value}" if value is not None else "-", end = " ")

    def display_inorder(self, index=0):
        if index >= self.capacity or self.tree[index] is None:
            return
        self.display_inorder(2 * index + 1)  # Left child
        print(self.tree[index], end=" ")
        self.display_inorder(2 * index + 2)  # Right child

    def display_preorder(self, index=0):
        if index >= self.capacity or self.tree[index] is None:
            return
        print(self.tree[index], end=" ")
        self.display_preorder(2 * index + 1)  # Left child
        self.display_preorder(2 * index + 2)  # Right child

    def display_postorder(self, index=0):
        if index >= self.capacity or self.tree[index] is None:
            return
        self.display_postorder(2 * index + 1)  # Left child
        self.display_postorder(2 * index + 2)  # Right child
        print(self.tree[index], end=" ")

    def display_level_order(self):
        print("Level Order Traversal:")
        for value in self.tree:
            if value is not None:
                print(value, end=" ")
        print()
        


if __name__ == "__main__":
    bt = BinaryTree(15)

    bt.insert_root(1)
    bt.insert_left(0, 2)
    bt.insert_right(0, 3)
    bt.insert_left(1, 4)
    bt.insert_right(1, 5)
    bt.insert_left(2, 6)
    bt.insert_right(2, 7)

    bt.display()