def dfs(graph, visited, v):
    visited[v] = True
    print(v, end=" ")

    for i in range(len(graph)):
        if graph[v][i] == 1 and not visited[i]:
            dfs(graph, visited, i)

n = int(input("Enter the number of vertices: "))
graph = []
for i in range(n):
    while True:
        row = list(map(int, input(
            f"Enter the adjacency row {i} ({n} values): "
        ).split()))
        if len(row) == n:
            break
        print(f"Each row must contain exactly {n} values.")
    graph.append(row)

print(graph)

visited = [False] * n
print("DFS traversal starting from vertex 0:")
dfs(graph, visited, 0)

"""
input:
Enter the number of vertices: 5
Enter the adjacency row 0 (5 values): 0 1 1 0 0
Enter the adjacency row 1 (5 values): 1 0 0 1 1
Enter the adjacency row 2 (5 values): 1 0 0 0 1
Enter the adjacency row 3 (5 values): 0 1 0 0 0
Enter the adjacency row 4 (5 values): 0 1 1 0 0
output:
DFS traversal starting from vertex 0:
0 1 3 4 2
"""