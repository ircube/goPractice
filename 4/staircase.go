/*
Staircase detail

This is a staircase of size : n = 4

   #
  ##
 ###
####
Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

Write a program that prints a staircase of size .

Function Description

Complete the  function with the following parameter(s):

int n : an integer
Print

Print a staircase as described above. No value should be returned.
Note: The last line is not preceded by spaces. All lines are right-aligned.

Input Format

A single integer, , denoting the size of the staircase.

Constraints

 0 < n <= 100 .

Sample Input

6
Sample Output

     #
    ##
   ###
  ####
 #####
######
Explanation

The staircase is right-aligned, composed of # symbols and spaces, and has a height and width of n = 6.
*/

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
