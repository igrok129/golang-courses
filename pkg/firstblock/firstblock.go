// Package firstblock - The basics.
package firstblock

import (
	"fmt"
	"math"

	"golang-courses/pkg/util"
)

// 1.1 About a course.
// 1.2 About go.
// 1.3 First program.
// 1.4 Data types.
// 1.5 Variables and arithmetic operations, data input/output.

// Exercises.

// Exercise151 Write a program that sequentially performs the following operations with the input number:
// 1. The number is multiplied by 2.
// 2. Then 100 is added to the number.
//
// Sample input: 1
// Sample output: 102
func Exercise151() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(100 + (a * 2))
}

// Exercise152 Petya was in a hurry to go to school and incorrectly wrote a program that first finds the squares
// of two numbers and the sums them up. Fix his program.
//
// Sample input:
// 2
// 2
// Sample output: 8
//
/*func main(){

	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * 2
	c = a + b
	fmt.Println(c)
}*/
func Exercise152() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	a *= a
	b *= b
	c := a + b

	fmt.Print(c)
}

// Exercise153 Given a whole number, find its square.
//
// Input data format: The input is a single integer.
// Output data format: The program should output the square of the given number.
//
// Sample input: 3
// Sample output: 9
func Exercise153() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(a * a)
}

// Exercise154 Given a natural number, print its last digit.
//
// Input data format: The input is a natural number N, not exceeding 10000.
// Output data format: Print one integer - the answer to the exercise.
//
// Sample input: 123
// Sample output: 3
func Exercise154() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(a % 10)
}

// Exercise155 Given a non-negative integer. Find the number of tens (that is, the second digit from the right).
//
// Input data format: The input is a natural number not exceeding 10000.
// Output data format: Print one integer - the number of tens.
//
// Sample input: 2010
// Sample output: 1
func Exercise155() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(util.GetSpecOrder(a, 1))
}

// Exercise156 The hour hand has turned D degrees since the beginning of the day. Determine how many whole
// hours H and whole minutes M are now.
//
// Input data: The input to the program is an integer D (0 < D < 360).
// Output data: Display the phrase: It is ... hours ... minutes.
//
// Sample input: 90
// Sample output: It is 3 hours 0 minutes.
func Exercise156() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	h := a / 30
	m := (a % 30) * 2

	fmt.Printf("It is %d hours %d minutes.", h, int(m))
}

// 1.6 Comments.

// Exercises.

// Exercise161 Remove unnecessary comments so that the program displays the number 100.
//
// Sample input:
// Sample output: 100
//
//func main(){
//	// a:=44
//	/*
//	   var a2 int = 10
//	*/
//	a2 = a2 * 10
//	fmt.Println(a2)
//}
func Exercise161() {
	var a2 = 10

	a2 = a2 * 10

	fmt.Println(a2)
}

// 1.7 Constants.

// Exercises.

// Exercise171 Fix the error in the program below:
//
// Sample input:
// Sample output: 18
//
/*func main(){
	var a int = 8
	const b int = 10
	a = a + b
	b = b + a
	fmt.Println(a)
}*/
func Exercise171() {
	var a = 8
	const b = 10

	a = a + b

	fmt.Println(a)
}

// 1.8 Conditional expressions.
// 1.9 Conditional constructions.

// Exercises.

// Exercise191 The input is an integer. If the number is positive - display the message "Positive number",
// if the number is negative - "Negative number". If zero is supplied, display the message "Zero".
//
// Sample input: 5
// Sample output: Positive number
func Exercise191() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	switch {
	case a < 0:
		fmt.Print("Negative number")
	case a > 0:
		fmt.Print("Positive number")
	default:
		fmt.Print("Zero")
	}
}

// Exercise192 Given a three-digit number, determine if all of its digits are different.
//
// Input data format: The input is a natural three-digit number.
// Output data format: Print "YES" if all digits of the number are different, otherwise - "NO".
//
// Sample input 1: 237
// Sample input 2: 117
// Sample output 1: YES
// Sample output 2: NO
func Exercise192() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	first := util.GetSpecOrder(a, 2)
	second := util.GetSpecOrder(a, 1)
	third := util.GetSpecOrder(a, 0)

	switch {
	case first == second || first == third || second == third:
		fmt.Print("NO")
	default:
		fmt.Print("YES")
	}
}

