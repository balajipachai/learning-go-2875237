package main

import (
	"bufio"
	"fmt"
	"io"
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

	// The new() function only allocates memory and does not initialize it
	m := make(map[string]int)

	m["balaji"] = 25

	fmt.Println(m)

	// Writing to a file
	content := "Balaji, is the wisest blockchain developer"
	file, err := os.Create("file.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Successfully written %v characters to the file\n", length)
	defer file.Close() // wait until all operation is completed
	defer readFile("file.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Data read from the file: \n", string(data))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
