package tree_and_graph

func numIslands(grid [][]byte) int {
	//return numIslans_1(grid)
	//return numIslans_2(grid)
	return numIslans_3(grid)
}

// 深度遍历
func numIslans_1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var num int
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])
	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if i+1 < m && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if j+1 < n && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

// 广度遍历
func numIslans_2(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	var num int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				num++
				grid[i][j] = '0'
				var queue [][2]int
				queue = append(queue, [2]int{i, j})
				for len(queue) != 0 {
					item := queue[0]
					queue = queue[1:]
					ii := item[0]
					jj := item[1]
					if ii-1 >= 0 && grid[ii-1][jj] == '1' {
						queue = append(queue, [2]int{ii - 1, jj})
						// 必须在此赋值为'0' 否则会导致无限广度搜索
						grid[ii-1][jj] = '0'
					}
					if ii+1 < m && grid[ii+1][jj] == '1' {
						queue = append(queue, [2]int{ii + 1, jj})
						grid[ii+1][jj] = '0'
					}
					if jj-1 >= 0 && grid[ii][jj-1] == '1' {
						queue = append(queue, [2]int{ii, jj - 1})
						grid[ii][jj-1] = '0'
					}
					if jj+1 < n && grid[ii][jj+1] == '1' {
						queue = append(queue, [2]int{ii, jj + 1})
						grid[ii][jj+1] = '0'
					}
				}
			}
		}
	}
	return num
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind() *UnionFind {
	uf := UnionFind{}
	return &uf
}

func (uf *UnionFind) SetParent(parent []int) {
	uf.parent = parent
}

func (uf *UnionFind) SetRank(rank []int) {
	uf.rank = rank
}

func (uf *UnionFind) find(index int) int {
	if uf.parent[index] != index {
		// 压缩路径
		uf.parent[index] = uf.find(uf.parent[index])
	}
	return uf.parent[index]
}

func (uf *UnionFind) Union(x, y int) {
	rootx := uf.find(x)
	rooty := uf.find(y)
	if rootx != rooty {
		// 合并
		// rooty分值大于rootx, 分值大的为父节点
		if uf.rank[rootx] < uf.rank[rooty] {
			uf.parent[rootx] = rooty
		} else if uf.rank[rootx] > uf.rank[rooty] {
			uf.parent[rooty] = rootx
		} else {
			uf.parent[rooty] = rootx
			uf.rank[rootx]++
		}
	}
}

func (uf *UnionFind) GetCount() int {
	var count int
	for i := 0; i < len(uf.parent); i++ {
		if uf.parent[i] == i {
			count++
		}
	}
	return count
}

func numIslans_3(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	uf := NewUnionFind()
	parent := make([]int, 0, m*n)
	rank := make([]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent = append(parent, i*n+j)
			} else {
				parent = append(parent, -1)
			}
			rank = append(rank, 0)
		}
	}
	uf.SetParent(parent)
	uf.SetRank(rank)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(i*n+j, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union(i*n+j, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*n+j, i*n+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j, i*n+j+1)
				}
			}
		}
	}
	return uf.GetCount()
}
