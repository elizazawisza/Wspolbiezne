package main

import (
	"fmt"
	"math/rand"
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

func maybeResolve(b bool, c chan *machine) chan *machine {
	if !b {
		return nil
	}
	return c
}



func multiplicationMachine(Id int, writesT chan *machine, verbose bool, backdoor chan bool, damage BreakdownReport) {
	var busy = false
	var isBroken = false
	var taskCounter = 0
	var breakdownsCounter = 0
	var multiItem *machine
	multiplicationMachineStatistics[Id] = machineStatistics{Id, taskCounter, breakdownsCounter}
	for {
		select {
		case multiItem = <-maybeResolve(!busy, writesT):
			if isBroken {
				if verbose {
					fmt.Println("Multiplication machine", Id, "is broken")
				}
			}else {
				busy = true
				if verbose {
					fmt.Println("Employee", multiItem.id, "put task :", multiItem.value, "into multiplication machine", Id)
				}
				multiItem.value.result = multiplication(multiItem.value.firstArgument, multiItem.value.secondArgument)

				time.Sleep(MultiplicationDelay)
				if verbose {
					fmt.Println("Employee", multiItem.id, "took task :", multiItem.value, "from multiplication machine", Id)
				}
				multiItem.response <- multiItem.value
				busy = false
				taskCounter++
				isBroken = rand.Float32() < BreakDownProbability
				if isBroken{
					breakdownsCounter++
					damage.multiMachines <- Id
				}
				multiplicationMachineStatistics[Id].computedTaskCounter = taskCounter
				multiplicationMachineStatistics[Id].breakdownCounter = breakdownsCounter

			}
		case isBrokenMessage := <- backdoor:
			if isBrokenMessage == false {
				isBroken = false
			}
		}
	}
}

func additionMachine(Id int, writesT chan *machine, verbose bool, backdoor chan bool, damage BreakdownReport) {
	var busy = false
	var isBroken = false
	var taskCounter = 0
	var breakdownsCounter = 0
	additionMachineStatistics[Id] = machineStatistics{Id, taskCounter ,breakdownsCounter}
	for {
		select {
		case addItem := <-maybeResolve(!busy, writesT):
			if isBroken {
				if verbose {
					fmt.Println("Addition machine", Id, "is broken")
				}
			}else{
				busy = true
				if verbose {
					fmt.Println("Employee", addItem.id, "put task :", addItem.value, "into adding machine", Id)
				}
				addItem.value.result = addition(addItem.value.firstArgument, addItem.value.secondArgument)

				time.Sleep(AdditionDelay)
				if verbose {
					fmt.Println("Employee", addItem.id, "took task :", addItem.value, "from adding machine", Id)
				}
				addItem.response <- addItem.value
				busy = false
				taskCounter++
				isBroken = rand.Float32() < BreakDownProbability
				if isBroken {
					breakdownsCounter++
					damage.addMachines <- Id
				}
				additionMachineStatistics[Id].computedTaskCounter = taskCounter
				additionMachineStatistics[Id].breakdownCounter = breakdownsCounter
			}
			case isBrokenMessage := <- backdoor:
				if isBrokenMessage == false {
					isBroken = false
				}
		}
	}
}