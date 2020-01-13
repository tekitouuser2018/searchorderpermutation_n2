package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanText() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func scanInt64() int64 {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}
	return int64(a)
}

func scanFloat() float64 {
	sc.Scan()
	a, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func scanTextSlice(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = scanText()
	}
	return a
}

func scanSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func scanFloatSlice(n int) []float64 {
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = scanFloat()
	}
	return a
}

func scanInt64Slice(n int64) []int64 {
	a := make([]int64, n)
	for i := int64(0); i < n; i++ {
		a[i] = scanInt64()
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	// length of a permutation
	n := scanInt()
	// permutation
	p := scanSlice(n)

	memo := make([]int, 0)
	for i := 1; i <= n; i++ {
		memo = append(memo, i)
	}

	m := n - 1
	partial := factorial(m)

	pNum := int64(1)

	//O(n)
	for i := 0; i < n-1; i++ {
		num := p[i]
		// O(log n)
		index := binarySearch(memo, num)
		// O(n)
		memo = deleteNumOfArray(memo, index)
		pNum += partial * int64(index)
		partial /= int64(m)
		m--

	}

	fmt.Println(pNum)

}

func binarySearch(arr []int, number int) int {

	l := -1
	r := len(arr)
	index := 0
	for {
		mid := l + (r-l)/2
		tmp := arr[mid]

		if r-l <= 1 || tmp == number {
			index = mid
			break
		} else if tmp < number {
			l = mid
		} else {
			r = mid
		}
	}
	return index
}

func factorial(n int) int64 {
	sum := int64(1)
	for i := 2; i <= n; i++ {
		sum *= int64(i)
	}
	return sum
}

func deleteNumOfArray(arr []int, position int) []int {
	left := arr[:position]
	r := arr[position+1:]
	appendArr := append(left, r...)
	return appendArr
}
