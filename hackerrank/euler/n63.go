// https://www.hackerrank.com/contests/projecteuler/challenges/euler063
package main

import (
	"fmt"
	"math/big"
)

func main() {

	var n int

	fmt.Scanf("%d", &n)

	for i := 1; i < 10; i++ {
		var result big.Int
		result.Exp(big.NewInt(int64(i)), big.NewInt(int64(n)), nil)
		str := result.String()
		if len(str) == n {
			fmt.Println(str)
		}
	}

}
