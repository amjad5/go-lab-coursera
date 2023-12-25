package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
	Write a program which prompts the user to first enter a name, and then enter an address. 
	Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. 
	Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
	*/

	var name,addr string 

	personMap:= make(map[string]string)

	fmt.Printf("Please enter your name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil{
		panic(err)
	}
	
	fmt.Printf("Please enter your address: ")
	_, err2 := fmt.Scanf("%s", &addr)
	if err2 != nil{
		panic(err2)
	}
	
	personMap["name"] = name
	personMap["address"] = addr

	jsonData, err3 := json.Marshal(personMap)
	if err3 != nil{
		panic(err3)
	}

	println(string(jsonData))



}

