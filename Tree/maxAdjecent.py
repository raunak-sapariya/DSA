"""
Find the maximum adjacent value for each node in a graph represented by an adjacency matrix.

[
    [0, 6, 0, 0, 16],
    [5, 0, 11, 12, 16],
    [0, 6, 0, 12, 0],
    [0, 6, 11, 0, 16],
    [5, 6, 0, 12, 0]
]

result will be
node - max adjacent value - max adjacent node
0 - 16 - 4
1 - 16 - 4
2 - 12 - 3
3 - 16 - 4
4 - 12 - 3
"""


def max_adjacent(matrix):
    n = len(matrix)
    result = []

    for i in range(n):
        max_value = float('-inf')
        max_node = -1

        for j in range(n):
            if matrix[i][j] != 0:
                if matrix[i][j] > max_value:
                    max_value = matrix[i][j]
                    max_node = j

        result.append((i, max_value, max_node))

    return result

if __name__ == "__main__":
    adjacency_matrix = [
        [0, 6, 0, 0, 16],
        [5, 0, 11, 12, 16],
        [0, 6, 0, 12, 0],
        [0, 6, 11, 0, 16],
        [5, 6, 0, 12, 0]
    ]

    print("Adjacency Matrix:")
    for row in adjacency_matrix:
        print(row)

    result = max_adjacent(adjacency_matrix)

    for node, max_value, max_node in result:
        print(f"\nNode {node} - Max Adjacent Value: {max_value} - Max Adjacent Node: {max_node}")
