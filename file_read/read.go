package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct{
	fname string
	lname string
}

/*
	Write a program which reads information from a file and represents it in a slice of structs. 
	Assume that there is a text file which contains a series of names. 
	Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

	Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. 
	Each field will be a string of size 20 (characters).

	Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file 
	and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and 
	after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After 
	reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

func main(){
	fmt.Println("starting......")
	var fileName string

	fmt.Println("Please enter a file name along with extension i.e. names.txt: ")
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil{
		panic(err)
	}

	fileObj, fileError := os.Open(fileName)
	if fileError != nil{
		fmt.Println("File not found in current directory!")
		fmt.Println("Exiting.......!")
		return 
	}

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(fileObj)
 
    fileScanner.Split(bufio.ScanLines)

	personSlice := make([]Name, 0)

	fmt.Println("Reading each line and storing into slice of structs.....")
    for fileScanner.Scan() {
		fullName := strings.Split(fileScanner.Text(), " ")
		personSlice = append(personSlice, Name{fname: fullName[0], lname: fullName[1]})
    }

	fmt.Println("Looping through slice of structs......")
	for _, ele := range personSlice {
		fmt.Println("firstName: " + ele.fname + " lastName: " + ele.lname)
	}
	fmt.Println("Reading Done......")
    fileObj.Close()
	fmt.Println("File gracefully closed......")

}
