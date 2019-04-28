package main

import (
	"fmt"
	"time"
)

func main() {

	var verboseMode =  false
	var additionMachines [AdditionMachinesAmount]chan *resolveTask
	var multiplicationMachines [MultiplicationMachinesAmount]chan *resolveTask
	var addToTaskList = make (chan *AddToDoList)
	var addToItemsList = make (chan *AddToMagazine)
	var takeFromToDoList = make (chan *TakeFromToDoList)
	var takeFromItemsList = make (chan *TakeFromMagazine)
	var magazinePrint = make (chan *Show)
	var listPrint = make (chan *Show)

	for{
		fmt.Println("Choose a simulator mode you'd like to work with ")
		fmt.Println("c - calm mode ")
		fmt.Println("v - verbose mode ")

		var simulationMode string
		fmt.Scanln(&simulationMode)
		if simulationMode == "c"{
			go calmMode(magazinePrint, listPrint)
			verboseMode = false
			break
		}else if simulationMode == "v"{
			verboseMode = true
			break
		}else{
			fmt.Println("Wrong data. Please try again")
		}

	}




	go Tasks(ToDoListSize, addToTaskList, takeFromToDoList, listPrint)
	go Magazine(MagazineSize, addToItemsList, takeFromItemsList, magazinePrint)

	for i := 0; i < MultiplicationMachinesAmount; i++ {
		multiplicationMachines[i] = make(chan *resolveTask)
		go multiplicationMachine(i, multiplicationMachines[i], verboseMode)
	}

	for a := 0; a < AdditionMachinesAmount; a++ {
		additionMachines[a] = make(chan *resolveTask)
		go additionMachine(a, additionMachines[a], verboseMode)
	}

	for i := 0; i < 1; i++ {
		go Ceo(addToTaskList, verboseMode)
	}
	for i := 0; i < Employers; i++ {
		go Employer(i, addToItemsList, takeFromToDoList, verboseMode, multiplicationMachines, additionMachines)
	}

	for i := 0; i < Clients; i++ {
		go Client(i, takeFromItemsList, verboseMode)
	}

	time.Sleep(time.Second * 1500)
}