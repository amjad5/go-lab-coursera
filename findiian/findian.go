package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	/*
		This program finds 'i' at the start of the string, 'n' at the end and it contains letter 'a'
		If a string doesn't match this criteria, program exits with a 'Not found!' message
	*/
	for{
		fmt.Printf("Enter a string: \n")
		reader := bufio.NewReader(os.Stdin)
		inputString, _ := reader.ReadString('\n')
		formattedStrLn := strings.TrimSpace(strings.ToLower(inputString))
		switch {
		case strings.HasPrefix(formattedStrLn, "i") && strings.HasSuffix(formattedStrLn, "n") && strings.Contains(formattedStrLn, "a"):
			fmt.Printf("Found!\n")
		default:
			fmt.Printf("Not Found!\n")
		}
	}
}