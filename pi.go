package pi

import (
	"math/big"
)

// Digits computes the digits of PI as given amount
// excluding first two characters (3.)
func Digits(n int64) string {
	i := big.NewInt(1)

	var x big.Int
	x.Exp(big.NewInt(10), big.NewInt(n+20), nil)
	x.Mul(big.NewInt(3), &x)

	var pi big.Int
	pi.SetBytes(x.Bytes())

	for x.Cmp(big.NewInt(0)) == 1 {
		var y, a, d big.Int
		y.Mul(&x, i)
		d.Mul(big.NewInt(4), a.Add(i, big.NewInt(1)))
		x.Div(&y, &d)

		d.Add(i, big.NewInt(2))
		a.Div(&x, &d)
		pi.Add(&pi, &a)

		i.Add(i, big.NewInt(2))
	}

	var r big.Int
	r.Exp(big.NewInt(10), big.NewInt(20), nil)
	return pi.Div(&pi, &r).String()
}
