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
