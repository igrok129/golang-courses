// Package thirdblock - Maps, Files, interfaces, multithreading and more.
package thirdblock

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"golang-courses/pkg/util"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

// 3.1 Maps.

// Exercises.

// Exercise311 The standard input is 10 space-separated integers (numbers can be repeated). You need ot calculate the
// result of the work function for each of the received numbers. The results of the calculations, separated by spaces,
// should be printed on a line. However, the work function is taking too long. The execution time allocated to you isn't
// enough to sequentially process each number, so you need ot implement caching of ready-made results and use them in
// your work.
//
// Sample input: 3 1 5 2 3 5 3 0 3 4
// Sample output: 2 0 6 1 2 6 2 -1 2 5
func Exercise311() {
	var a int
	res := make(map[int]int)

	for i := 0; i < 10; i++ {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Print(err)
			return
		}

		if v, ok := res[a]; ok {
			fmt.Printf("%d ", v)
		} else {
			res[a] = util.Work(a)
			fmt.Printf("%d ", res[a])
		}
	}
}

// 3.2 Type conversion.

// Exercises.

// Exercise321 Write a function that converts an input int64 to uint16.
func Exercise321(a int64) uint16 {
	return uint16(a)
}

// Exercise322 Write a function that reads two string variables and returns an int64 number that you need to get by adding
// these strings.
//
// Sample input: %^80 hhhhh20&&&&nd
// Sample output: 100
func Exercise322(a, b string) int64 {
	n1, err := strconv.Atoi(clearStr(a))
	if err != nil {
		fmt.Print(err)
		return 0
	}

	n2, err := strconv.Atoi(clearStr(b))
	if err != nil {
		fmt.Print(err)
		return 0
	}

	return int64(n1 + n2)
}

// clearStr remove any symbols from a string, but numbers.
func clearStr(a string) string {
	var res string

	for _, v := range a {
		if !unicode.Is(unicode.Number, v) {
			continue
		} else {
			res += string(v)
		}
	}

	return res
}

