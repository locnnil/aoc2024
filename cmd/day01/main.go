package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/locnnil/aoc2024.git/pkg/request"
	"github.com/locnnil/aoc2024.git/pkg/sorting"
)

func main() {
	in, err := request.ReadInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	nums := make([]string, 0, 2000)
	fmt.Printf("Words:\n %v", nums)

	nums = strings.Fields(in)

	fmt.Println("Number of words:\n", len(nums))
	fmt.Println("Counting words:\n", len(nums))

	hf := len(nums) / 2
	fmt.Printf("\nHalf: %v\n", hf)
	rigth := make([]int, 0, hf)
	left := make([]int, 0, hf)

	// done := make([]chan bool, 2)
	done := []chan bool{make(chan bool), make(chan bool)}

	// go split(done[0], &nums, &left, func(i int) bool { return i%2 == 0 })
	// go split(done[1], &nums, &rigth, func(i int) bool { return i%2 != 0 })

	// for i, n := range [][]int{left, rigth} {
	// 	go Spliter(done[i], &nums, &n, i)
	// }
	// <-done[0]
	// <-done[1]

	go Spliter(done[0], &nums, &left, 0)
	go Spliter(done[1], &nums, &rigth, 1)
	<-done[0]
	<-done[1]

	// fmt.Println("Number of words:\n", len(nums))
	// fmt.Println("Words: ", in)
	fmt.Println("Right:\nsize->", len(rigth), "\n itens->\n", rigth)
	fmt.Println("Left:\nsize->", len(left), "\n itens->\n", left)

	sorting.QuickSort(left)
	sorting.QuickSort(rigth)

	fmt.Println("Sorted Left:\nsize->", len(left), "\n itens->\n", left)
	fmt.Println("Sorted Right:\nsize->", len(rigth), "\n itens->\n", rigth)

	if len(left) != len(rigth) {
		panic("The number of elements in the left and right slices are different")
	}
	size := len(left)

	sum := 0
	for i := 0; i < size; i++ {
		// fmt.Printf("sortedLeft[%v]: %v\n", i, left[i])
		// fmt.Printf("sortedRigth[%v]: %v\n", i, rigth[i])
		sum += absDiff(left[i], rigth[i])
	}

	fmt.Println("\n\n[FINAL] Sum: ", sum)
}

// Avoiding the use of the math.Abs function
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Spliter(ok chan bool, nums *[]string, parity *[]int, init int) {
	fmt.Println("\nSize of nums: ", len(*nums))

	for i := init; i < len(*nums); i += 2 {
		num, err := strconv.Atoi((*nums)[i])
		if err != nil {
			panic(err)
		}
		*parity = append(*parity, num)
	}
	ok <- true
}

type ParityChecker func(int) bool

func Split(ok chan bool, nums *[]string, parity *[]int, check ParityChecker) {
	for i, n := range *nums {
		if check(i) {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			*parity = append(*parity, num)
		}
	}
	ok <- true
}
