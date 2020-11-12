package main

import "fmt"

/*
iota is a predefined identifier representing successive untyped
integer constants. It is zero indexed and can be used to declare
a set of related constants
*/

func main() {
	simpleIotaDemo()

	moreIotaDemo()

}

func simpleIotaDemo() {
	// note the weird looking syntax.
	const (
		a = iota
		b
		c
		d
		e
		f
		g
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d == 3)
	fmt.Println(e == 4)
	fmt.Println(f == 5)
	fmt.Println(g == 100) // finish on a false

	// note the reuse of the const keyword resets where iota
	// starts from.
	const (
		z = iota
		y
		x
	)

	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(x)

	// as we have two sets of values from using iota
	// we can express the comparisons
	fmt.Println(a == z) // true
	fmt.Println(b == y) // true
	fmt.Println(c == x) // true
}

func moreIotaDemo() {
	const (
		MON = iota
		TUE
		WED
		THU
		FRI
		SAT
		SUN
	)
	fmt.Println(MON, TUE, WED, THU, FRI, SAT, SUN)

	// note we used the blanks for the ignored days or closed days.
	const (
		Mon = iota
		Tue
		_
		Thu
		Fri
		Sat
		_
	)

	fmt.Println(Mon, Tue, Thu, Fri, Sat)

}
