package school

import "fmt"

func StudentGrades() {
	var results = make(map[string][]int)
	results["Dora"] = []int{5, 5, 4, 3, 4}
	results["David"] = []int{2, 3, 2, 1, 4}
	results["Lena"] = []int{4, 4, 2, 3, 4}
	results["Mark"] = []int{5, 5, 5, 3, 4}
	results["Ela"] = []int{3, 4, 5, 5, 4}

	for student, grades := range results {
		var max = grades[0]
		for _, grade := range grades {
			if grade > max {
				max = grade
			}
		}
		fmt.Printf("%s's highest grade is %d \n", student, max)
	}
}
