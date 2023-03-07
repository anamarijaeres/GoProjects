package numOperations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeInput(sc *bufio.Scanner) (string, []string) {
	sc.Scan()
	listNums := sc.Text()
	items := strings.Split(listNums, ",")
	return listNums, items

}

func performOperation(n1 float64, n2 float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong :((((")
		}
	}()

	fmt.Print("Enter one of the mathematical operations you want to perform on the numbers(/,*,+,-):\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()

	switch operation {
	case "/":
		if n2 == 0 {
			panic("Panic")
		}
		fmt.Printf("%.3f", n1/n2)
	case "*":
		fmt.Printf("%.3f", n1*n2)
	case "+":
		fmt.Printf("%.3f", n1+n2)
	case "-":
		fmt.Printf("%.3f", n1-n2)
	default:
		fmt.Println("Given operation is invalid, please try again")
		performOperation(n1, n2)
	}
}

func UserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	defer func() {
		if r := recover(); r != nil {
			UserInput()
		}

	}()

	fmt.Print("\nEnter two decimal numbers, separated by coma:\n")
	listNums, items := takeInput(scanner)

	for len(items) != 2 {
		fmt.Printf("%s are not two decimal numbers, separated by comma please enter again: \n", listNums)
		listNums, items = takeInput(scanner)
	}
	items[0] = strings.TrimSpace(items[0])
	items[1] = strings.TrimSpace(items[1])
	n1, err1 := strconv.ParseFloat(items[0], 64)
	n2, err2 := strconv.ParseFloat(items[1], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("One of the numbers is not a decimal number,try again!")
		panic("panic")
	} else {
		performOperation(n1, n2)
	}

}
