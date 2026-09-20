from collections import deque


def bfs(graph, start):
	visited = {start}
	queue = deque([start])
	result = []

	while queue:
		vertex = queue.popleft()
		result.append(vertex)

		for neighbor, connected in enumerate(graph[vertex]):
			if connected and neighbor not in visited:
				visited.add(neighbor)
				queue.append(neighbor)

	return result


n = int(input("Enter number of vertices: "))
print("Enter adjacency matrix:")
graph = [list(map(int, input().split())) for _ in range(n)]
start = int(input(f"Enter starting vertex (0-{n - 1}): "))

if 0 <= start < n:
	print("BFS Traversal:", *bfs(graph, start))
else:
	print("Invalid starting vertex")

