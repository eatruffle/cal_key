package calculateKey

import (
	"fmt"
	"math/big"
)

const q = "115792089237316195423570985008687907852837564279074904382605163141518161494337"

func PrimeModInverse(x big.Int, p big.Int) big.Int {
	t := x.Exp(&x, p.Sub(&p, big.NewInt(2)), &p)
	return *t
}

func LagrangeCoefficient(myId string, ids []string, field big.Int) *big.Int {
	coefficient := big.NewInt(1)
	for _, id := range ids {
		if id == myId {
			continue
		}

		n := new(big.Int)
		myIdB, ok := n.SetString(myId, 10)

		if !ok {
			fmt.Println("SetString: error")
		}

		m := new(big.Int)
		idB, ok := m.SetString(id, 10)

		if !ok {
			fmt.Println("SetString: error")
		}
		tmp := PrimeModInverse(*idB.Mod(idB.Sub(idB, myIdB), &field), field)

		tmp = *idB.Mod(idB.Mul(&tmp, idB), &field)

		coefficient.Mul(coefficient, &tmp)
	}
	return coefficient
}

func CalculateKeys(myId string, ids []string) big.Int {

	prvKey := big.NewInt(0)

	value := big.NewInt(12)

	z := new(big.Int)
	z.SetString(q, 10)

	f1 := prvKey.Add(prvKey, value.Mul(value, LagrangeCoefficient(myId, ids, *z)))

	prvKey = prvKey.Mod(f1, z)

	return *prvKey
}
