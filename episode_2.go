package main

import (
	"fmt"
	"math/rand"
)

const birthYearOfTheTimeTraveler = 1997
const maxTimeThatCanBeSpentInTimeTravel = 5


func episode2_start() {
	avengersHeadquarters()
}

func avengersHeadquarters() {
	quantumTunnel()
	fmt.Println("Returned normally from quantumTunnel to AvengersHeadquarters.")
}

func quantumTunnel() {
	var scientists = [2]string{
		"Hulk",
		"Tony Stark",
	}
	switch scientistMonitoringTimeTravel := scientists[rand.Intn(len(scientists))]; scientistMonitoringTimeTravel { // Switch without a condition is the same as switch true.
	case "Hulk": // Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go
		defer fmt.Println("You recovered because hulk was here!")
		defer func() {
			if r := recover(); r != nil { // nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
				fmt.Println("Came back from year: ", r)
				fmt.Println("Saved Ant Man from getting stuck in Quantum realm")
			}
		}() // defers the execution of a function until the surrounding function returns.
	case "Tony Stark":
		defer fmt.Println("oh no! hulk is not here! going into panic mode!!!!!!!!!")
	default:
		defer fmt.Println("oh no! no one is here! going into panic mode!!!!!!!!!")
	}
	fmt.Println("Starting time travel from year", 2024)
	goToYear(getPreviousYear(2024))
}

func goToYear(year int) {
	fmt.Println("Traveling in ", year)           // A deferred function’s arguments are evaluated when the defer statement is evaluated. So this will always return value of year that is there before the above statement gets executed
	defer fmt.Println("Coming back from ", year) // Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	// In this example, the expression "year" is evaluated when the Println call is deferred, so it doesnt matter if it gets modified after this point in a program. It will print the exact value which it had stored while evaluating this statement.
	if year < birthYearOfTheTimeTraveler {
		fmt.Println("Panicking: You cant go beyond your birth year!")
		panic(fmt.Sprintf("%v", year))
	}
	goToYear(getPreviousYear(year))
	hour := 1
	for hour < maxTimeThatCanBeSpentInTimeTravel { // // basic for consists of init statement, condition expression and post statement. if only condition expression is provided for loop behaves like while. If even thats not provided loop goes into infinite iterations
		hour += 1
	}
}

func getPreviousYear(year int) (previousYear int) {
	mischievousValueAddedByLokiHopingItWillTriggerPanic := 1990
	previousYear = mischievousValueAddedByLokiHopingItWillTriggerPanic
	defer func() { previousYear = year - 1 }() // In this example, a deferred function changes the return value previousYear after the surrounding function returns. Thus this will never give mischievousValueAddedByLokiHopingItWillTriggerPanic
	return previousYear
}
