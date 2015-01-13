// https://www.hackerrank.com/contests/projecteuler/challenges/euler065
package main

import (
	"fmt"
	"math/big"
)

func constants(n big.Int) big.Int {
	remain := new(big.Int)
	remain.Mod(&n, big.NewInt(3))

	if remain.Cmp(big.NewInt(0)) == 0 {
		return *big.NewInt(1)
	} else if remain.Cmp(big.NewInt(1)) == 0 {
		rt := new(big.Int)
		rt.Div(&n, big.NewInt(3))
		rt.Add(rt, big.NewInt(1))
		rt.Mul(rt, big.NewInt(2))
		return *rt
	} else {
		return *big.NewInt(1)
	}
}

type fraction struct {
	n, d big.Int // numerator, denominator
}

func (a fraction) add(b fraction) fraction {
	// try to reduce

	n := new(big.Int)
	d := new(big.Int)

	n.Mul(&a.n, &b.d).Add(n, new(big.Int).Mul(&a.d, &b.n))
	d.Mul(&a.d, &b.d)

	return fraction{*n, *d}
}

func (f fraction) reciprocal() fraction {
	return fraction{n: f.d, d: f.n}
}

func cal(n int64) fraction {
	if n <= 1 {
		return fraction{*big.NewInt(2), *big.NewInt(1)}
	}

	last := fraction{constants(*big.NewInt(n - 2)), *big.NewInt(1)}

	for i := n - 3; i >= 0; i-- {
		c := constants(*big.NewInt(i))
		last = fraction{c, *big.NewInt(1)}.add(last.reciprocal())
	}

	return last.reciprocal().add(fraction{*big.NewInt(2), *big.NewInt(1)})
}

func sumOfDigits(n *big.Int) int {
	s := n.String()
	sum := 0
	for _, c := range s {
		sum += int(byte(c) - byte('0'))
	}

	return sum
}

func main() {
	var n int64
	fmt.Scanf("%d", &n)
	result := cal(n)

	fmt.Println(sumOfDigits(&result.n))
}
