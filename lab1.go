//this is the programme which we are covering string, integer, float, array, slice, conditional statement example with slice


package main

import "fmt"

func main(){
	//string example
	name := "seshu"
	fmt.Printf("Hello, %s!\n", name)

	//integer example
	age := 15
	if age >= 18{
		fmt.Println("you are elgiable for movie.")
	} else {
		fmt.Println("you are not elgiable for movie.")
    }

	//float example
	fat := 3.33563
	fmt.Printf("this value of fat approximately %f.\n", fat)

	//the array example
	cashWithdraw := [5]int{50, 100, 200, 500, 2000}
	fmt.Println("The type of cash in atm:", cashWithdraw)

	//slice example 
	workingDays := []string{"monday", "tuesday", "wednesday", "thursday", "friday"}
	fmt.Println("this are the working days:", workingDays)

	//conditional statment example with slice
	 pin := []int{1, 3}
	if len(pin)>4{
		fmt.Println("you entered the correct password pin:")
		
	}else {
		fmt.Println("you entered the password pin not less then 4 numbers:")
	}
	
}