// Exercise193 Given a non-negative integer. Find and print the first digit of the number.
//
// Input data format: The input is a natural number not exceeding 10000.
// Output data format: Print one integer - the first digit of the given number.
//
// Sample input: 12345
// Sample output: 1
func Exercise193() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(util.GetSpecOrder(a, util.GetNumLength(a)-1))
}

// Exercise194 Determine if the ticket is lucky. A ticket is considered lucky, in the six-digit number of which,
// the sum of the first three digits coincides with the sum of the last three.
//
// Input data format: The ticket number is given at the entrance - a six-digit number.
// Output data format: Print "YES" if the ticket is lucky, otherwise - "NO".
//
// Sample input: 613244
// Sample output: YES
func Exercise194() {
	var a, sumFirst, sumLast int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := util.GetNumLength(a) - 1; i >= 0; i-- {
		if i > 2 {
			sumFirst += util.GetSpecOrder(a, i)
			continue
		}

		sumLast += util.GetSpecOrder(a, i)
	}

	if sumFirst == sumLast {
		fmt.Print("YES")
		return
	}

	fmt.Print("NO")
}

// Exercise195 Determine if a given year is a leap year:
// - is a multiple of 400.
// - is a multiple of 4, but not a multiple of 100.
//
// Input data: A single number is entered - the number of the year (integer, positive, does not exceed 10000).
// Output data: It's required to display the word "YES" if the year is a leap and "NO" otherwise.
//
// Sample input: 2000
// Sample output: YES
func Exercise195() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	switch {
	case (a%400 == 0) || (a%4 == 0 && a%100 != 0):
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}

// 1.10 Cycles.

// Exercises.

// Exercise1101 Write a program that prints the squares of natural numbers from 1 to 10. The square of each number
// must be printed on a new line.
//
// Sample input:
// Sample output:
// 1
// 4
// 9
// 16
// 25
// 36
// 49
// 64
// 81
// 100
func Exercise1101() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

// Exercise1102 It's required to write a program, during which two natural numbers A and B are read from the
// keyboard (each isn't more than 100, A < B). Print the sum of all numbers from A to B inclusively.
//
// Sample input: 1 5
// Sample output: 15
func Exercise1102() {
	var a, b, sum int

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	for ; a <= b; a++ {
		sum += a
	}

	fmt.Print(sum)
}

// Exercise1103 Write a program that finds the sum of two-digit numbers divisible by 8 in a sequence of numbers.
// The program in the first line receives the number N - the number of numbers in the sequence, in the second line - N
// numbers included in this sequence.
//
// Sample input:
// 5
// 38 24 800 8 16
// Sample output: 40
func Exercise1103() {
	var a, b, sum int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		if (util.GetNumLength(b) == 2) && b%8 == 0 {
			sum += b
		}
	}

	fmt.Print(sum)
}

// Exercise1104 The sequence consists of natural numbers and ends with the number 0. Determine the number of elements
// of this sequence that are equal to its largest element.
//
// Input data format: A non-empty sequence of natural numbers is introduced, ending with the number 0.
// Input data output: Output the answer to the exercise.
//
// Sample input: 1 3 3 1 0
// Sample output: 2
func Exercise1104() {
	var a, biggest, count int

	for {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return
		}

		if a == 0 {
			break
		}

		if biggest < a {
			biggest = a
			count = 0
		}

		if biggest == a {
			count++
		}
	}

	fmt.Print(count)
}

// Exercise1105 Find the first number between 1 and N inclusively, that is a multiple of C but not multiple of D.
//
// Input data: Entered 3 natural numbers N, C, D, each of which does not exceed 10000.
// Output data: Output the answer to the exercise.
//
// Sample input: 20 3 5
// Sample output: 3
func Exercise1105() {
	var n, c, d int

	if _, err := fmt.Scan(&n, &c, &d); err != nil {
		fmt.Print(err)
		return
	}

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
	}
}

// Exercise1106 Write a program that reads integers from the console, one number per line. For each number entered check:
// - if the numbers is less than 10, then skip this number.
// - if the number is greater than 100, then stop reading the numbers.
// - otherwise print this number back to the console on a separate line.
//
// Sample input: 30 11 7 101
// Sample output:
// 30
// 11
func Exercise1106() {
	var a int

	for {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return
		}

		if a > 100 {
			return
		}

		if a < 10 {
			continue
		}

		fmt.Println(a)
	}
}

