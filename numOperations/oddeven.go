package numOperations

import "fmt"

func Numbers() {
	var a int = 3
	var b int = 4
	if a < b {
		var res = float32(a) / float32(b)
		fmt.Printf("%.2f \n", res)
	} else if b < a {
		fmt.Println(a >> b)
	} else {
		if a%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	}
}
