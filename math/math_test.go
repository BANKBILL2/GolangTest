package math
import(
	"testing"
	"GolangClass/math"
)

func TestPrimeBetwnn1To10(t *testing.T) {
	given := 10

	prime(given)
}

func TestPowerBase2X4(t *testing.T) {
	givenB := 2
	givenX := 4
	want := 16

	get := math.Power(givenB, givenX)
	if want != get {
		t.Errorf("given b %d and x %d want %d but get %d\n", givenB, givenX, want, get)
	}
}