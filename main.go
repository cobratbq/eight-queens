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
	Solve(12)
}

func Solve(size int) {
	solution := make([]uint8, 0)
	constraints := make([]uint8, size)
	solve(size, solution, constraints)
}

func solve(size int, solution []uint8, constraints []uint8) {

	if len(solution) >= size {
		//base case: all queens placed, we're done
		fmt.Printf("Solution found: %v\n", solution)
		return
	}

	nextConstr, potentials := prepareNext(size, constraints)

	//check potential queen positions
	for _, pos := range potentials {

		//copy constraints for next step
		tryConstr := make([]uint8, 0, size)
		tryConstr = append(tryConstr, nextConstr...)

		//copy current solution for further exploration in next step
		trySolution := make([]uint8, 0, len(solution)+1)
		trySolution = append(trySolution, solution...)
		trySolution = append(trySolution, pos)

		//try cell 'pos' and prepare next step constraints
		tryConstr[pos] |= CONSTANT
		if pos > 0 {
			tryConstr[pos-1] |= DECREASING
		}
		if pos < uint8(size-1) {
			tryConstr[pos+1] |= INCREASING
		}

		//try to solve with current position
		solve(size, trySolution, tryConstr)
	}

	return
}

func prepareNext(size int, currentConstr []uint8) (nextConstr []uint8, potentials []uint8) {
	nextConstr = make([]uint8, size)
	potentials = make([]uint8, 0)

	for i, constraint := range currentConstr {

		if constraint == UNUSED {
			//If a cell is completely empty, record it as a possible next step.
			potentials = append(potentials, uint8(i))
		} else {
			//If a cell isn't completely empty, translate the constraints to the next step.

			if constraint&CONSTANT == CONSTANT {
				//mark position for prior queen's vertical influence
				nextConstr[i] |= CONSTANT
			}
			if i < size-1 && constraint&INCREASING == INCREASING {
				//mark position for prior queen's forward diagonal influence
				nextConstr[i+1] |= INCREASING
			}
			if i > 0 && constraint&DECREASING == DECREASING {
				//mark position for prior queen's backward diagonal influence
				nextConstr[i-1] |= DECREASING
			}
		}
	}

	return
}
