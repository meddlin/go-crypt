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

	fmt.Print("Seed entered: " + text)
	fmt.Println("Option: " + option)

	if string(option[0]) == "E" || string(option[0]) == "e" {
		encrypt(text)
	}
	if string(option[0]) == "D" || string(option[0]) == "d" {
		decrypt(text)
	}
}

// Key concepts : Iterate over a string with a basic loop
func encrypt(in_str string) {
	encodedString := ""

	// The (- 2) is a cheap way to bypass the \r\n tacked onto the end of the input
	for i := 0; i < (len(in_str) - 2); i++ {
		encodedVal := in_str[i] + 5
		encodedString += string(encodedVal)

		// Uncomment to see how each byte moves in the ROT-13 process
		//playByPlay(i, in_str, encodedVal)
	}

	fmt.Printf("%v\n", "Encrypted!")
	fmt.Println(encodedString)
}

// Decrypt function performs the "mathematical reverse" of the encrypt function
func decrypt(in_str string) {
	decodedString := ""

	for i := 0; i < (len(in_str) - 2); i++ {
		decodedVal := in_str[i] - 5			// Too large of an offset will break the current ASCII cipher
		decodedString += string(decodedVal)

		// Uncomment to see how each byte moves in the ROT-13 process
		//playByPlay(i, in_str, decodedVal)
	}

	fmt.Printf("%v\n", "Decrypted!")
	fmt.Println(decodedString)
}

// Displays the "diagnostics" of the cipher movement
func playByPlay(counter int, in_str string, val byte) {
	fmt.Printf("%v", counter)
	fmt.Printf("%v\n", ": " + string(in_str[counter]) + " => " + string(val))
}