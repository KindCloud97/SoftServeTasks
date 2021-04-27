package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Дано число m > 1.
// Получить наибольшее k,
// при котором 4^k < m.

// .\SoftServe_academy\Exercises\107\107.go

func Power(x, pow int) int {
	result := x
	for i := 1; i < pow; i++ {
		result = result * x
	}
	return result
}

func multIn() []int {
	var scanner = bufio.NewScanner(os.Stdin)
	var nums []int
	fmt.Println("Enter some numbers:")
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nums
		}
		nums = append(nums, value)
	}
	return nums
}

func MakeCalculations() {
	var n []int = multIn()
	for i := 0; i < len(n); i++ {
		fmt.Println("Наибольшее значение k =", Task107(n[i]))
	}
}

func Task107(m int) int {
	var k = 0
	for result := 0; result < m; {
		result = Power(4, k)
		k = k + 1
	}

	return k - 1
}

func main() {
	MakeCalculations()
}
