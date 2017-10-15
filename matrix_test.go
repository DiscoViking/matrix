package matrix

import "testing"

func TestNewMatrixInternalStructure(t *testing.T) {
	m := New(3, 4)
	if m.rows != 3 {
		t.Errorf("Matrix width should be 3")
	}

	if m.columns != 4 {
		t.Errorf("Matrix height should be 4")
	}

	if len(m.cells) != 12 {
		t.Fatalf("Matrix should have 12 cells")
	}

	for i := 0; i < 12; i++ {
		if m.cells[i] != 0 {
			t.Errorf("Matrix cell %v should equal 0", i)
		}
	}
}

func TestSetAndGet(t *testing.T) {
	m := New(2, 2)
	err := m.Set(0, 1, 1.2)
	if err != nil {
		t.Fatalf("Error setting cell (0, 1): %v", err)
	}

	assertCellHasValue(t, m, 0, 1, 1.2)
}

func TestNewWithValues(t *testing.T) {
	m, err := NewWithValues(2, 2,
		1.1, 2.1,
		1.2, 2.2,
	)
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	assertCellHasValue(t, m, 0, 0, 1.1)
	assertCellHasValue(t, m, 0, 1, 2.1)
	assertCellHasValue(t, m, 1, 0, 1.2)
	assertCellHasValue(t, m, 1, 1, 2.2)
}

func TestDimensions(t *testing.T) {
	m := New(5, 9)
	w, h := m.Dimensions()

	if w != 5 {
		t.Errorf("width should be 5, got %v", w)
	}

	if h != 9 {
		t.Errorf("height should be 9, got %v", h)
	}
}

func TestEquals(t *testing.T) {
	m, err := NewWithValues(2, 3,
		1, 2, 3,
		4, 5, 6,
	)
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	n, err := NewWithValues(2, 3,
		1, 2, 3,
		4, 5, 6,
	)
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	if !m.Equals(n) {
		t.Errorf("m didn't equal n")
	}

	if !n.Equals(m) {
		t.Errorf("n didn't equal m")
	}
}

func TestAdd(t *testing.T) {
	m := mustCreate(t, 2, 3,
		1, 2, 3,
		4, 5, 6,
	)

	n := mustCreate(t, 2, 3,
		7, 8, 9,
		-1, -2, -3,
	)

	exp := mustCreate(t, 2, 3,
		8, 10, 12,
		3, 3, 3,
	)

	sum, err := m.Add(n)
	if err != nil {
		t.Fatalf("Failed to add matrices: %v", err)
	}

	if !sum.Equals(exp) {
		t.Fatalf("Unexpected result: %v", sum)
	}
}

func TestSubtract(t *testing.T) {
	m := mustCreate(t, 2, 3,
		7, 8, 9,
		-1, -2, -3,
	)

	n := mustCreate(t, 2, 3,
		1, 2, 3,
		4, 5, 6,
	)

	exp := mustCreate(t, 2, 3,
		6, 6, 6,
		-5, -7, -9,
	)

	sub, err := m.Subtract(n)
	if err != nil {
		t.Fatalf("Failed to subtract matrices: %v", err)
	}

	if !sub.Equals(exp) {
		t.Fatalf("Unexpected result: %v", sub)
	}
}

func TestAddFailsWithMismatchedSizes(t *testing.T) {
	m := New(1, 2)
	n := New(1, 3)
	_, err := m.Add(n)
	if err == nil {
		t.Fatalf("Should have hit an error")
	}
}

func TestScalarMultiply(t *testing.T) {
	m := mustCreate(t, 2, 3,
		1, 2, 3,
		4, 5, 6,
	)

	exp := mustCreate(t, 2, 3,
		2, 4, 6,
		8, 10, 12,
	)

	mult := m.ScalarMultiply(2)
	if !mult.Equals(exp) {
		t.Errorf("Expected %v, got %v", exp, mult)
	}
}

func TestMultiply(t *testing.T) {
	m := mustCreate(t, 2, 3,
		1, 2, 3,
		4, 5, 6,
	)
	n := mustCreate(t, 3, 2,
		7, 8,
		9, 10,
		11, 12,
	)
	exp := mustCreate(t, 2, 2,
		58, 64,
		139, 154,
	)

	product, err := m.Multiply(n)
	if err != nil {
		t.Fatalf("Failed to multiply matrices: %v", err)
	}

	if !product.Equals(exp) {
		t.Fatalf("Expected %v, got %v", exp, product)
	}
}

func assertCellHasValue(t *testing.T, m Matrix, x, y int, exp float64) {
	v, err := m.Get(x, y)
	if err != nil {
		t.Errorf("Error getting cell (%v, %v): %v", x, y, err)
	}

	if v != exp {
		t.Errorf("Expected cell (%v, %v) to be %v, got %v", x, y, exp, v)
	}
}

func mustCreate(t *testing.T, rows, columns int, values ...float64) Matrix {
	m, err := NewWithValues(rows, columns, values...)
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	return m
}
