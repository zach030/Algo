package union_find

type UnionFind struct {
	count int
	parent []int
	size []int
}

func NewUF(n int)*UnionFind{
	uf := &UnionFind{count:  n,parent: make([]int,n),size: make([]int,n)}
	for i := 0; i < n; i++ {
		uf.parent[i]=i
		uf.size[i]=1
	}
	return uf
}

func (u *UnionFind)Union(a,b int){
	ra,rb := u.Find(a),u.Find(b)
	if ra==rb{
		return
	}
	if u.size[ra]>u.size[rb]{
		u.parent[rb]=ra
		u.size[rb]+=u.size[ra]
	}else{
		u.parent[ra]=rb
		u.size[ra]+=u.size[rb]
	}
	u.count--
}

// 返回某节点的根节点
func (u *UnionFind)Find(x int)int{
	for u.parent[x]!=x{
		u.parent[x]=u.parent[u.parent[x]]
		x=u.parent[x]
	}
	return x
}

func (u *UnionFind)Connected(a,b int)bool {
	return u.Find(a) == u.Find(b)
}

func (u *UnionFind)Count()int{
	return u.count
}


// 130 被围绕的区域
func solve(board [][]byte)  {
	if len(board)==0{
		return
	}
	m,n:=len(board),len(board[0])//m 行 n 列
	uf :=NewUF(m*n+1)
	dummy :=m*n
	// 先找边界0，与dummy相连
	for i := 0; i < n; i++ {
		if board[0][i]=='O'{
			uf.Union(i,dummy)
		}
		if board[m-1][i]=='O'{
			uf.Union(n*(m-1)+i,dummy)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O'{
			uf.Union(i * n, dummy)
		}
		if board[i][n - 1] == 'O' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	// 方向数组 d 是上下左右搜索的常用手法
	d := [][]int{{1,0}, {0,1}, {0,-1}, {-1,0}}
	for  i := 1; i < m - 1; i++{
		for j := 1; j < n - 1; j++{
			if board[i][j] == 'O'{
				// 将此 O 与上下左右的 O 连通
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == 'O' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	// 所有不和 dummy 连通的 O，都要被替换
	for i := 1; i < m - 1; i++{
		for j := 1; j < n - 1; j++{
			if !uf.Connected(dummy, i * n + j) {
				board[i][j] = 'X'
			}
		}
	}
}

// 128 最长连续序列
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

// 岛屿数量
// class Solution {
//    class UnionFind {
//        int count;
//        int[] parent;
//        int[] rank;
//
//        public UnionFind(char[][] grid) {
//            count = 0;
//            int m = grid.length;
//            int n = grid[0].length;
//            parent = new int[m * n];
//            rank = new int[m * n];
//            for (int i = 0; i < m; ++i) {
//                for (int j = 0; j < n; ++j) {
//                    if (grid[i][j] == '1') {
//                        parent[i * n + j] = i * n + j;
//                        ++count;
//                    }
//                    rank[i * n + j] = 0;
//                }
//            }
//        }
//
//        public int find(int i) {
//            if (parent[i] != i) parent[i] = find(parent[i]);
//            return parent[i];
//        }
//
//        public void union(int x, int y) {
//            int rootx = find(x);
//            int rooty = find(y);
//            if (rootx != rooty) {
//                if (rank[rootx] > rank[rooty]) {
//                    parent[rooty] = rootx;
//                } else if (rank[rootx] < rank[rooty]) {
//                    parent[rootx] = rooty;
//                } else {
//                    parent[rooty] = rootx;
//                    rank[rootx] += 1;
//                }
//                --count;
//            }
//        }
//
//        public int getCount() {
//            return count;
//        }
//    }
//
//    public int numIslands(char[][] grid) {
//        if (grid == null || grid.length == 0) {
//            return 0;
//        }
//
//        int nr = grid.length;
//        int nc = grid[0].length;
//        int num_islands = 0;
//        UnionFind uf = new UnionFind(grid);
//        for (int r = 0; r < nr; ++r) {
//            for (int c = 0; c < nc; ++c) {
//                if (grid[r][c] == '1') {
//                    grid[r][c] = '0';
//                    if (r - 1 >= 0 && grid[r-1][c] == '1') {
//                        uf.union(r * nc + c, (r-1) * nc + c);
//                    }
//                    if (r + 1 < nr && grid[r+1][c] == '1') {
//                        uf.union(r * nc + c, (r+1) * nc + c);
//                    }
//                    if (c - 1 >= 0 && grid[r][c-1] == '1') {
//                        uf.union(r * nc + c, r * nc + c - 1);
//                    }
//                    if (c + 1 < nc && grid[r][c+1] == '1') {
//                        uf.union(r * nc + c, r * nc + c + 1);
//                    }
//                }
//            }
//        }
//
//        return uf.getCount();
//    }
//}
//
//作者：LeetCode
//链接：https://leetcode-cn.com/problems/number-of-islands/solution/dao-yu-shu-liang-by-leetcode/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