// Exercise1107 The deposit in the bank is X rubles. It increases annually by P percent, after which the fractional part
// is discarded. Every year the amount of the contribution becomes larger. Determine in how many years the contribution
// will be at least Y rubles.
//
// Input data: The program receives three natural numbers as input: X, P, Y.
// Output data: The program should output a single integer.
//
// Sample input: 100 10 200
// Sample output: 8
func Exercise1107() {
	var x, p, y, count int

	if _, err := fmt.Scan(&x, &p, &y); err != nil {
		fmt.Print(err)
		return
	}

	for x < y {
		x += (x * p) / 100
		count++
	}

	fmt.Print(count)
}

// Exercise1108 Two numbers are given. Determine the digits included in the record of both the first and second numbers.
//
// Input data: The program receives two numbers as input. It is guaranteed that digits in numbers do not repeat.
// Output data: Output the answer to the exercise.
//
// Sample input: 564 8954
// Sample output: 5 4
func Exercise1108() {
	var a, b int
	var result []interface{}

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	for lenA := util.GetNumLength(a) - 1; lenA >= 0; lenA-- {
		for lenB := util.GetNumLength(b) - 1; lenB >= 0; lenB-- {
			if specA := util.GetSpecOrder(a, lenA); specA == util.GetSpecOrder(b, lenB) {
				result = append(result, specA)
			}
		}
	}

	fmt.Print(result...)
}

// 1.11 Formatted output.

// Exercises.

// Exercise1111 The input is a float64 number. You need to output the converted number according to the rule:
// the number is squared, then the fractional part is cut off so that 4 decimal places remain. But if the number is equal
// to or less than zero, output: "number R is not suitable", where R is the original number with 2 digits after the
// decimal point and 2 in width. And if the number is greater than 10000 - display the original number in exponential
// notation.
//
// Sample input 1: -000012.2123
// Sample input 2: 1000001
// Sample input 3: 12.12345678
// Sample output 1: number -12.21 is not suitable
// Sample output 2: 1.000001e+06
// Sample output 3: 146.9782
func Exercise1111() {
	var a float64

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	if a <= 0 {
		fmt.Printf("number %2.2f is not suitable", a)
		return
	}

	if a > 10000 {
		fmt.Printf("%e", a)
		return
	}

	fmt.Printf("%.4f", a*a)
}

// 1.12 Arrays and slices.

// Exercises.

// Exercise1121 In the first step, 10 positive integers are given on standard input, which must be written in input order
// into an array of 10 elements. The type of the numbers included in the array must correspond to the smallest possible
// unsigned integer. The name of the array that you must create yourself - workArray (required condition).
// At the second stage, 3 more pairs of numbers are supplied to the standard input - the indices of the elements of
// this array that need to be swapped.
//
// Sample input: 99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
// Sample output: 99 137 151 187 117 71 20 287 93 67
func Exercise1121() {
	var a, b int
	var workArray [10]uint8
	var substitution [6]uint8

	for i := 0; i < len(workArray); i++ {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return
		}

		workArray[i] = uint8(a)
	}

	for i := 0; i < len(substitution); i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		substitution[i] = uint8(b)
	}

	for _, pVal := range [3]int{0, 2, 4} {
		first := workArray[substitution[pVal]]
		second := workArray[substitution[pVal+1]]

		workArray[substitution[pVal]] = second
		workArray[substitution[pVal+1]] = first
	}

	for _, v := range workArray {
		fmt.Printf("%d ", v)
	}
}

// Exercise1122 Write a program that takes as input the number N (N >= 4), and then N int-elements of slice. The output
// should be the 4th element of this slice.
//
// Sample input:
// 5
// 41 -231 24 49 6
// Sample output: 49
func Exercise1122() {
	var a, b int
	var c []int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		c = append(c, b)
	}

	fmt.Print(c[3])
}

// Exercise1123 The input is five integers, which are written to the array. However, this part of the program has already
// been written. You need to write a code snippet that can find and display the biggest number in this array.
//
// Sample input: 2 3 56 45 21
// Sample output: 56
func Exercise1123() {
	array := [5]int{}
	var a int

	for i := 0; i < 5; i++ {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return
		}

		array[i] = a
	}

	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}

	fmt.Print(max)
}

