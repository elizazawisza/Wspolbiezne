package main

import (
	"time"
)

func Service(breakdownReport BreakdownReport, verbose bool, multiplicationBackdoors []chan bool,
	additionBackdoors []chan bool){
	var totalMachinesAmount = AdditionMachinesAmount + MultiplicationMachinesAmount
	var serviceManager = service{make(chan int, totalMachinesAmount),
		make(chan int, totalMachinesAmount), make(chan int, totalMachinesAmount),
		make(chan int, totalMachinesAmount)}
	var currentFixAddMachine [AdditionMachinesAmount] bool
	var currentFixMultiMachine [MultiplicationMachinesAmount] bool

	for i := 0; i < AdditionMachinesAmount; i++ {
		currentFixAddMachine[i] = false
	}
	for i := 0; i < MultiplicationMachinesAmount; i++ {
		currentFixMultiMachine[i] = false
	}
	for i := 0; i < Servicemen; i++ {
		go serviceman(i, serviceManager, verbose, multiplicationBackdoors, additionBackdoors )
	}

	for {
		select {
		case machineId := <-breakdownReport.addMachines:
			if !currentFixAddMachine[machineId] {
				currentFixAddMachine[machineId] = true
				serviceManager.currentRepairAdd <- machineId
			}
		case machineId := <-breakdownReport.multiMachines:
			if !currentFixMultiMachine[machineId] {
				currentFixMultiMachine[machineId] = true
				serviceManager.currentRepairMulti <- machineId
			}
		case machineId := <-serviceManager.repairedAddMachine:
			currentFixAddMachine[machineId] = false

		case machineId := <-serviceManager.repairedMultiMachine:
			currentFixMultiMachine[machineId] = false
		}
		time.Sleep(ServiceDelay)
	}
}
