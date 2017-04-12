package Leetcode

//solution with less code
func islandPerimeter(grid [][]int) int {
	if len(grid)==0 {return 0}
	if len(grid[0])==0 {return 0}
	nr := len(grid)
	nc := len(grid[0])
	numIslands := 0
	numSB := 0

	for i:=0; i<nr; i++{
		for j:=0; j<nc; j++{
			if grid[i][j] == 1 {
				numIslands++
				if i<nr-1 {
					numSB += grid[i+1][j]
				}
				if j<nc-1{
					numSB += grid[i][j+1]
				}
			}
		}
	}

	return numIslands*4 - numSB*2
}

//solution using if-else
func islandPerimeter(grid [][]int) int {
	if len(grid)==0 {return 0}
	if len(grid[0])==0 {return 0}
	nr := len(grid)
	nc := len(grid[0])
	res := 0

	for i:=0; i<nr; i++{
		for j:=0; j<nc; j++{
			if grid[i][j] == 1 {
				var u, r, d, l int
				if i==0 {
					u = 0
				} else{
					u = grid[i-1][j]
				}
				if i==nr-1 {
					d = 0
				} else{
					d = grid[i+1][j]
				}
				if j==nc-1 {
					r = 0
				} else{
					r = grid[i][j+1]
				}
				if j==0 {
					l = 0
				} else{
					l = grid[i][j-1]
				}
				res += 4-u-r-d-l
			}
		}
	}

	return res
}

//solution using grid expasion
func islandPerimeter(grid [][]int) int {
	if len(grid)==0 {return 0}
	if len(grid[0])==0 {return 0}
	nr := len(grid)
	nc := len(grid[0])
	res := 0

	expGrid := make([][]int, nr+2)
	expGrid[0] = make([]int, nc+2)
	expGrid[nr+1] = make([]int, nc+2)

	for i:=1;i<=nr;i++{
		expGrid[i] = make([]int, nc+2)
		for j:=1;j<=nc;j++{
			expGrid[i][j] = grid[i-1][j-1]
		}
	}

	for i:=1; i<=nr; i++{
		for j:=1; j<=nc; j++{
			if expGrid[i][j] == 1 {
				res += 4-expGrid[i-1][j]- expGrid[i][j+1]-expGrid[i+1][j]-expGrid[i][j-1]
			}
		}
	}
	return res
}
