// calculator with golang

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a number")
	reader := bufio.NewReader(os.Stdin)
	first_number, _ := reader.ReadString('\n')
	first_number = strings.TrimSpace(first_number)
	int_first_num, _ := strconv.Atoi(first_number)

	fmt.Println("Enter a sign (+, -, *, /)")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Enter a second number")
	second_number, _ := reader.ReadString('\n')
	second_number = strings.TrimSpace(second_number)
	int_second_num, _ := strconv.Atoi(second_number)

	switch input {
	case "+":
		fmt.Println("Your Output is :", int_first_num+int_second_num)
	case "-":
		fmt.Println("Your Output is :", int_first_num-int_second_num)
	case "*":
		fmt.Println("Your Output is :", int_first_num*int_second_num)
	case "/":
		if int_first_num > int_second_num && int_second_num != 0 {
			fmt.Println("Your Output is :", int_first_num/int_second_num)
		} else if int_second_num == 0 {
			fmt.Println("Zero division error")
		}else if(int_first_num  < int_second_num){
			fmt.Println("Second number should be less than first number")
		}

	default:
		fmt.Println("Enter a valid sign")

	}
}
