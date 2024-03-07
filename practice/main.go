package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your entered text is: ", input)

	fmt.Print("Enter a number: ")
	aNumber, _ := reader.ReadString('\n')
	// Convert the entered string to a number using strconv package
	aInt, err := strconv.ParseInt(strings.TrimSpace(aNumber), 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println("Entered number is : ", aInt)

}
