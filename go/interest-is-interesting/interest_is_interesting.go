package interest

const negInterest = 3.213
const smallInterest = 0.5
const medInterest = 1.621
const ballerInterest = 2.475
const medMin = 1000
const ballerMin = 5000

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return float32(negInterest)
	case balance >= 0 && balance < medMin:
		return float32(smallInterest)
	case balance >= medMin && balance < ballerMin:
		return float32(medInterest)
	case balance >= ballerMin:
		return float32(ballerInterest)
	default:
		return 0
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	switch {
	case balance < 0:
		return balance * negInterest / 100
	case balance >= 0 && balance < medMin:
		return balance * smallInterest / 100
	case balance >= medMin && balance < ballerMin:
		return balance * medInterest / 100
	case balance >= ballerMin:
		return balance * ballerInterest / 100
	default:
		return 0
	}
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for {
		if balance >= targetBalance {
			break
		}
		years++
		balance = AnnualBalanceUpdate(balance)
	}
	return years
}
