package main

import (
	"fmt"
)

const (
	UNUSED     uint8 = 0x0
	CONSTANT   uint8 = 0x1
	DECREASING uint8 = 0x2
	INCREASING uint8 = 0x4
)

func main() {
	constraints := [8]uint8{UNUSED, UNUSED, UNUSED, UNUSED, UNUSED, UNUSED, UNUSED, UNUSED}
	Solve(constraints)
}

func Solve(constr [8]uint8) {
	solution := make([]uint8, 0)
	solve(solution, constr)
}

func solve(solution []uint8, constr [8]uint8) {

	if len(solution) >= 8 {
		fmt.Printf("Solution found: %v\n", solution)
		return
	}

	nextConstr, potentials := prepareNext(&constr)

	for _, emptyPos := range potentials {
		tryConstr := nextConstr

		trySolution := make([]uint8, 0, len(solution)+1)
		trySolution = append(trySolution, solution...)
		trySolution = append(trySolution, emptyPos)

		tryConstr[emptyPos] |= CONSTANT
		if emptyPos > 0 {
			tryConstr[emptyPos-1] |= DECREASING
		}
		if emptyPos < 7 {
			tryConstr[emptyPos+1] |= INCREASING
		}

		solve(trySolution, tryConstr)
	}

	return
}

func prepareNext(currentConstr *[8]uint8) (nextConstr [8]uint8, potentials []uint8) {
	potentials = make([]uint8, 0)

	for i := uint8(0); i < 8; i++ {

		if currentConstr[i] == UNUSED {
			//If a cell is completely empty, record it as a possible next step.
			potentials = append(potentials, i)
		} else {
			//If a cell isn't completely empty, translate the constraints to the next step.

			if currentConstr[i]&CONSTANT == CONSTANT {
				nextConstr[i] |= CONSTANT
			}
			if i < 7 && currentConstr[i]&INCREASING == INCREASING {
				nextConstr[i+1] |= INCREASING
			}
			if i > 0 && currentConstr[i]&DECREASING == DECREASING {
				nextConstr[i-1] |= DECREASING
			}
		}
	}

	return
}
