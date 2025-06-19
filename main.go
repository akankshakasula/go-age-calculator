// Age calculator
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	year := time.Now().Year()
	fmt.Println("Enter your year of birth")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	intInput, _ := strconv.Atoi(input)
	age := year - intInput
	fmt.Printf("you will be %v in %v", age, year)
}
