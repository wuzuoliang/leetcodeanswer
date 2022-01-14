package main

import "fmt"

func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))

	fmt.Println(F(5, 1, 1))
}

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuvie(n-1)
	/**
	Rescuvie(5)
	{5 * Rescuvie(4)}
	{5 * {4 * Rescuvie(3)}}
	{5 * {4 * {3 * Rescuvie(2)}}}
	{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
	{5 * {4 * {3 * {2 * 1}}}}
	{5 * {4 * {3 * 2}}}
	{5 * {4 * 6}}
	{5 * 24}
	120
	*/
}

func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}

	return RescuvieTail(n-1, a*n)
	/**
	RescuvieTail(5, 1)
	RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
	RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
	RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
	RescuvieTail(1, 60*2)=RescuvieTail(1, 120)
	120
	*/
}

func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}

	return F(n-1, a2, a1+a2)
	/**
	F(5,1,1)
	F(4,1,1+1)=F(4,1,2)
	F(3,2,1+2)=F(3,2,3)
	F(2,3,2+3)=F(2,3,5)
	F(1,5,3+5)=F(1,5,8)
	F(0,8,5+8)=F(0,8,13)
	8
	*/
}
