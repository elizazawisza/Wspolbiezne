package main

import "fmt"

func calmMode(showMagazine chan *Show , showToDoList chan *Show){
	var option string
	for{
		fmt.Println("The simulator is running in the calm mode.")
		fmt.Println( "You can get the information about the current state of simulation by typing:")
		fmt.Println("1 - information about the current magazine state")
		fmt.Println("2 - information about the current task to do")
		fmt.Println("3 - information about the employees")
		fmt.Println("4 - information about the machines")


		fmt.Scanln(&option)
		if option == "1"{
			fmt.Println()
			fmt.Println("Below you can see the current state of Magazine ")
			magazine := &Show{
				response: make (chan []string)}
			showMagazine <- magazine
			response := <-magazine.response
			fmt.Println(response)
			fmt.Println()
			fmt.Println()

		}else if option == "2"{
			fmt.Println()
			fmt.Println("Below you can see the current state of ToDoList ")
			list := &Show{
				response:make (chan []string)}
			showToDoList <-list
			response := <- list.response
			fmt.Println(response)
			fmt.Println()
			fmt.Println()


		}else if option == "3"{
			fmt.Println()
			fmt.Println("Below you can see the current statistics of employees ")

			for _, employer := range employersStatistics {
				fmt.Println("Employee Id: ", employer.id, "  type :", employer.employerType, "  successes :", employer.counter)
			}
			fmt.Println()
			fmt.Println()


		} else if option == "4"{
			fmt.Println()
			fmt.Println("Below you can see the current statistics of addition Machines ")
			for _, machine := range additionMachineStatistics {
				fmt.Println("Machine Id: ", machine.id, " amount of resolved Tasks :", machine.counter)
			}
			fmt.Println()
			fmt.Println("Below you can see the current statistics of multiplication Machines ")
			for _, machine := range multiplicationMachineStatistics {
				fmt.Println("Machine Id: ", machine.id,  "  amount of resolved Tasks :", machine.counter)
			}
			fmt.Println()


		}else{
			fmt.Println("Wrong data, please try again")
			continue
		}
	}
}