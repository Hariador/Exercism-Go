package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	numCows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", e.numCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		nErr := errors.New("division by zero")
		return 0.0, nErr
	}
	amount, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			if amount < 0 {
				nErr := errors.New("negative fodder")
				return 0.0, nErr
			}
			if amount > 0 {
				amount = amount * 2
			}

		} else {
			return 0, err
		}
	}
	if amount < 0 {
		nErr := errors.New("negative fodder")
		return 0.0, nErr
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows}
	}

	return amount / float64(cows), nil

}
