package util

func MaxInt(x int, y int) int  {
	if x > y {
		return x
	}
	return y
}

func IsEven(num int) bool {
	if num % 2 == 0 {
		return true
	}
	return false
}

func IsOdd(num int) bool {
	return !IsEven(num)
}