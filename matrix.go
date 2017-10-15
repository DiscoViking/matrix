package matrix

import "errors"

type Matrix struct {
	rows    int
	columns int
	cells   []float64
}

func New(rows, columns int) Matrix {
	return Matrix{
		rows:    rows,
		columns: columns,
		cells:   make([]float64, rows*columns),
	}
}

func NewWithValues(rows, columns int, vals ...float64) (Matrix, error) {
	if len(vals) != rows*columns {
		return Matrix{}, errors.New("Must provide rows*columns values")
	}

	m := New(rows, columns)
	m.cells = vals
	return m, nil
}

func (m Matrix) Dimensions() (int, int) {
	return m.rows, m.columns
}

func (m Matrix) Set(x, y int, v float64) error {
	ix, err := m.findCell(x, y)
	if err != nil {
		return err
	}

	m.cells[ix] = v
	return nil
}

func (m Matrix) Get(x, y int) (float64, error) {
	ix, err := m.findCell(x, y)
	if err != nil {
		return 0, err
	}

	return m.cells[ix], nil
}

func (m Matrix) Equals(n Matrix) bool {
	if m.rows != n.rows || m.columns != n.columns {
		return false
	}

	for i := 0; i < m.rows*m.columns; i++ {
		if m.cells[i] != n.cells[i] {
			return false
		}
	}

	return true
}

func (m Matrix) Add(n Matrix) (Matrix, error) {
	if m.rows != n.rows || m.columns != n.columns {
		return Matrix{}, errors.New("matrix dimensions must match")
	}

	sum := New(m.rows, m.columns)
	for i := 0; i < m.rows*m.columns; i++ {
		sum.cells[i] = m.cells[i] + n.cells[i]
	}

	return sum, nil
}

func (m Matrix) Subtract(n Matrix) (Matrix, error) {
	if m.rows != n.rows || m.columns != n.columns {
		return Matrix{}, errors.New("matrix dimensions must match")
	}

	sum := New(m.rows, m.columns)
	for i := 0; i < m.rows*m.columns; i++ {
		sum.cells[i] = m.cells[i] - n.cells[i]
	}

	return sum, nil
}

func (m Matrix) ScalarMultiply(x float64) Matrix {
	mult := New(m.rows, m.columns)
	for i := 0; i < m.rows*m.columns; i++ {
		mult.cells[i] = m.cells[i] * x
	}

	return mult
}

func (m Matrix) Multiply(n Matrix) (Matrix, error) {
	if m.columns != n.rows {
		return Matrix{}, errors.New("left matrix must have same number of columns as right matrix has rows")
	}

	product := New(m.rows, n.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < n.columns; j++ {
			var v float64 = 0
			for k := 0; k < m.columns; k++ {
				v += m.cells[k+i*m.columns] * n.cells[j+k*n.columns]
			}
			product.Set(i, j, v)
		}
	}

	return product, nil
}

func (m Matrix) findCell(r, c int) (int, error) {
	if r >= m.rows || r < 0 {
		return 0, errors.New("row out of bounds")
	}

	if c >= m.columns || r < 0 {
		return 0, errors.New("column out of bounds")
	}

	return c + (r * m.columns), nil
}
