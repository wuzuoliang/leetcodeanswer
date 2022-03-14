package Code

import "testing"

func Test886(t *testing.T) {
	t.Log(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	t.Log(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
}

var coudlSplit bool
var colorMap886 []bool
var visited886 []bool

func possibleBipartition(n int, dislikes [][]int) bool {
	coudlSplit = true
	colorMap886 = make([]bool, n+1)
	visited886 = make([]bool, n+1)
	graph := buildGraph(n, dislikes)
	for i := 1; i <= n; i++ {
		if visited886[i] == true {
			continue
		}
		dfs886(graph, i)
	}
	return coudlSplit
}

func buildGraph(n int, dislikes [][]int) (grap [][]int) {
	grap = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		grap[i] = make([]int, 0)
	}
	for _, edge := range dislikes {
		a := edge[0]
		b := edge[1]
		grap[a] = append(grap[a], b)
		grap[b] = append(grap[b], a)
	}
	return grap
}

func dfs886(graph [][]int, cur int) {
	if !coudlSplit {
		return
	}

	visited886[cur] = true

	for _, next := range graph[cur] {
		if visited886[next] == false {
			// 相邻节点 w 没有被访问过
			// 那么应该给节点 w 涂上和节点 v 不同的颜色
			colorMap886[next] = !colorMap886[cur]
			// 继续遍历 w
			dfs886(graph, next)
		} else {
			// 相邻节点 w 已经被访问过
			// 根据 v 和 w 的颜色判断是否是二分图
			if colorMap886[next] == colorMap886[cur] {
				// 若相同，则此图不是二分图
				coudlSplit = false
			}
		}
	}
}

/**
private boolean ok = true;
private boolean[] color;
private boolean[] visited;

public boolean possibleBipartition(int n, int[][] dislikes) {
    // 图节点编号为 1...n
    color = new boolean[n + 1];
    visited = new boolean[n + 1];
    // 转化成邻接表表示图结构
    List<Integer>[] graph = buildGraph(n, dislikes);

    for (int v = 1; v <= n; v++) {
        if (!visited[v]) {
            traverse(graph, v);
        }
    }

    return ok;
}

// 建图函数
private List<Integer>[] buildGraph(int n, int[][] dislikes) {
    // 图节点编号为 1...n
    List<Integer>[] graph = new LinkedList[n + 1];
    for (int i = 1; i <= n; i++) {
        graph[i] = new LinkedList<>();
    }
    for (int[] edge : dislikes) {
        int v = edge[1];
        int w = edge[0];
        // 「无向图」相当于「双向图」
        // v -> w
        graph[v].add(w);
        // w -> v
        graph[w].add(v);
    }
    return graph;
}

// 和之前的 traverse 函数完全相同
private void traverse(List<Integer>[] graph, int v) {
    if (!ok) return;
    visited[v] = true;
    for (int w : graph[v]) {
        if (!visited[w]) {
            color[w] = !color[v];
            traverse(graph, w);
        } else {
            if (color[w] == color[v]) {
                ok = false;
            }
        }
    }
}
*/
