package main

import (
	"fmt"
	"time"
)

func main() {

	var verboseMode =  false
	var additionMachines [AdditionMachinesAmount]chan *machine
	var multiplicationMachines [MultiplicationMachinesAmount]chan *machine
	var addToTaskList = make (chan *AddToDoList)
	var addToItemsList = make (chan *AddToMagazine)
	var takeFromToDoList = make (chan *TakeFromToDoList)
	var takeFromItemsList = make (chan *TakeFromMagazine)
	var magazinePrint = make (chan *Show)
	var listPrint = make (chan *Show)
	var multiplicationBackdoors = make([]chan bool, 0)
	var additionBackdoors = make([]chan bool, 0)
	var breakdown = BreakdownReport{make(chan int,3),make(chan int, 3)}

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


	for i := 0; i < MultiplicationMachinesAmount; i++ {
		multiplicationMachines[i] = make(chan *machine)
		var backdoor = make(chan bool)
		multiplicationBackdoors = append(multiplicationBackdoors, backdoor)
		go multiplicationMachine(i, multiplicationMachines[i], verboseMode, backdoor, breakdown)
	}

	for a := 0; a < AdditionMachinesAmount; a++ {
		additionMachines[a] = make(chan *machine)
		var backdoor = make(chan bool)
		additionBackdoors = append(multiplicationBackdoors, backdoor)
		go additionMachine(a, additionMachines[a], verboseMode, backdoor, breakdown)
	}



	go Tasks(ToDoListSize, addToTaskList, takeFromToDoList, listPrint)
	go Magazine(MagazineSize, addToItemsList, takeFromItemsList, magazinePrint)
	go Service(breakdown, verboseMode, multiplicationBackdoors, additionBackdoors)


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