package main

import (
	"fmt"
	"time"
)

var multiplicationMachineStatistics [MultiplicationMachinesAmount]machineStatistics
var additionMachineStatistics [AdditionMachinesAmount]machineStatistics

func addition(firstElement int, secondElement int) int {
	return firstElement + secondElement
}

func multiplication(firstElement int, secondElement int) int {
	return firstElement * secondElement
}

func maybeResolve(b bool, c chan *resolveTask) chan *resolveTask {
	if !b {
		return nil
	}
	return c
}

func multiplicationMachine(Id int, writesT chan *resolveTask, verbose bool) {
	var busy = false
	var counter = 0
	multiplicationMachineStatistics[Id] = machineStatistics{Id,  counter}
	for {
		select {
		case addItem := <-maybeResolve(!busy, writesT):
			busy = true
			if verbose == true {
				fmt.Println("Employee", addItem.id, "put task :", addItem.value, "into multiplication machine", Id)
			}
			addItem.value.result = multiplication(addItem.value.firstArgument , addItem.value.secondArgument)

			time.Sleep(MultiplicationDelay)
			if verbose == true {
				fmt.Println("Employee", addItem.id, "took task :", addItem.value, "from multiplication machine", Id)
			}
			addItem.response <- addItem.value
			busy = false
			counter++
			multiplicationMachineStatistics[Id].counter = counter
		}
	}
}

func additionMachine(Id int, writesT chan *resolveTask, verbose bool) {
	var busy = false
	var counter = 0
	additionMachineStatistics[Id] = machineStatistics{Id, counter}
	for {
		select {
		case addItem := <-maybeResolve(!busy, writesT):
			busy = true
			if verbose == true {
				fmt.Println("Employee", addItem.id, "put task :", addItem.value, "into adding machine", Id)
			}
			addItem.value.result = addition(addItem.value.firstArgument, addItem.value.secondArgument)

			time.Sleep(AdditionDelay)
			if verbose == true {
				fmt.Println("Employee", addItem.id, "took task :", addItem.value, "from adding machine", Id)
			}
			addItem.response <- addItem.value
			busy = false
			counter++
			additionMachineStatistics[Id].counter = counter
		}
	}
}
