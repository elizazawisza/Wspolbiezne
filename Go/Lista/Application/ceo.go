package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


func createNewTask() Task{
	tmp := rand.Intn(2)
	operationTable  := []string {"addition", "multiplication"}
	var operationName = operationTable[tmp]
	return Task {
		firstArgument: rand.Intn(56789),
		secondArgument: rand.Intn(56789),
			operation: operationName }

}

func TaskToString(task Task) string{
	return strconv.Itoa(task.firstArgument) + " " + strconv.Itoa(task.secondArgument) + " " + task.operation
}

func Ceo(addTask chan *AddToDoList, verbose bool) {
	for{
		var newTask = createNewTask()
		var add = &AddToDoList{
			value:newTask,
			response:make(chan string)}
		addTask <- add
		response := <-add.response
		if verbose {
			fmt.Println(response)
		}
		time.Sleep(DelayCeo)
	}

}
