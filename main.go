package main

import (
	"fmt"
	"os"
)

func main() {
	// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------
	var (
		args = os.Args
		l = len(args) - 1
	)
	
if l == 0 {
	fmt.Printf("Give me args")
} else if l == 1 {
	fmt.Printf("There is one: %q\n", args[1])
} else if l == 2 {
	fmt.Printf(
		`There are two: "%s %s"`+"\n",
		args[1], args[2],
	)
} else {
	fmt.Printf("There are %d arguments\n", 1)
}

// isSphere, radius := true, 200

// if isSphere && radius >= 200 {
// 	fmt.Printf("It's a big sphere.")
// } else {
// 	fmt.Printf("I don't know.")
// }

// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up

// age := 10
// if age > 60 {
// 	fmt.Printf("Getting older")
// } else if age > 30 {
// 	fmt.Printf("Getting wiser")
// } else if age > 20 {
// 	fmt.Printf("Adulthood")
// } else if age > 10 {
// 	fmt.Printf("Young blood")
// } else {
// 	fmt.Printf("Booting up")
// }

	// name, lastName := os.Args[1], os.Args[2]
	// msg := "my name is %s and my last name is %s\n"
	// fmt.Printf(msg, name, lastName)

	// fmt.Printf(`"Hello word"`)

	// temp := 29.5
	// fmt.Printf("Temperature is %f in my area now\n", temp)

	// tf:= false
	// fmt.Printf("These are %t claims\n", tf)

	// fmt.Printf("My name is %s and my last name is %s\n", "isaac", "Wanger")
	// msg:= "My name is %s and my last name is %s\n"
	// fmt.Printf(msg, "isaac", "wanger")

	// age := 30
	// fmt.Printf("I am %d years old\n", age)

	// const (
	// 	Spring = (iota + 1) * 3
	// 	Summer
	// 	Fall
	// 	Winter
	// )
	// fmt.Println(Winter, Spring, Summer, Fall)

	// const (
	// 	jan = iota + 1
	// 	Feb
	// 	March
	// )
	// fmt.Println(jan, Feb, March)

	// const (
	// 	Nov = 11 - iota // 11 - 0 = 11
	// 	Oct             // 11 - 1 = 10
	// 	Sep             // 11 - 2 = 9
	// )

	// fmt.Println(Sep, Oct, Nov)


	// const (
	// 	monday = iota
	// 	tuesday
	// 	wednesday
	// 	thursday
	// 	friday
	// 	saturday
	// 	sunday
	// )
	// fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	// name := os.Args[1]
	// fmt.Println(strings.ToLower(os.Args[1]))

	//---------------------------------------------------
	// EXERCISE: Raw Concat
	//
	//  1. Initialize the name variable
	//     by getting input from the command line
	//
	//  2. Use concatenation operator to concatenate it
	//     with the raw string literal below
	//
	// NOTE
	//  You should concatenate the name variable in the correct
	//  place.
	//
	// EXPECTED OUTPUT
	//  Let's say that you run the program like this:
	//   go run main.go inanç
	//
	//  Then it should output this:
	//   hi inanç!
	//   how are you?

	// name := os.Args[1]

	// replace and concatenate the `name` variable
	// after `hi ` below
	// fmt.Println(name)

	// msg := `Hi`  +  name  + `! how are you?`

	// fmt.Println(msg)
	// ---------------------------------------------------------

	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	// json := `
	// {
	// 	"items":
	// 	[{
	// 		"item":
	// 		{
	// 			"name": "Isaac Tee"

	// 		}
	// 	}]
	// }
	// `
	// fmt.Println(json)

	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	// json := "\n" +
	// 	"{\n" +
	// 	"\t\"Items\": [{\n" +
	// 	"\t\t\"Item\": {\n" +
	// 	"\t\t\t\"name\": \"Teddy Bear\"\n" +
	// 	"\t\t}\n" +
	// 	"\t}]\n" +
	// 	"}\n"

	// fmt.Println(json)

	// 	path := `c:\\program files\\duper super\\fun.txt\
	// 	c:\\program files\\really\\funny.png`

	// fmt.Println(path)

	// msg := os.Args[1]
	// l := len(msg)

	// s := msg + strings.Repeat("!", l)
	// s = strings.ToUpper(s)

	// fmt.Println(s)

	// counting in strings
	// utf8.RuneCountInString(name)

	// dir, _ := path.Split("desktop/go-full-course")
	// fmt.Println(dir)

	// red, blue := "red", "blue"
	// red, blue = blue, red
	// fmt.Println(red, blue)

	// var (
	// 	planet string
	// 	isTrue bool
	// 	temp   float64
	// )

	// planet = "Mars"
	// isTrue = true
	// temp = 19.5
	// fmt.Println("Air is good on ", planet)
	// fmt.Println("It's ", isTrue)
	// fmt.Println("It is", temp, "degrees")

	// DO NOT TOUCH THIS
	// var (
	// 	lang    string
	// 	version int
	// )

	// ADD YOUR CODE BELOW
	// lang = "Go"
	// version = 2

	// DO NOT TOUCH THIS
	// fmt.Println(lang, "version", version)

	// UNCOMMENT THE CODE BELOW:
	// var (
	// 	perimeter        int
	// 	width, height = 5, 6
	// )

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT
	// perimeter = 2 *(width * height)
	// fmt.Println(perimeter)

	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	// n := 0.

	// ADD YOUR CODE BELOW
	// n = float64(3.14 * 2)
	// fmt.Println(n)

	// color := "green"

	// // `"dark " + color` is an expression

	// color = "dark " + color

	// fmt.Println(color)

	// color := "green"
	// color = "blue"
	// fmt.Println(color)

	// speed := 100
	// force := 2.5

	// // type conversion
	// speed = int(float64(speed) * force)
	// fmt.Println(speed)

	// Walking with os.Args
	// fmt.Printf("%#v\n", os.Args)

	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("1st argument:", os.Args[1])
	// fmt.Println("2nd argument:", os.Args[2])
	// fmt.Println("3rd argument:", os.Args[3])

	// fmt.Println("Number of items inside os.Args:", len(os.Args))
}