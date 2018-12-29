package main

import (
	"fmt"
	"bufio"
	"os"
)

// Application enty point
func main() {
	fmt.Printf("Welcome to the ASCII cipher!\n")

	// Setup to accept user input as parameters from terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Print("Encrypt/Decrypt (E / d)? : ")
	option, _ := reader.ReadString('\n')

	fmt.Println("Seed entered: " + text)
	fmt.Println("Option: " + option)

	encrypt(text)
}

// Key concepts : Iterate over a string with a basic loop
func encrypt(in_str string) {
	encodedString := ""

	// The (- 2) is a cheap way to bypass the \r\n tacked onto the end of the input
	for i := 0; i < (len(in_str) - 2); i++ {
		encodedVal := in_str[i] * 2

		fmt.Printf("%v", i)
		fmt.Printf("%v\n", ": " + string(in_str[i]) + " => " + string(encodedVal))

		encodedString += string(encodedVal)
	}

	fmt.Printf("%v\n", "Encoded!")
	fmt.Println(encodedString)
}