// Exercise323 In the spreadsheet editors we are used to, there is a convenient representation of a number with a separator
// of digits in the form of space, in addition, in Russia, the integer part is separated from the fractional by a comma.
// A set of such numbers was exported to CSV format, where the symbol ";" is used as a separator. On standard input you get
// 2 such real numbers. As a result you need to output the quotient of dividing the first number by the second with an
// accuracy of four decimal places.
//
// Sample input: 1 149,6088607594936;1 179,0666666666666
// Sample output: 0.9750
func Exercise323() {
	str := util.ReadString()
	str = strings.Replace(str, ",", ".", -1)
	str = strings.Replace(str, " ", "", -1)

	numbers := strings.Split(str, ";")

	first, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	second, err := strconv.ParseFloat(numbers[1], 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%.4f", first/second)
}

// 3.3 Anonymous functions.

// Exercises.

// Exercise331 You must declare a function of the form func(uint) uint, which can later be called by the name fn. The input
// function receives a positive integer (uint), return a number of the same type in which there are no odd digits and the
// digit 0. If the resulting number is 0, then the number 100 is returned. The numbers in the returned number have the same
// order as in the original number.
//
// Sample input: 727178
// Sample output: 28
func Exercise331(i uint) {
	fn := func(i uint) uint {
		var resStr string

		if i == 0 {
			return 100
		}

		numStr := strconv.FormatUint(uint64(i), 10)

		for _, v := range []rune(numStr) {
			if n, err := strconv.ParseInt(string(v), 10, 64); n == 0 || n%2 != 0 || err != nil {
				continue
			} else {
				resStr += string(v)
			}
		}

		resUint, err := strconv.ParseUint(resStr, 10, 64)
		if err != nil {
			fmt.Print(err)
			return 0
		}

		if output := uint(resUint); output != 0 {
			return output
		} else {
			return 100
		}
	}

	fmt.Print(fn(i))
}

// 3.4 Interfaces.

// Exercises.

// Exercise341 You get 3 values of type empty interface: if all goes well, then the first 2 values will be float64. The
// third value will be ideally be a string, which can have the values: "+", "-", "*", "/" (a certain mathematical operation).
// But such ideal cases won't always be, you can get values of other types, and also the string in the third value may not
// apply to one of the indicated mathematical operations.
// The output of the program should be like this:
// 1. In a normal situation, you should print the result of a mathematical operation to the standard output with an accuracy
// of 4 digits after the decimal point.
// 2. If the first or second value isn't a float64 type, you should print an error message like: value=get_value:value_type
// for example: value=true:bool
// 3. If the third value is of the wrong type, or if a sign is passed that is not related to the above mathematical operations,
// the error message should be: unknown operation.
// It's guaranteed that there can be only one error in the arguments, so if you saw that it contains an error when checking
// the first value, print an error message and exit the program, checking the remaining arguments no longer matters, and
// the checking system will accept 2 error messages as a violation of the terms of the assignment.
func Exercise341() {
	v1, v2, operand := util.ReadTask()

	if err := checkVal(v1, v2); err != nil {
		fmt.Print(err)
		return
	}

	if err := checkOperand(operand); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(doMath(v1.(float64), v2.(float64), operand.(string)))
}

// checkVal check whether v float64 or not.
func checkVal(v ...interface{}) error {
	for i := 0; i < len(v); i++ {
		switch v[i].(type) {
		case float64:
			continue
		default:
			return fmt.Errorf("value=%v: %T", v[i], v[i])
		}
	}

	return nil
}

// checkOperand check whether v is known operand or not.
func checkOperand(v interface{}) error {
	if fmt.Sprintf("%T", v) != "string" {
		return fmt.Errorf("unknown operation")
	}

	switch v.(string) {
	case "+":
	case "-":
	case "*":
	case "/":
		return nil
	default:
		return fmt.Errorf("unknown operation")
	}

	return nil
}

// doMath calculate result of operand two numbers.
func doMath(v1, v2 float64, operand string) string {
	var res float64

	switch operand {
	case "+":
		res = v1 + v2
	case "-":
		res = v1 - v2
	case "*":
		res = v1 * v2
	case "/":
		res = v1 / v2
	}

	return fmt.Sprintf("%.4f", res)
}

// Exercise342
// 1. You must declare a new type that satisfies the fmt.Stringer interface. Your type should allow it to print like this:
// [  XXX]: where the spaces are "empty" battery capacity and X is "charged".
// 2. On standard input, you get a string consisting of exactly 10 digits: 0 or 1 (the order 0/1 is random). Your task is
// to read this string in any possible way and create an object of the type you declared at the first stage based on this string.
//
// Sample input: 1000010011
// Sample output: [      XXXX]
func Exercise342() {
	var charge string

	if _, err := fmt.Scan(&charge); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(Battery{charge: charge})
}

type Battery struct {
	charge string
}

func (b Battery) String() string {
	var count int

	charge := make([]int, 10)
	res := "["

	for _, v := range []rune(b.charge) {
		if v == 49 {
			count++
		}
	}

	for i, _ := range charge {
		if i >= len(charge)-count {
			res += "X"
		} else {
			res += " "
		}
	}

	res += "]"

	return res
}

// 3.5 Working with files.

// Exercises.

// Exercise351 Integers in the range 0 - 100 are supplied to the standard input, each number is supplied to the standard
// input on a new line (the number of numbers isn't known). It's required to read all these numbers and print their sum to
// the standard output.
//
// Sample input:
// 33
// 47
// 12
// 79
// 15
// Sample output: 186
func Exercise351() {
	var res int

	r := bufio.NewWriter(os.Stdout)
	q := bufio.NewScanner(os.Stdin)

	for q.Scan() {
		n, err := strconv.Atoi(q.Text())
		if err != nil {
			fmt.Print(err)
			return
		}

		res += n
	}

	if _, err := r.WriteString(strconv.FormatInt(int64(res), 10)); err != nil {
		fmt.Print(err)
		return
	}

	if err := r.Flush(); err != nil {
		fmt.Print(err)
		return
	}
}

// Exercise352 - Finding a file in a given format and processing it.
// This task will help you understand the encoding/csv and path/filepath packages, although the archive/zip
// package can also be used to solve it (since the task file is provided in this format). The test archive which you can
// download form our repository (https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_csv_1), contains
// a set of folders and files. One of these files is a CSV data file, while the other files don't contain structured data.
// It's required to find and read this single file with structured data (this is a 10x10 table, the separator is a comma),
// and as answer, you must specify the number located on line 5 and position 4.
//
// Sample input:
// Sample output: 42
func Exercise352() {
	r, err := zip.OpenReader("taskdata/exercise352/task.zip")
	if err != nil {
		fmt.Print(err)
		return
	}

	defer func(r *zip.ReadCloser) {
		if err := r.Close(); err != nil {
			fmt.Print(err)
			return
		}
	}(r)

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		frc, err := f.Open()
		if err != nil {
			fmt.Print(err)
			return
		}

		crf := csv.NewReader(frc)
		rows, err := crf.ReadAll()

		if err != nil || len(rows) < 4 || len(rows[4]) < 2 {
			continue
		}

		fmt.Print(rows[4][2])

		return
	}
}

