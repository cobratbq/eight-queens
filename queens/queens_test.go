package queens

import (
	"testing"
)

func TestSolve2(t *testing.T) {
	result := Solve(2)
	if result != nil {
		t.Fatal("Did not expect to get a solution.")
	}
}

func TestSolve3(t *testing.T) {
	result := Solve(3)
	if result != nil {
		t.Fatal("Did not expect to get a solution.")
	}
}

func TestSolve4(t *testing.T) {
	result := Solve(4)
	if result == nil {
		t.Fatal("Expected a solution.")
	}
}

func TestSolve5(t *testing.T) {
	result := Solve(5)
	if result == nil {
		t.Fatal("Expected a solution.")
	}
}

func TestSolve8(t *testing.T) {
	result := Solve(8)
	if result == nil {
		t.Fatal("Expected a solution.")
	}
}

func BenchmarkSolve5(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Solve(5)
	}
}

func BenchmarkSolve8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(8)
	}
}