// Exercise1124 Given an array of integers. Write a program that will print the elements of an array whose indices are even.
//
// Input data: Firstly, the number N is given - number of elements in the array (1 <= N <= 100). Then N numbers are written
// with a space - the elements of the array.
//
// Sample input:
// 6
// 4 5 3 4 2 3
// Sample output: 4 3 2
func Exercise1124() {
	var a, b int
	var res []interface{}

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		if i%2 == 0 {
			res = append(res, b)
		}
	}

	fmt.Print(res...)
}

// Exercise1125 Given a sequence of integers. Write a program that counts the number of positive numbers among the elements
// of a sequence.
//
// Input data: Firstly, the number N is given - number of elements in the array (1 <= N <= 100). Then N numbers are written
// with a space - the elements of the array.
// Output data: It's necessary to output a single number - the number of positive elements in the sequence.
//
// Sample input:
// 5
// 1 2 3 -1 -4
// Sample output: 3
func Exercise1125() {
	var a, b, res int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		if b > 0 {
			res++
		}
	}

	fmt.Print(res)
}

// 1.13 Tasks resolving.

// Task1131 Given a three-digit number. Find the sum of its digits.
//
// Input data format: The input is a three-digit number.
// Output data format: Print a single integer - the sum of the digits of the entered number.
//
// Sample input: 745
// Sample output: 16
func Task1131() {
	var a, b int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < 3; i++ {
		b += util.GetSpecOrder(a, i)
	}

	fmt.Print(b)
}

// Task1132 Given a three-digit number. Flip it over and then take it out.
//
// Input data format: The input is a three-digit number that does not end in zero.
// Output data format: Output the inverted number.
//
// Sample input: 843
// Sample output: 348
func Task1132() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < 3; i++ {
		fmt.Print(util.GetSpecOrder(a, i))
	}
}

// Task1133 it's the K-th second of the day. Determine how many whole hours H and whole minutes M have passed
// since the beginning of the day.
//
// Input data: The input to the program is an integer K (0 < K < 86399).
// Output data: Display the phrase: It is ... hours ... minutes.
//
// Sample input: 13257
// Sample output: It is 3 hours 40 minutes.
func Task1133() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	hours := a / 3600
	minutes := (a - (hours * 3600)) / 60

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}

// Task1134 Three numbers are given - A, B, C (A < B < C) - the lengths of the sides of the triangle.
// We need to check if the triangle is a right triangle. If it is, output "Rectangular". Else output "Non-rectangular".
//
// Sample input: 6 8 10
// Sample output: Rectangular
func Task1134() {
	var a, b, c int

	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		fmt.Print(err)
		return
	}

	if c*c == a*a+b*b {
		fmt.Print("Rectangular")
		return
	}

	fmt.Print("Non-rectangular")
}

// Task1135
//
// Input data: Three natural numbers are given. Determine if there is a triangle with such sides.
// Output data: If the triangle exists, print the string "Exists", otherwise print the string "Does not exist".
//
// Sample input: 4 5 6
// Sample output: Exist
func Task1135() {
	var a, b, c int

	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		fmt.Print(err)
		return
	}

	if a+b > c && a+c > b && b+c > a {
		fmt.Print("Exist")
		return
	}

	fmt.Print("Does not exist")
}

// Task1136 Two numbers are given. Find their arithmetic mean.
//
// Input data format: The input is two positive integers A and B.
// Output data format: The program should print the arithmetic mean of the numbers A and B (the answer can be an integer
// or a fractional number).
//
// Sample input 1: 3 5
// Sample input 2: 277 154
// Sample output 1: 4
// Sample output 2: 215.5
func Task1136() {
	var a, b float64

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print((a + b) / 2)
}

// Task1137 Given numbers determine the number of numbers that are equal to zero.
//
// Input data: A natural number N is entered, and then N numbers.
// Output data: Print the number of numbers that are equal to zero.
//
// Sample input: 5 1 8 100 0 12
// Sample output: 1
func Task1137() {
	var a, b, count int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		if b == 0 {
			count++
		}
	}

	fmt.Print(count)
}

