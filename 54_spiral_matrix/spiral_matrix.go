// 解题的核心思路是按照右、下、左、上的顺序遍历数组，并使用四个变量圈定未遍历元素的边界：
package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	leftBound := 0
	rightBound := n - 1
	upBound := 0
	downBound := m - 1
	res := make([]int, 0)
	for ; len(res) < m * n; {
		if upBound <= downBound {
			for i:=leftBound; i<= rightBound; i++ {
				res = append(res, matrix[upBound][i])
			}
			upBound++
		}

		if leftBound <= rightBound {
			for i := upBound; i <= downBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}

		if upBound <= downBound {
			for i:= rightBound; i>=leftBound; i-- {
				res = append(res, matrix[downBound][i])
			}
			downBound--
		}

		if leftBound <= rightBound {
			for i := downBound; i>= upBound; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}
	return res
}
