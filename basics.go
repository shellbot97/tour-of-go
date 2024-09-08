package main // Programs start running in package main.

import (
	"fmt"
	"math"
	"math/rand" // package name is the same as the last element of the import path. so if you want to call methods in this package you need to do rand.MethodName()
) // this is called "factored" import statement alternatively you can import each package in a separate line.

const superhero = "Superman" // Constants are declared like variables, but with the const keyword. Constants cannot be declared using the := syntax.

var loveInterest, keepAwayFrom = "Loise Lane", "Kryptonite" // short assignment wont work out of the function. Here you will have to use `var`. If initial value is not present, you will have to specify the type.

var nemesisWithWeapons = map[string]string{
	"Lex Luther":  "Kryptonite",
	"General Zod": "Plasma Carbine",
	"Darkseid":    "Omega Beams",
} // A map maps keys to values. Syntax of this arrangement is map[key_type]value_type. This is covered in details in incoming session.

var nemesis = [3]string{
	"Lex Luther",
	"General Zod",
	"Darkseid",
} // An array's length is part of its type, so arrays cannot be resized.

var canBeKilled bool // Variables declared without an explicit initial value are given their zero value. Incase of string it is "", incase if int it is 0, incase of boolean it is false.

var herosHearingAbility = float64(math.MaxInt64) // The expression T(v) converts the value v to the type T.

func swap(x, y string) (string, string) { // while accepting arguments, function declares argument_type after argument_names(alternatively you can specify argument_type for each argument separately) and specifies return_type/s after the argument_list
	// when a method or variable needs to be exported out of the package, it needs to start with a capital letter. Otherwise it can start with small letters
	return y, x // A function can return any number of results.
}

func turnIntoSuperHero(ordinaryMan string) (hero string) { // naming the return value in go is called `Naked return`
	fmt.Print(ordinaryMan, " turns into ") // fmt's Print, prints the provided arguments to the standard output (usually the console) without adding a newline at the end.
	hero, _ = swap(ordinaryMan, superhero) // when you dont want to use a variable you can catch it in _. Declaring a variable and not using will throw a compile time error
	fmt.Println(hero)                      // println, prints the provided arguments to the standard output with a newline appended at the end. It also adds spaces between arguments.
	return
}

func main() {
	villain := nemesis[rand.Intn(len(nemesis))]

	hero := ""

	fmt.Println(villain, " captures ", loveInterest)

	canBeBeaten := canBeKilled // declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

	if troubleDistance := rand.Float64(); troubleDistance < herosHearingAbility { // if statement can start with a short statement to execute before the condition. This is covered in details in incoming session.
		hero = turnIntoSuperHero("Clarke Kent")
	} else {
		fmt.Printf("%s cant hear %f Kms away", hero, troubleDistance) // Variables declared inside an if short statement are also available inside any of the else blocks.
	}

	if nemesisWithWeapons[villain] == keepAwayFrom {
		canBeBeaten = true
	}

	fmt.Printf("%s has %s and hence %s is beaten? : %t", villain, nemesisWithWeapons[villain], hero, canBeBeaten) // Printf formats the output according to a format string and prints it to the standard output. The format string specifies how the subsequent arguments should be formatted. To find out the type of a variable you can use `%T` along with fmt's  Printf()
}
