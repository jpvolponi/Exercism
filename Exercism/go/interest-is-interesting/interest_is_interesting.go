package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float32
	var blc float32 = float32(balance)
	switch {
	case blc < 0.0:
		rate = 3.213
	case blc >= 0.0 && blc < 1000.0:
		rate = 0.5
	case blc >= 1000.0 && blc < 5000.0:
		rate = 1.621
	default:
		rate = 2.475
	}
	return rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	panic("Please implement the YearsBeforeDesiredBalance function")
}
