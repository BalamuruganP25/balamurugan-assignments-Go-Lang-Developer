package main

import (
	House "HouseRobber/src/Assignment"
	"fmt"
)

func main() {

	for {
		var user_input int

		fmt.Println("HouseRobber Assignment")
		fmt.Println("Press 1 for getting inputs from users")
		fmt.Print("please enter ->")
		fmt.Scanf("%d", &user_input)

		if user_input == 1 {

			var total_number int
			var manual_number int
			fmt.Print("Enter Number of elements to be Inserted -->")
			fmt.Scanf("%d", &total_number)
			arr := make([]int, 0)
			for i := 0; i < total_number; i++ {

				fmt.Scanf("%d", &manual_number)
				arr = append(arr, manual_number)

			}

			fmt.Println("Thanks for entering your manual inputs")

			House.Rob(arr)

		} else {

			fmt.Println("please give valid option")

		}

	}

}