// Task1138 Find the number of minimum elements in the sequence.
//
// Input data: A natural number N is entered, and then N integers of the sequence.
// Output data: Output the number of minimal elements of the sequence.
//
// Sample input: 3 21 11 4
// Sample output: 1
func Task1138() {
	var a, b, count, min int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < a; i++ {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Print(err)
			return
		}

		if i == 0 {
			min = b
		}

		if b < min {
			min = b
			count = 0
		}

		if b == min {
			count++
		}
	}

	fmt.Print(count)
}

// Task1139 Given a number, find its digital root.
//
// Input data: Enter one natural number N, not exceeding 10⁷.
// Output data: Output the digital root of the number N.
//
// Sample input: 3456
// Sample output: 9
func Task1139() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(digitalRoot(a, 0))
}

// digitalRoot count digital root of number.
func digitalRoot(number int, i int) int {
	if i == 3 {
		return number
	}

	var sum int

	for i := 0; i < util.GetNumLength(number); i++ {
		sum += util.GetSpecOrder(number, i)
	}

	return digitalRoot(sum, i+1)
}

// Task11310 Find the largest number on the segment from A to B that is a multiple of 7.
//
// Input data: Two integers A and B (A <= B) are entered.
// Output data: Find the largest number on the segment from A to B, divisible by 7, or print NO - if there are none.
//
// Sample input: 100 500
// Sample output: 497
func Task11310() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	for a <= b {
		if b%7 == 0 {
			fmt.Print(b)
			return
		}
		b--
	}

	fmt.Print("NO")
}

// Task11311 For the given number N, complete the phrase "There is grazing in the meadow ..." with one of the possible
// continuations: N korov, N korova, N korovy, conjugating the word korova(cow) correctly.
//
// Input data: Given a number N (0 < N < 100).
// Output data: The program should display the entered number N and one of the words (in Latin): korov, korova, korovy.
// for instance: 1 korova, 2 korovy, 5 korov.
//
// Sample input: 10
// Sample output: 10 korov
func Task11311() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%d %s", a, getRightCow(a))
}

// getRightCow conjugate word korova for a number.
func getRightCow(num int) string {
	korov := map[int]bool{0: true, 5: true, 6: true, 7: true, 8: true, 9: true}
	korovy := map[int]bool{2: true, 3: true, 4: true}

	if (num <= 20 && num >= 5) || korov[num%10] {
		return "korov"
	}

	if num%10 == 1 {
		return "korova"
	}

	if korovy[num] || korovy[num%10] {
		return "korovy"
	}

	return ""
}

// Task11312 Given a number N, print out all integer values of a power of two not exceeding N, in ascending order.
//
// Input data: Enter a natural number.
// Output data: Output the answer to the exercise.
//
// Sample input: 50
// Sample output: 1 2 4 8 16 32
func Task11312() {
	var a int
	var res []interface{}

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i <= a; i++ {
		if good := int(math.Pow(2, float64(i))); good <= a {
			res = append(res, good)
		} else {
			break
		}
	}

	fmt.Print(res...)
}

// Task11313 Given a natural number A > 1. Determine what kind of Fibonacci number it is. If A is not a Fibonacci number,
// print the number -1.
//
// Input data: Enter a natural number.
// Output data: Output the answer.
//
// Sample input: 8
// Sample output: 6
func Task11313() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for k, v := range util.GetFib(100000) {
		if v == a {
			fmt.Print(k)
			return
		}
	}

	fmt.Print(-1)
}

// Task11314 Given a natural number N. Print its binary representation.
//
// Input data: Given a single number N.
// Output data: It's necessary to derive the required representation of the number N.
//
// Sample input: 12
// Sample output: 1100
func Task11314() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%b", a)
}

// Task11315 Remove the given digit from a natural number.
//
// Input data: Enter a natural number and the digit to be removed.
// Output data: Output a number without the given digits.
//
// Sample input:
// 38012732
// 3
// Sample output: 801272
func Task11315() {
	var num, remove int

	if _, err := fmt.Scan(&num, &remove); err != nil {
		fmt.Print(err)
		return
	}

	for i := util.GetNumLength(num) - 1; i > 0; i-- {
		if fig := util.GetSpecOrder(num, i); fig == remove {
			continue
		} else {
			fmt.Print(fig)
		}
	}
}
