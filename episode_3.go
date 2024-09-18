package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type room struct {
	name              string
	doorToTheNextRoom *room
} // A struct is a collection of fields. A struct literal denotes a newly allocated struct value by listing the values of its fields (literals are constant values written directly into the code).

type batcave struct {
	entrance *room
}

var roomsInBatcave = map[string]string{
	"entrace":   "Batcave Entrance",
	"room1":     "Batcave Computer",
	"room2":     "Batcave Armory",
	"last_room": "Batcave Garage",
} // A map maps keys to values. These are mutable

var (
	room1 = &room{name: roomsInBatcave["entrace"], doorToTheNextRoom: room2} // using `&` you can assign pointer of room to room1. A pointer holds the memory address of a value. you can also change the value of i with help of *p. This is known as "dereferencing" or "indirecting" for eg. p := &i; *p = 21.
	// Unlike C, Go has no pointer arithmetic. This is when you perform arithmetic operations on the pointer itself, not the value it points to. Pointer arithmetic allows incrementing or decrementing the pointer to move to different memory addresses. Go does not allow this.
	room2 = &room{name: roomsInBatcave["room1"], doorToTheNextRoom: room3}
	room3 = &room{name: roomsInBatcave["room2"], doorToTheNextRoom: room4}
	room4 = &room{name: roomsInBatcave["last_room"], doorToTheNextRoom: nil}
) // just like factored imports, you can also collectively declare variables

var countDown = [10]int{
	10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
}

func start() {
	bcave := batcave{entrance: room1}
	device := bcave.teleport()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("All rooms visited")
		}
	}()

	for {
		device()
	}
}

func (batcave *batcave) teleport() func() string { // accessing i with help of *p.
	currentRoom := batcave.entrance
	countDownForRoom := countDown[:] //  The type []T is a slice with elements of type T. As opposed to arrays slices are dynamic in size.
	// A slice does not store any data, it just describes a section of an underlying array.
	// In the slice, a[num1:num2] num1 is the starting index (inclusive). num2 is the ending index (exclusive).
	// a[num1:] means that slice from an element to the end of the array
	// a[:num1] means that slice from the beginning to a specific element
	// a[:] means a slice that starts from the first element and goes up to, but does not exclude, the last element.
	return func() string {
		// go functions maybe closure
		// a closure is a function that references variables outside of its body
		// imagine a scenario when you have to find a sum of 0 to 5. Typically this would require following approach
		/*
			var sum = 0
			func add(new_int int) {
				sum = sum + new_int
			}
			func main() {
				for i, _ := range [5]int{} {
					add(i)
				}
				fmt.Println("Sum : ", sum)
			}
		*/
		// but if you can bind the variable sum with the function add, it would make it easier. With closure we can do this
		roomName := currentRoom.name // Struct fields are accessed using a dot.
		fmt.Println("Entered, and planted bomb in ", roomName)

		var rooms []string = []string{"Batcave Armory", "Batcave Computer", "Batcave Entrance"} // A slice literal is like an array literal without the length.
		if roomName == rooms[rand.Intn(len(rooms))] {
			fmt.Println("Batman is here in this room, he defuses the bomb..")
			countDown[0] = -1 // Slices are mutable ie. CHANGING THE ELEMENTS OF A SLICE MODIFIES THE CORRESPONDING ELEMENTS OF ITS UNDERLYING ARRAY. This doesnt apply to append or delete ie. Appending/Deleting an element from a slice in Go doesn't append or delete it in the underlying array
			// OTHER SLICES THAT SHARE THE SAME UNDERLYING ARRAY WILL SEE THOSE CHANGES.
		}
		var second int
		for second = range countDownForRoom {
			if countDownForRoom[second] == -1 {
				fmt.Println("Bombs defused!!!!!!!")
				break
			} else {
				warningMessage := func(remainingSeconds int) string {
					return "Bomb will explode in " + strconv.Itoa(remainingSeconds) + " seconds"
				}
				fmt.Println(bombExplosionWarning(warningMessage, countDownForRoom[second])) // functions can be passed around as an argument to other functions
			}
		}
		if second != 0 {
			fmt.Println("Skadoosh! BOOM! Skapow! ", "Bomb exploded in ", currentRoom) // interesting thing to note here is, printing pointer to a complex data structure like struct gives out a more readable form `&{Batcave Computer 0x100a6ae40}` as opposed to simpler data types like int (which gives memory address)
		}
		currentRoom = currentRoom.doorToTheNextRoom
		return roomName
	}
}

func bombExplosionWarning(warning func(int) string, remainingSeconds int) string {
	return warning(remainingSeconds)
}
