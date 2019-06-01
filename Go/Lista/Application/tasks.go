package main

import "strconv"

func maybeAddToTaskList(b bool,c chan *AddToDoList) chan *AddToDoList{
	if !b {
		return nil
	}
	return c
}

func maybeTakeFromTaskList(b bool,c chan *TakeFromToDoList) chan *TakeFromToDoList{
	if !b {
		return nil
	}
	return c
}

func tasksToString(tasks []Task) []string{
	var strs []string
	for i := 0; i< len(tasks); i++ {
		strs = append(strs, strconv.Itoa(tasks[i].firstArgument) + " " + strconv.Itoa(tasks[i].secondArgument) + " " + tasks[i].operation)
		strs = append(strs, "\n")
	}
	return strs
}

func taskToString(task Task) []string{
	var strs []string
		strs = append(strs, strconv.Itoa(task.firstArgument) + " " + strconv.Itoa(task.secondArgument) + " " + task.operation)
	return strs
}


func Tasks (limit int, addToTaskList chan *AddToDoList, takeFromTaskList chan *TakeFromToDoList, showToDoList chan *Show){

	var taskList = make([]Task, 0)
	for{
		select {
		case add := <-maybeAddToTaskList(len(taskList) < limit, addToTaskList):
			taskList = append(taskList, add.value)
			add.response <- "CEO add new Task to the Task List: " + TaskToString(add.value)

		case take := <-maybeTakeFromTaskList(len(taskList) > 0, takeFromTaskList):
			take.task <- taskList[0]
			taskList = append(taskList[:0], taskList[0+1:]...)
			take.response <- "Employer took one task from the Task List: "
		case show := <-showToDoList:
			show.response <- tasksToString(taskList)
		}
	}

}