// Exercise353 - Step by step data search.
// The test file, which you can download from out repository (https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_sep_1)
// contains a long series of numbers separated by ";'. It's required to find at what position the number 0 is located and
// indicate it as an answer.
//
// Sample input:
// Sample output: 10527
func Exercise353() {
	fr, err := os.Open("taskdata/exercise353/task.data")
	if err != nil {
		fmt.Print(err)
		return
	}

	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '\n' {
				return 0, data[:i], bufio.ErrFinalToken
			}

			if data[i] == ';' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		return 0, data, bufio.ErrFinalToken
	}

	x := bufio.NewScanner(fr)
	x.Split(splitFunc)

	for i := 1; x.Scan(); i++ {
		if x.Text() == "0" {
			fmt.Print(i)
		}
	}
}

// 3.6 JSON.

// Exercises.

// Exercise361 The standard input receives data about the students of the university group in JSON format:
/*{
	"ID": 134,
	"Number": "Number-1274",
	"Year": 2,
	"Students": [
		{
			"LastName": "LastName",
			"FirstName": "FirstName",
			"MiddleName": "MiddleName",
			"Birthday": "4.04.1970",
			"Address": "632432, Moscow, str. Street 12-21",
			"Phone": "+7(948)709-47-24",
			"Rating": [1, 2, 3]
		},
		{
			//...
		}
	]
}*/
// The information about each student contains information about the grades he received (Rating). It's required to read
// the data and calculate the average number of grades received by the students of the group. The response to the task
// needs to be written to standard output in JSON format in the following form:
/*{
	"Average": 14.1
}*/
func Exercise361() {
	var sum float64

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
		return
	}

	sg := &StudentGroup{}

	if err := json.Unmarshal(input, sg); err != nil {
		fmt.Print(err)
		return
	}

	for _, student := range sg.Students {
		sum += float64(len(student.Rating))
	}

	res, err := json.MarshalIndent(struct {
		Average float64
	}{Average: sum / float64(len(sg.Students))}, "", "    ")
	if err != nil {
		fmt.Print(err)
		return
	}

	if _, err := os.Stdout.Write(res); err != nil {
		fmt.Print(err)
		return
	}
}

type StudentData struct {
	LastName, FirstName, MiddleName, Birthday, Address, Phone string
	Rating                                                    []int
}

type StudentGroup struct {
	ID, Year int
	Number   string
	Students []StudentData
}

