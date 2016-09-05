// Copyright 2014 Steve Kaliski.

package matrix

import (
	"errors"
	"reflect"
)

var (
	errPositiveNumberRequired = errors.New("positive number required")
	errIndexOutOfRange        = errors.New("index out of range")
	errDimensionMismatch      = errors.New("matrices dimensions do not match")
	errIsNotSquare            = errors.New("must be a n x n matrix")
	errRowsMustBeSameSize     = errors.New("rows must contain the same number of elements")
)

// Matrix defines a two-dimensional matrix comprised of
// rows and columns.
type Matrix struct {
	data [][]float64
}

// New creates a new matrix.
func New(data [][]float64) (*Matrix, error) {
	var numRow = len(data)
	var numCol int

	if numRow == 0 {
		return nil, errPositiveNumberRequired
	}

	if len(data[0]) == 0 {
		return nil, errPositiveNumberRequired
	}

	numCol = len(data[0])
	for _, row := range data {
		if len(row) != numCol {
			return nil, errRowsMustBeSameSize
		}
	}

	return &Matrix{data: data}, nil
}

// NewIdentity creates a new n x n identity matrix.
func NewIdentity(n int) (*Matrix, error) {
	if n <= 0 {
		return nil, errPositiveNumberRequired
	}

	var data = make([][]float64, n)
	for i := 0; i < len(data); i++ {
		data[i] = make([]float64, n)
		for j := 0; j < len(data[i]); j++ {
			if i == j {
				data[i][j] = 1
			} else {
				data[i][j] = 0
			}
		}
	}

	matrix, err := New(data)
	if err != nil {
		return nil, err
	}

	return matrix, nil
}

// GetRowCount determines the total number of rows in the matrix.
func (m *Matrix) GetRowCount() int {
	return len(m.data)
}

// GetRow gets the indexed row.
func (m *Matrix) GetRow(index int) ([]float64, error) {
	if index < m.GetRowCount() {
		return m.data[index], nil
	}

	return nil, errIndexOutOfRange
}

// GetColumnCount determines the number of columns in the matrix.
func (m *Matrix) GetColumnCount() int {
	return len(m.data[0])
}

// GetColumn gets the indexed column.
func (m *Matrix) GetColumn(index int) ([]float64, error) {
	var column []float64
	rowCount := m.GetRowCount()

	if index < 0 || index > rowCount {
		return column, errIndexOutOfRange
	}

	for i := 0; i < rowCount; i++ {
		column = append(column, m.data[i][index])
	}

	return column, nil
}

// GetElement retrieves an element from the matrix.
func (m *Matrix) GetElement(i, j int) (float64, error) {
	if (i < 0 || i >= m.GetRowCount()) || (j < 0 || j >= m.GetColumnCount()) {
		return 0, errIndexOutOfRange
	}

	return m.data[i][j], nil
}

// SetElement sets an element in the matrix.
func (m *Matrix) SetElement(i, j int, val float64) error {
	if (i < 0 || i >= m.GetRowCount()) || (j < 0 || j >= m.GetColumnCount()) {
		return errIndexOutOfRange
	}

	m.data[i][j] = val
	return nil
}

// Scale executes scalar multiplication on a matrix.
func (m *Matrix) Scale(val float64) {
	for i := 0; i < m.GetRowCount(); i++ {
		for j := 0; j < m.GetColumnCount(); j++ {
			m.data[i][j] = val * m.data[i][j]
		}
	}
}

// Transpose executes a transposition on the matrix.
func (m *Matrix) Transpose() {
	var numRow = m.GetRowCount()
	var numCol = m.GetColumnCount()
	var data = make([][]float64, numCol)

	for i := 0; i < numCol; i++ {
		data[i] = make([]float64, numRow)
		for j := 0; j < numRow; j++ {
			data[i][j] = m.data[j][i]
		}
	}

	m.data = data
}

// Determinant calculates the determinant of the matrix.
// Must be an n x n matrix.
func (m *Matrix) Determinant() (float64, error) {
	if !m.IsSquare() {
		return 0, errIsNotSquare
	}

	var numCol = m.GetColumnCount()
	var numRow = m.GetRowCount()
	var determinant float64
	var diagLeft float64
	var diagRight float64

	for j := 0; j < numCol; j++ {
		diagLeft = m.data[0][j]
		diagRight = m.data[0][j]

		for i := 0; i < numRow; i++ {
			diagRight *= m.data[i][(((j+i)%numCol)+numCol)%numCol]
			diagLeft *= m.data[i][(((j-i)%numCol)+numCol)%numCol]
		}

		determinant += diagRight - diagLeft
	}

	return determinant, nil
}

// IsSquare determines if the matrix is an n x n matrix.
func (m *Matrix) IsSquare() bool {
	return m.GetRowCount() == m.GetColumnCount()
}

// IsEqual determines if a matrix is equal to another matrix.
func (m *Matrix) IsEqual(b *Matrix) bool {
	return reflect.DeepEqual(m.data, b.data)
}

func (m *Matrix) IsSameSize(b *Matrix) bool {
	return m.GetRowCount() == b.GetRowCount() &&
		m.GetColumnCount() == b.GetColumnCount()
}

// Add adds two matrices together.
func (m *Matrix) Add(b *Matrix) error {
	if !m.IsSameSize(b) {
		return errDimensionMismatch
	}

	for i := 0; i < m.GetRowCount(); i++ {
		for j := 0; j < m.GetColumnCount(); j++ {
			current, err := m.GetElement(i, j)
			if err != nil {
				return err
			}
			additive, err := b.GetElement(i, j)
			if err != nil {
				return err
			}
			m.SetElement(i, j, current+additive)
		}
	}

	return nil
}
