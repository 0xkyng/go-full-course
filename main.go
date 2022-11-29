package main

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("desktop/go-full-course")
	fmt.Println(dir)


	// ? ?= path.Split("secret/file.txt")

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