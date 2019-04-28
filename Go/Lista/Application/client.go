package main

import (
	"fmt"
	"time"
)

func Client(Id int, buyItem chan *TakeFromMagazine, verbose bool){
	for{
		buy := &TakeFromMagazine{
			id: Id,
			item: make(chan int),
			response: make(chan string)}
		buyItem <- buy
		item := <-buy.item
		response := <-buy.response
		if verbose {
			fmt.Println(response, item)
		}
		time.Sleep(DelayClient)
	}

}