// Exercise362 This task is focused on real work with data in JSON format. For the solution, we will use the OKVED reference
// book. The data structure you need is contained in the structure-file, and the actual data that you need to decode is
// contained in the data-file. The files are hosted in our repository (https://github.com/semyon-dev/stepik-go/tree/master/work_with_json).
// In order to show that you were able to decode the document, you need to write down as a response the sum of the global_id
// fields of all elements encoded in the file.
//
// Sample input:
// Sample output: 763804981288
func Exercise362() {
	var sum int
	var jsonStr []OKVEDData

	file, err := os.Open("taskdata/exercise362/data-20190514T0100.json")
	if err != nil {
		fmt.Print(err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Print(err)
		return
	}

	data := make([]byte, fileInfo.Size())
	if _, err := file.Read(data); err != nil {
		fmt.Print(err)
		return
	}

	if err := json.Unmarshal(data, &jsonStr); err != nil {
		fmt.Print(err)
		return
	}

	for _, v := range jsonStr {
		sum += v.GlobalId
	}

	fmt.Print(sum)
}

type OKVEDData struct {
	GlobalId int `json:"global_id"`
}

// 3.7 Working with date and time.

// Exercises.

// Exercise371 The standard input is a string representation of a date and time in the following format: 1986-04-16T05:20:00+06:00
// Your task is to convert this string to Time, and then output it in UnixDate format: Wed Apr 16 05:20:00 +0600 1986
//
// Sample input: 1986-04-16T05:20:00+06:00
// Sample output: Wed Apr 16 05:20:00 +0600 1986
func Exercise371() {
	var input string

	if _, err := fmt.Scan(&input); err != nil {
		fmt.Print(err)
		return
	}

	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		print(err)
		return
	}

	fmt.Print(t.Format(time.UnixDate))
}

// Exercise372 The standard input is a string representation of the date and time of a specific event in the following
// format: 2020-05-15 08:00:00
// If the time of the event is before lunch (13-00), then nothing needs to be changed, it's enough to output the date to
// the standard output in the same format.
// If the event should happen after lunch, you need to reschedule it to the same time the next day, and then print it to
// standard output in the same format.
//
// Sample input: 2020-05-15 08:00:00
// Sample output: 2020-05-15 08:00:00
func Exercise372() {
	format := "2006-01-02 15:04:05"
	t, err := time.Parse(format, util.ReadString())
	if err != nil {
		fmt.Print(err)
		return
	}

	y, m, d := t.Date()
	afternoon, err := time.Parse("2006-January-02 15:04:05", fmt.Sprintf("%d-%s-%d 13:00:00", y, m, d))
	if err != nil {
		fmt.Print(err)
		return
	}

	if t.Before(afternoon) {
		fmt.Print(t.Format(format))
		return
	}

	fmt.Print(t.Add(time.Hour * 24).Format(format))
}

// Exercise373 The standard input is a string representation of two dates separated by a comma (see the example for the
// data format). You need to convert the received data to the Time type, and then display the duration of the period between
// the smaller and larger dates.
//
// Sample input: 13.03.2018 14:00:15,12.03.2018 14:00:15
// Sample output: 24h0m0s
func Exercise373() {
	format := "02.01.2006 15:04:05"
	dates := strings.Split(util.ReadString(), ",")

	d1, err := time.Parse(format, dates[0])
	if err != nil {
		fmt.Print(err)
		return
	}

	d2, err := time.Parse(format, dates[1])
	if err != nil {
		fmt.Print(err)
		return
	}

	if d1.Before(d2) {
		fmt.Print((time.Since(d1) - time.Since(d2)).Round(time.Second))
		return
	}

	fmt.Print((time.Since(d2) - time.Since(d1)).Round(time.Second))
}

// Exercise374 The length of the period is given on standard input (the format is shown in the example). In addition, you
// are given a date in Unix-time format: 1589570165 as a constant of type int64. It's required to read the length of a period,
// convert it to the Duration type, and then output (in UnixDate format) the date and time that you get when you add the
// period to the standard date.
//
// Sample input: 12 мин. 13 сек.
// Sample output: Fri May 15 19:28:18 UTC 2020
const now = 1589570165

func Exercise374() {
	period := strings.Replace(util.ReadString(), "мин.", "m", 1)
	period = strings.Replace(period, "сек.", "s", 1)
	period = strings.ReplaceAll(period, " ", "")

	duration, err := time.ParseDuration(period)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(time.Unix(int64(duration.Round(time.Second).Seconds())+now, 0).UTC().Format(time.UnixDate))
}

