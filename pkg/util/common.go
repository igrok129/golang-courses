// Package util - Helper functions.
package util

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

// GetFib get the fibonacci sequence with length len.
func GetFib(len int) []int {
	fib := []int{0, 1}

	for i := 2; i < len; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib
}

// ReadString read string from standard console input.
func ReadString() string {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	return strings.TrimRight(str, "\n")
}

// GetNumLength get length of num.
func GetNumLength(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

// GetSpecOrder get value of number in specific position (from right to left:)).
func GetSpecOrder(input int, pos int) int {
	return (input % int(math.Pow(10, float64(pos+1)))) / int(math.Pow(10, float64(pos)))
}

// Emulated functions that are called in the task itself.

// Divide function for Exercise261.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant divide by zero")
	}

	return a / b, nil
}

// Work function for Exercise311.
func Work(x int) int {
	if x > 3 {
		return x + 1
	}

	return x - 1
}

// ReadTask function for Exercise341.
func ReadTask() (interface{}, interface{}, interface{}) {
	return 42.12, 44.14, "/"
}

// Work2 function for Exercise391.
func Work2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("accomplished")
}
