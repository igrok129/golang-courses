// Package secondblock - Functions, structures, pointers and more.
package secondblock

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"golang-courses/pkg/util"
)

// 2.1 Functions.

// Exercises.

// Exercise211 Write a function that will take the string text and output it.
//
// Sample input: Hello!
// Sample output: Hello!
func Exercise211(text string) {
	fmt.Print(text)
}

// Exercise212 Write a function that finds the smallest of the four numbers entered the same function.
//
// Input data: Entered four numbers.
// Output data: It's necessary to return from the function the smallest of the 4 given numbers.
//
// Sample input: 4 5 6 7
// Sample output: 4
func Exercise212() int {
	var a, min int

	for i := 0; i < 3; i++ {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return 0
		}

		if i == 0 {
			min = a
			continue
		}

		if min > a {
			min = a
		}
	}

	return min
}

// Exercise213 Write a "voting function" that returns the value (0 or 1) that occurs most often among the values X, Y, Z.
//
// Input data: 3 numbers are entered - X, Y, Z.
// Output data: It's necessary to return the value of the function from X, Y, Z.
//
// Sample input: 0 0 1
// Sample output: 0
func Exercise213(x, y, z int) int {
	if x+y+z >= 2 {
		return 1
	}

	return 0
}

// Exercise214 Write a function, that return N-th number from Fibonacci sequence.
//
// Input data: Entered a number N.
// Output data: Return the answer to the exercise.
//
// Sample input: 4
// Sample output: 3
func Exercise214(n int) int {
	return util.GetFib(100000)[n]
}

// Exercise215 Write a function, that takes a variable number of int arguments and returns the number of arguments received
// by the function and their sum.
func Exercise215(numbers ...int) (int, int) {
	var sum int
	for _, v := range numbers {
		sum += v
	}

	return len(numbers), sum
}

// 2.2 Packages.
// 2.3 Pointers.

// Exercises.

// Exercise231 Write a function, that multiplies the values referenced by two pointer and prints the resulting product
// to the console.
//
// Sample input: 2 2
// Sample output: 4
func Exercise231(x1, x2 *int) {
	fmt.Print(*x1 * *x2)
}

// Exercise232 Swap the values of the variables referenced by the pointers. After that, the variables need to be displayed.
//
// Sample input: 2 4
// Sample output: 4 2
func Exercise232(x1, x2 *int) {
	a := *x1
	b := *x2

	*x1 = b
	*x2 = a

	fmt.Print(*x1, *x2)
}

// 2.4 Structures.

// Exercises.

// Exercise241 You need to implement a structure with properties-fields On, Ammo and Power, with types bool, int, int
// respectively. This structure must have methods: Shoot and RideBike, which take no arguments, but return a bool value.
// If On == false, then both methods will return false.
// You can shoot only if there is Ammo (then Ammo is reduced by one, and the method returns true), if it's not there,
// the method will return false. The RideBike method works the same, but only depends on the Power property.
func Exercise241() *TestStruct {
	return new(TestStruct)
}

type TestStruct struct {
	On          bool
	Ammo, Power int
}

func (t *TestStruct) Shoot() bool {
	if t.On == false {
		return false
	}

	if t.Ammo > 0 {
		t.Ammo--
		return true
	}

	return false
}

func (t *TestStruct) RideBike() bool {
	if t.On == false {
		return false
	}

	if t.Power > 0 {
		t.Power--
		return true
	}

	return false
}

// 2.5 Strings.

// Exercises.

// Exercise251 The input is a string. It needs to be determined whether it's correct or not. A valid string starts with
// a capital letter and ends with a dot. If the string is correct, print "Right", else print "Wrong".
//
// Sample input: Be or not to be.
// Sample output: Right
func Exercise251() {
	str := util.ReadString()

	if !strings.HasSuffix(str, ".") {
		fmt.Print("Wrong")
		return
	}

	for i, v := range str {
		if i == 0 && !unicode.IsUpper(v) {
			fmt.Print("Wrong")
			return
		}
	}

	fmt.Print("Right")
}

// Exercise252 The input is a string. We need to determine if it's a palindrome. If the string is a palindrome, output
// "Palindrome", otherwise, output "No".
//
// Sample input: топот
// Sample output: Palindrome
func Exercise252() {
	var reverse []rune

	direct := []rune(util.ReadString())

	for i := len(direct) - 1; i >= 0; i-- {
		reverse = append(reverse, direct[i])
	}

	if string(reverse) == string(direct) {
		fmt.Print("Palindrome")
		return
	}

	fmt.Print("No")
}

