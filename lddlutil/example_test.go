package lddlutil_test

import (
	"fmt"
	"math"

	"github.com/Legenddigital/lddld/lddlutil"
)

func ExampleAmount() {

	a := lddlutil.Amount(0)
	fmt.Println("Zero Atom:", a)

	a = lddlutil.Amount(1e8)
	fmt.Println("100,000,000 Atoms:", a)

	a = lddlutil.Amount(1e5)
	fmt.Println("100,000 Atoms:", a)
	// Output:
	// Zero Atom: 0 LDDL
	// 100,000,000 Atoms: 1 LDDL
	// 100,000 Atoms: 0.001 LDDL
}

func ExampleNewAmount() {
	amountOne, err := lddlutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := lddlutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := lddlutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := lddlutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 LDDL
	// 0.01234567 LDDL
	// 0 LDDL
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := lddlutil.Amount(44433322211100)

	fmt.Println("Atom to kCoin:", amount.Format(lddlutil.AmountKiloCoin))
	fmt.Println("Atom to Coin:", amount)
	fmt.Println("Atom to MilliCoin:", amount.Format(lddlutil.AmountMilliCoin))
	fmt.Println("Atom to MicroCoin:", amount.Format(lddlutil.AmountMicroCoin))
	fmt.Println("Atom to Atom:", amount.Format(lddlutil.AmountAtom))

	// Output:
	// Atom to kCoin: 444.333222111 kLDDL
	// Atom to Coin: 444333.222111 LDDL
	// Atom to MilliCoin: 444333222.111 mLDDL
	// Atom to MicroCoin: 444333222111 μLDDL
	// Atom to Atom: 44433322211100 Atom
}
