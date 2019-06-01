package main

import (
	"fmt"
	"time"
)

var servicemanStatistic[Servicemen] servicemanStatistics

func serviceman(Id int, service service, verbose bool, addBackdoors []chan bool, multiBackdoors []chan bool) {
	var counterAddFixed = 0
	var counterMultiFixed = 0
	servicemanStatistic[Id] = servicemanStatistics{Id,  counterMultiFixed, counterAddFixed}
	for {
		select {
		case machineIndex := <-service.currentRepairAdd:
			time.Sleep(ServicemanDelay)
			addBackdoors[machineIndex] <- false
			counterAddFixed++
			if verbose {
				fmt.Println("Addition Machine number: ", machineIndex," was repaired by serviceman : ", Id , "current fixed: ", counterAddFixed)
			}
			servicemanStatistic[Id].addFixed = counterAddFixed
			service.repairedAddMachine <- machineIndex
		case machineIndex := <-service.currentRepairMulti:
			time.Sleep(ServicemanDelay)
			multiBackdoors[machineIndex] <- false
			counterMultiFixed++
			if verbose {
				fmt.Println("Multiplication Machine number ", machineIndex," was repaired by serviceman : ", Id)
			}
			servicemanStatistic[Id].multiFixed = counterMultiFixed
			service.repairedMultiMachine <- machineIndex
		}
	}

}
