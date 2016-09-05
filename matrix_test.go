// Copyright 2014 Steve Kaliski.

package matrix

import (
	"reflect"
	"testing"
)

var (
	err               error
	testMatrix        *Matrix
	testIdentity      *Matrix
	sampleIdentity, _ = New([][]float64{[]float64{1, 0}, []float64{0, 1}})
	sampleSquare, _   = New([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}, []float64{7, 8, 9}})
)

func TestNewValid(t *testing.T) {
	testMatrix, err = New([][]float64{[]float64{1, 2, 3}, []float64{2, 3, 4}})
	if err != nil {
		t.Error("unexpected result", err)
	}
}

func TestNewInvalid(t *testing.T) {
	_, err = New([][]float64{[]float64{1, 2}, []float64{2, 3, 4}})
	if err == nil && err != errRowsMustBeSameSize {
		t.Error("exepected error")
	}
}

func TestNewIdentity(t *testing.T) {
	testIdentity, _ = NewIdentity(2)
	if !testIdentity.IsSquare() {
		t.Error("failed to create n x n matrix")
	}
	if !testIdentity.IsEqual(sampleIdentity) {
		t.Error("failed to create identity matrix")
	}
}

func TestGetRowCount(t *testing.T) {
	if testMatrix.GetRowCount() != 2 {
		t.Error("failed to get proper row count")
	}
}

func TestGetRowValid(t *testing.T) {
	row, err := testMatrix.GetRow(0)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(row, []float64{1, 2, 3}) {
		t.Error("failed to get proper row")
	}
}

func TestGetRowInvalid(t *testing.T) {
	_, err := testMatrix.GetRow(2)
	if err == nil {
		t.Error("did not throw error")
	}
}

func TestGetColumnValid(t *testing.T) {
	col, err := testMatrix.GetColumn(0)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(col, []float64{1, 2}) {
		t.Error("failed to get proper column")
	}
}

func TestGetColumnInvalid(t *testing.T) {
	_, err := testMatrix.GetColumn(5)
	if err == nil {
		t.Error("did not throw error")
	}
}

func TestSetElementValid(t *testing.T) {
	err := sampleSquare.SetElement(0, 0, 2)
	if err != nil {
		t.Error(err)
	}
}

func TestSetElementInvalid(t *testing.T) {
	err := sampleSquare.SetElement(-1, 0, 2)
	if err == nil {
		t.Error("expected error")
	}
}

func TestScale(t *testing.T) {
	testMatrix.Scale(2)
	row, _ := testMatrix.GetRow(0)
	if !reflect.DeepEqual(row, []float64{2, 4, 6}) {
		t.Error("failed to properly scale matrix")
	}
}

func TestTranspose(t *testing.T) {
	testMatrix.Transpose()
	row, _ := testMatrix.GetRow(0)
	if !reflect.DeepEqual(row, []float64{2, 4}) {
		t.Error("failed to properly transpose matrix")
	}
}

func TestDeterminantValid(t *testing.T) {
	det, err := sampleSquare.Determinant()
	if err != nil {
		t.Error(err)
	}
	if det != -15 {
		t.Error("unexepted result", det)
	}
}

func TestDeterminantInvalid(t *testing.T) {
	_, err := testMatrix.Determinant()
	if err == nil {
		t.Error("unexpected result, matrix is not square")
	}
}

func TestIsSquare(t *testing.T) {
	if testMatrix.IsSquare() {
		t.Error("unexpected result, matrix is not square")
	}
}

func TestIsEqual(t *testing.T) {
	if !testIdentity.IsEqual(sampleIdentity) {
		t.Error("2 x 2 identity matrices should be equal")
	}
}

func TestAdd(t *testing.T) {
	addMatrix, _ := New([][]float64{[]float64{3, 4}, []float64{4, 5}, []float64{5, 6}})
	err := addMatrix.Add(testMatrix)
	if err != nil {
		t.Error(err)
	}

	expectedResult, _ := New([][]float64{[]float64{5, 8}, []float64{8, 11}, []float64{11, 14}})
	if !addMatrix.IsEqual(expectedResult) {
		t.Error("failed to add matrices correctly")
	}
}