// 3.8 Parallelism part1.

// Exercises.

// Exercise381 Write a function that takes a channel and a number N of type int. It's necessary to return the value N + 1
// to the channel.
func Exercise381(c chan int, n int) {
	c <- n + 1
}

// Exercise382 Write a function that takes a channel and a string. It's necessary to send the same string to the given
// channel 5 times, adding a space to it.
func Exercise382(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

// Exercise383 Write a pipeline element (function) that remembers the previous value and sends values to the next stage
// of the pipeline only if it differs from what came before. Your function must accept two channels - inputStream and
// outputStream, in the first you will receive strings, in the second you should send values without repeats. As a result
// the outputStream should contain values that are not repeated in a row.
func Exercise383(inputStream chan string, outputStream chan string) {
	var prev string

	for s := range inputStream {
		if prev == s {
			continue
		}
		prev = s

		outputStream <- s
	}

	close(outputStream)
}

// 3.9 Parallelism part2.

// Exercises.

// Exercise391 You need to call the work() function in a separate goroutine and wait for the results of its execution.
func Exercise391() {
	done := make(chan struct{})

	go func(d chan struct{}) {
		util.Work2()
		close(done)
	}(done)

	<-done
}

// Exercise392 You need to call the work() function 10 times in separate goroutines and wait for the results of the
// execution of the called functions.
func Exercise392() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			util.Work2()
			wg.Done()
		}(wg)
	}

	wg.Wait()
}

// Exercise393 The function takes 3 channels as arguments, and returns a channel of type <-chan int.
// - if the argument is received from the firstChannel, you must send the square of the argument to the output channel.
// - if the argument is received from the secondChannel, you must send the result of multiplying the argument by 3 to the
// output channel.
// - if the argument is received from the stopChannel, you just need to terminate the function.
// Your function will receive only one value in one of the channels - received the value, processed it, and completed work.
func Exercise393(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		select {
		case first := <-firstChan:
			output <- first * first
		case second := <-secondChan:
			output <- second * 3
		case <-stopChan:
			return
		default:
			return
		}
	}()

	return output
}

// Exercise394 This function takes two read-only channels as an argument, returns a read-only channel. Through the arguments
// channel, the function will receive a series of numbers, and through the done channel, a signal to complete the work.
// When the signal about the completion of work is received, the function must send the sum of the received numbers to the
// output channel.
func Exercise394(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		var sum int

		for {
			select {
			case argument := <-arguments:
				sum += argument
			case <-done:
				output <- sum
				return
			}
		}
	}()

	return output
}

// Exercise395 Need to write a function that does the following N times:
// - read one number from each of the two channels in1 and in2.
// - compute f(x1) + f(x2).
// - write the received value to out.
// The fn function can run for a long time, waiting for something or performing calculations.
//
// Input data format:
// - The number of iterations is passed through the N argument.
// - Integers are passed through the in1 and int2 pipe-arguments.
// - A function to handle numbers before addition is passed through the fn argument.
// Output data format: The channel for outputting the results is passed through the out argument.
func Exercise395(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	numbers := make([]int, n*2)
	results := make(chan Res, n*2)

	go func() {
		defer close(done1)
		for i := 0; i < n; i++ {
			numbers[i] = <-in1
		}
	}()

	go func() {
		defer close(done2)
		for i := 0; i < n; i++ {
			numbers[n+i] = <-in2
		}
	}()

	go func(fn func(int) int) {
		<-done1
		<-done2
		for i, v := range numbers {
			go func(fn func(int) int, i, v int) {
				results <- Res{v: fn(v), id: i}
			}(fn, i, v)
		}
	}(fn)

	go func() {
		output := make([]int, n*2)
		for i := 0; i < n*2; i++ {
			r := <-results
			output[r.id] = r.v
		}

		for i := 0; i < n; i++ {
			out <- output[i] + output[n+i]
		}
	}()
}

type Res struct {
	v, id int
}
