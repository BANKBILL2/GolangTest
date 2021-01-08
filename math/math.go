package math

// Power calculate power
func Power(base, exponent int) int{
	result := 1;
	for i:=1; i <= exponent; i++ {
		result *= base
	}
	return result
}