// Exercise253 Given two strings X and S. You need to find and print the first occurrence of the substring S in the
// string X. If the substring S isn't in the string X, print -1.
//
// Sample input: awesome es
// Sample output: 2
func Exercise253() {
	var x, s string

	if _, err := fmt.Scan(&x, &s); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(strings.Index(x, s))
}

// Exercise254 A string is given as input, you need to make another string from it, leaving only odd characters.
//
// Sample input: ihgewlqlkot
// Sample output: hello
func Exercise254() {
	var input, res string

	if _, err := fmt.Scan(&input); err != nil {
		fmt.Print(err)
		return
	}

	for i, v := range input {
		if i%2 != 0 {
			res += string(v)
		}
	}

	fmt.Print(res)
}

// Exercise255 A string is given. You need to remove all characters that occur more that once and output the result.
//
// Sample input: zaabcbd
// Sample output: zcd
func Exercise255() {
	var str string
	var res []rune

	if _, err := fmt.Scan(&str); err != nil {
		fmt.Print(err)
		return
	}

	strRune := []rune(str)

	for _, v := range strRune {
		count := 0

		for i := 0; i < len(strRune); i++ {
			if v == strRune[i] {
				count++
			}

			if count == 1 && i == len(strRune)-1 {
				res = append(res, v)
			}
		}
	}

	fmt.Print(string(res))
}

// Exercise256 Your task is to check whether the password entered by the user matches the specified requirements. The
// password must be at least 5 characters long and contain only numbers and letters of the Latin alphabet. The input is a
// password string. If the password matches the requirements - print "OK", otherwise print "Wrong password".
//
// Sample input: fdsghdfgjsdDD1
// Sample output: OK
func Exercise256() {
	var str string

	if _, err := fmt.Scan(&str); err != nil {
		fmt.Print(err)
		return
	}

	strRune := []rune(str)

	for _, v := range strRune {
		if !unicode.In(v, unicode.Latin, unicode.Number) {
			fmt.Print("Wrong password")
			return
		}
	}

	if len(strRune) < 5 {
		fmt.Print("Wrong password")
		return
	}

	fmt.Print("OK")
}

// 2.6 Errors handling.

// Exercises.

// Exercise261 You have to call the divide function which divides two numbers, but it returns not only the result of the
// division, but also information about the error, In case of any error, you should print "error", if there is no error -
// the result of the function.
//
// Sample input: 10 5
// Sample output: 2
func Exercise261() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	if res, err := util.Divide(a, b); err != nil {
		fmt.Print("error")
		return
	} else {
		fmt.Print(res)
	}
}

// 2.7 Tasks resolving.

// Task271 The input is A and B - the legs of a right-angled triangle. Find the length of the hypotenuse.
//
// Sample input: 6 8
// Sample output: 10
func Task271() {
	var a, b float64

	if _, err := fmt.Scan(&a, &b); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(math.Sqrt(a*a + b*b))
}

// Task272 Given a string containing only English letters (uppercase and lowercase). Add "*" (asterisk) character between
// letters (you don't need to add "*" character before the first letter and after the last one).
//
// Input data: A string of non-zero length is entered. It's also known that the string length doesn't exceed 1000 characters.
// Output data: Output the string that will be obtained after adding the characters "*".
//
// Sample input: LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
// Sample output: L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
func Task272() {
	var str, res string

	if _, err := fmt.Scan(&str); err != nil {
		fmt.Print(err)
		return
	}

	for _, v := range str {
		res += fmt.Sprintf("*%v", string(v))
	}

	fmt.Print(strings.TrimLeft(res, "*"))
}

// Task273 Given a string containing only decimal digits. Find and display the largest number.
//
// Input data: A string of non-zero length is entered. It's also known that the string length doesn't exceed 1000 characters
// and the string contains only decimal digits.
// Output data: Print the biggest digit that occurs in the input string.
//
// Sample input: 1112221112
// Sample output: 2
func Task273() {
	var str string
	var max rune

	if _, err := fmt.Scan(&str); err != nil {
		fmt.Print(err)
		return
	}

	for i, v := range str {
		if i == 0 {
			max = v
			continue
		}

		if v > max {
			max = v
		}
	}

	fmt.Print(string(max))
}

// Task274 The input is an integer. It's necessary to square each digit of the number and output the resulting number.
//
// Sample input: 9119
// Sample output: 811181
func Task274() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Print(err)
		return
	}

	for i := util.GetNumLength(a) - 1; i >= 0; i-- {
		n := util.GetSpecOrder(a, i)
		fmt.Print(n * n)
	}
}
