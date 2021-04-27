package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Дано натуральное число n.
// Можно ли представить его в12
// виде суммы двух квадратов натуральных чисел?
// Если можно, то:
//  а)указать пару x, y таких натуральных чисел,
//  что n = x^2 + y^2;
//
//  б)указать пары x, y таких натуральных чисел,
//  что n = x^2 + y^2, x >= y.

// .\SoftServe_academy\Exercises\243\243.go

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

func SumOfSquares() {
	var n []int = multIn()
	var a, b []int
	for i := 0; i < len(n); i++ {
		a, b = Task243(n[i])
		if a[0] == -1 || b[0] == -1 {
			fmt.Printf("if n = %d\n\t No matches!\n", n[i])
		} else {
			fmt.Printf("if n = %d\n\tа)x and y = %d\tб)x and y = %d\n", n[i], a, b)
		}
	}
}

func Task243(n int) ([]int, []int) {
	var ok bool
	var a, b []int
	for x := n; x > 0; x-- {
		for y := n; y > 0; y-- {
			result := Power(x, 2) + Power(y, 2)
			if n == result {
				ok = true
				a = append(a, x, y)
				if x >= y {
					b = append(b, x, y)
				}
			}
		}
	}
	if !ok {
		a = append(a, -1)
		b = append(b, -1)
	}
	return a, b
}

func main() {
	SumOfSquares()
}
