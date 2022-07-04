package calculateKey

import (
	"fmt"
	"math/big"
	"testing"
)

func TestPrimeModInverse(t *testing.T) {
	fmt.Println(PrimeModInverse(*big.NewInt(2), *big.NewInt(2)))
}

func TestLagrangeCoefficient(t *testing.T) {

	test1 := "69983908575440197882660754237092652335022943333787171199661134501211333823076200"

	tests := []string{"1", "1212121"}

	z := new(big.Int)
	z.SetString(q, 10)

	fmt.Println(LagrangeCoefficient(test1, tests, *z))
}

func TestCalculateKeys(t *testing.T) {
	test1 := "121212000000000000000000000121"

	tests := []string{"1", "1212121"}

	prvKey := CalculateKeys(test1, tests)

	fmt.Println(prvKey.String())
}
