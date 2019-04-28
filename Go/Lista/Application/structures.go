package main

type Task struct{
	firstArgument int
	secondArgument int
	operation string
	result int
}

type AddToMagazine struct {
	id int
	value  int
	response chan string
}
type TakeFromMagazine struct {
	id int
	item  chan int
	response chan string
}

type AddToDoList struct {
	id int
	value Task
	response chan string
}
type TakeFromToDoList struct {
	id int
	task  chan Task
	response chan string
}

type Show struct {
	response  chan []string
}

type statistics struct {
	id           int
	employerType string
	counter      int
}

type resolveTask struct {
	id   int
	value  Task
	response chan Task
}