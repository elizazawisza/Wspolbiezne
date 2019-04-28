package main

import "strconv"

func maybeAddToMagazine(b bool,c chan *AddToMagazine) chan *AddToMagazine{
	if !b {
		return nil
	}
	return c
}

func maybeTakeFromMagazine(b bool,c chan *TakeFromMagazine) chan *TakeFromMagazine{
	if !b {
		return nil
	}
	return c
}

func itemsToString(items []int) []string{
	var strs []string
	for i := 0; i< len(items); i++ {
		strs = append(strs, strconv.Itoa(items[i]))
		strs = append(strs, "\n")
	}
	return strs
}

func Magazine(limit int, addItemToMagazine chan *AddToMagazine, takeItemFromMagazine chan *TakeFromMagazine, showMagazine chan *Show){
	var magazine = make([]int, 0)
	for{
		select {
		case add := <-maybeAddToMagazine(len(magazine) < limit, addItemToMagazine):
			magazine = append(magazine, add.value)
			add.response <- "Employer " + strconv.Itoa(add.id) + " finished the task: "

		case take := <-maybeTakeFromMagazine(len(magazine) > 0, takeItemFromMagazine):
			take.item <- magazine[0]
			magazine = append(magazine[:0], magazine[0+1:]...)
			take.response <- "Client " + strconv.Itoa(take.id) + " bought item with value"

		case show := <-showMagazine:
			show.response <- itemsToString(magazine)
		}
	}
}
