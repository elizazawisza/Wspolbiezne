package main

import (
	"fmt"
	"time"
)

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

func multiplicationMachine(id int, writesT chan *resolveTask, verbose bool) {
	var busy = false
	for {
		select {
		case addItem := <-maybeResolve(!busy, writesT):
			busy = true
			if verbose == true {
				fmt.Println("Employee", addItem.id, "put task :", addItem.value, "into multiplication machine", id)
			}
			addItem.value.result = multiplication(addItem.value.firstArgument , addItem.value.secondArgument)

			time.Sleep(MultiplicationDelay)
			if verbose == true {
				fmt.Println("Employee", addItem.id, "took task :", addItem.value, "from multiplication machine", id)
			}
			addItem.response <- addItem.value
			busy = false
		}
	}
}

func additionMachine(id int, writesT chan *resolveTask, verbose bool) {
	var busy = false
	for {
		select {
		case addItem := <-maybeResolve(!busy, writesT):
			busy = true
			if verbose == true {
				fmt.Println("Employee", addItem.id, "put task :", addItem.value, "into adding machine", id)
			}
			addItem.value.result = addition(addItem.value.firstArgument, addItem.value.secondArgument)

			time.Sleep(AdditionDelay)
			if verbose == true {
				fmt.Println("Employee", addItem.id, "took task :", addItem.value, "from adding machine", id)
			}
			addItem.response <- addItem.value
			busy = false
		}
	}
}
