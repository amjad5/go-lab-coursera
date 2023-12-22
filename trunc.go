package main

import (
	"fmt"
	// "strconv"
)

func main(){
	var floatVal float64
	fmt.Printf("Enter a number: ")
	_, err := fmt.Scanf("%f\n", &floatVal)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", int(floatVal))


}