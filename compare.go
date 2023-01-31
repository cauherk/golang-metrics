package metrics

func DeltaCompareFloat64(expected float64, actual float64, delta float64) bool {
	if expected > actual {
		return expected-actual <= delta
	}
	return actual-expected <= delta
}
