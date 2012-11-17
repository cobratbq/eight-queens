package queens

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	num := SolveAll(1)
	if num != 1 {
		t.Fatal("Expected 1 solutions.")
	}
}

func TestSolve2(t *testing.T) {
	num := SolveAll(2)
	if num != 0 {
		t.Fatal("Expected 0 solutions.")
	}
}

func TestSolve3(t *testing.T) {
	num := SolveAll(3)
	if num != 0 {
		t.Fatal("Expected 0 solutions.")
	}
}

func TestSolve4(t *testing.T) {
	num := SolveAll(4)
	if num != 2 {
		t.Fatal("Expected 2 solutions.")
	}
}

func TestSolve5(t *testing.T) {
	num := SolveAll(5)
	if num != 10 {
		t.Fatal("Expected 10 solutions.")
	}
}

func TestSolve6(t *testing.T) {
	num := SolveAll(6)
	if num != 4 {
		t.Fatal("Expected 4 solutions.")
	}
}

func TestSolve7(t *testing.T) {
	num := SolveAll(7)
	if num != 40 {
		t.Fatal("Expected 40 solutions.")
	}
}

func TestSolve8(t *testing.T) {
	num := SolveAll(8)
	if num != 92 {
		t.Fatal("Expected 92 solutions.")
	}
}

func TestSolve9(t *testing.T) {
	num := SolveAll(9)
	if num != 352 {
		t.Fatal("Expected 352 solutions.")
	}
}

func TestSolve10(t *testing.T) {
	num := SolveAll(10)
	if num != 724 {
		t.Fatal("Expected 724 solutions.")
	}
}

func TestSolve11(t *testing.T) {
	num := SolveAll(11)
	if num != 2680 {
		t.Fatal("Expected 2680 solutions.")
	}
}

func TestSolve12(t *testing.T) {
	num := SolveAll(12)
	if num != 14200 {
		t.Fatal("Expected 14200 solutions.")
	}
}
