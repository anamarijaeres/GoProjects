package numOperations

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)

}
func FibonacciEntry() {

	var n = 5

	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, fibonacci(i))
	}
	fmt.Println(result)

	result2 := []int{1, 1}

	for i := 0; i < n-2; i++ {
		newValue := result2[len(result2)-1] + result2[len(result2)-2]
		result2 = append(result2, newValue)

	}

	fmt.Println(result2)

}
