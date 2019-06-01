package main

import (
	"fmt"
	"math/rand"
	"time"
)
var types = [2]string{"patient", "impatient"}
var employersStatistics [Employers]statistics

func chooseAdditionMachine(machines [AdditionMachinesAmount]chan *machine) chan *machine {
	return machines[rand.Intn(len(machines))]
}
func chooseMultiplicationMachine(machines [MultiplicationMachinesAmount]chan *machine) chan *machine {
	return machines[rand.Intn(len(machines))]
}

func whoAmI() string {
	return types[rand.Intn(len(types))]
}

func getRandIdx(amount int) int{
	return rand.Intn(amount)
}

func Employer(Id int, addItem chan *AddToMagazine, takeTask chan *TakeFromToDoList, verbose bool,
	multiplicationMachines [MultiplicationMachinesAmount]chan *machine, additionMachines [AdditionMachinesAmount]chan *machine){
	var employerType = whoAmI()
	var counter = 0
	var idx = 0
	employersStatistics[Id] = statistics{Id, employerType, counter}

	for{
		taken := &TakeFromToDoList{
			id: Id,
			task: make(chan Task),
			response: make(chan string)}
		takeTask <- taken
		task := <- taken.task
		response := <- taken.response
		if verbose {
			fmt.Println(response, TaskToString(task))
		}


		var resolvedTask Task
		calc := &machine{
			id:       Id,
			value:    task,
			response: make(chan Task)}

		var machine chan *machine

		if task.operation == "addition" {
			idx = getRandIdx(AdditionMachinesAmount)
			machine = additionMachines[idx]
		} else if task.operation == "multiplication" {
			idx = getRandIdx(MultiplicationMachinesAmount)
			machine = multiplicationMachines[idx]
		}

		if employerType == "impatient" {
			done := false
			for !done {
				select {
				case machine <- calc:
					resolvedTask = <-calc.response
					done = true
				case <-time.After(PatientEmployerDelay):
					if task.operation == "addition" {
						machine = chooseAdditionMachine(additionMachines)
					} else if task.operation == "multiplication" {
						machine = chooseMultiplicationMachine(multiplicationMachines)
					}
				}
			}
		} else {
			machine <- calc
			resolvedTask = <-calc.response
		}
		myResult := resolvedTask.result
		counter++
		employersStatistics[Id].counter = counter

		add := &AddToMagazine{
			id: Id,
			value: myResult,
			response: make(chan string)}
		addItem <- add
		responseMagazine := <- add.response
		if verbose {
			fmt.Println(responseMagazine, resolvedTask)
		}
		time.Sleep(DelayEmployer)
	}
}