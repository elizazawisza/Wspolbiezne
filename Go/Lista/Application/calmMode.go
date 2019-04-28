package main

import "fmt"

func calmMode(showMagazine chan *Show , showToDoList chan *Show){
	var option string
	for{
		fmt.Println("The simulator is running in the calm mode.")
		fmt.Println( "You can get the information about the current state of simulation by typing:")
		fmt.Println("m - information about the current magazine state")
		fmt.Println("t - information about the current task to do")
		fmt.Println("e - information about the employees")


		fmt.Scanln(&option)
		if option == "m"{
			fmt.Println()
			fmt.Println("Below you can see the current state of Magazine ")
			magazine := &Show{
				response: make (chan []string)}
			showMagazine <- magazine
			response := <-magazine.response
			fmt.Println(response)
			fmt.Println()
			fmt.Println()

		}else if option == "t"{
			fmt.Println()
			fmt.Println("Below you can see the current state of ToDoList ")
			list := &Show{
				response:make (chan []string)}
			showToDoList <-list
			response := <- list.response
			fmt.Println(response)
			fmt.Println()
			fmt.Println()


		}else if option == "e"{
			fmt.Println()
			fmt.Println("Below you can see the current statistics of employees ")

			for _, employer := range employersStatistics {
				fmt.Println("Employee ID: ", employer.id, "  type :", employer.employerType, "  successes :", employer.counter)
			}
			fmt.Println()
			fmt.Println()


		} else{
			fmt.Println("Wrong data, please try again")
			continue
		}
	}
}