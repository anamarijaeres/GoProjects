package main

import (
	"fmt"

	"tsi.co/mod/hello/numOperations"
	"tsi.co/mod/hello/school"
)

func calculate(nums []int, pipe chan<- int) {
	sum := 0
	for _, num := range nums {
		fmt.Println(num)
		sum += num
	}

	pipe <- sum
}

func main() {

	fmt.Println("The first task was:")
	numOperations.Numbers()
	fmt.Println("The second task was:")
	numOperations.FizzBuzz()
	fmt.Println("The third task was:")
	school.StudentGrades()
	fmt.Println("The fourth task was:")
	numOperations.FibonacciEntry()
	fmt.Println("The fifth task was:")
	numOperations.UserInput()
	pipe := make(chan int, 2)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Trying out concurrency.")

	go calculate(nums[0:len(nums)/2], pipe)
	go calculate(nums[len(nums)/2:], pipe)

	fmt.Println("Received 2 values")
	val1 := <-pipe
	val2 := <-pipe
	fmt.Println("Sum of the one half ", val1)
	fmt.Println("Sum of the other half ", val2)
	fmt.Println("Sum ", val1+val2)

}
