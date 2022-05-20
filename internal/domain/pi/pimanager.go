package pi

// CalculatePi calculate n'th term of PI using Leibniz function
func CalculatePi(terms int) float64 {
	numerator := 4.0
	denominator := 1.0
	operation := 1.0
	pi := 0.0

	for i := 0; i < terms; i++ {
		pi += operation * (numerator / denominator)
		denominator += 2.0
		operation *= -1
	}

	return pi
}
