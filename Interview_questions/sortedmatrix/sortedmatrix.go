package sortedmatrix

import (
	"fmt"
	"math"
)

//2-D matrix sorted if increasing elements in each row, and next column first element is bigger then the last row last element

func IsInSortedMatrix(matr [][]int, numb int) {
	row := len(matr)
	col := len(matr[0])
	start := 0
	end := row*col - 1
	e := row * col

	for start <= end {
		m := int(math.Floor(float64(start+end) / 2))
		i := m % col
		j := m / e
		if matr[i][j] < numb {
			start = m + 1
		} else if matr[i][j] > numb {
			end = m - 1
		} else {
			fmt.Printf("Number %d found in the matrix!", numb)
			return
		}
	}
	fmt.Printf("Number %d did not founded in the matrix!", numb)

}
