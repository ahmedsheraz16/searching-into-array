package main

import (
	"fmt"
)

func checkNumber(num string) int {

	arr := [10]string{"malik", "gujjar", "khan", "chaudary", "sindi", "balochi", "qureshi", "qazi", "awan", "Pashto"}
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
func main() {
	var toCheck string
	fmt.Println("Value to search = ", toCheck)
	fmt.Scan(&toCheck)

	if i := checkNumber(toCheck); checkNumber(toCheck) != -1 {
		fmt.Print("\nFound at index", i)
	} else {
		fmt.Print("not found")
	}
}
