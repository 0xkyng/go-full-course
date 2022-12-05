package main


func main() {
	// ARRAYS.
	// var books [4]string
	// books[0] = "Kafka's Revenge"
	// books[1] = "stay golden"
	// books[2] = "Everythingship"
	// books[3] = "kafka's revenge 2nd edition"
// const (
// 	winter = 1
// 	summer = 3
// 	yearly = winter + summer
// )
// 	// Array Literal
// 	books := [yearly] string{
// 		"Kafka's Revenge",
// 		"stay golden",
// 		"Everythingship",
// 		"Kafaka's Revenge 2nd edition",
// 	}
// 	fmt.Println(books)

	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	// RANDOMNIZATION
	// t := time.Now() // current time

	// // or rand.Seed(time.Now().UnixNano())
	
	// rand.Seed(t.UnixNano()) // seeding time with unixnano time
	
	// guess := 10

	// for n := 0; n != guess; {
	// 	n = rand.Intn(guess + 1)
	// 	fmt.Printf("%d ", n)
	// }
	// fmt.Println()
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	// 	var sum int
	// 	for i := 1; i <= 10; i++ {
	// 		sum += 1
	// 		fmt.Println("sum:", sum)
	// 	}
	// RANGING
	//  for _, v := range os.Args[1:] {
	// 	fmt.Printf("%q\n", v)
	//  }

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// MULTIPLICATION TABLE
	// const max = 5
	// fmt.Printf("%5s", "X")
	// for i := 0; i <= max; i++ {
	// 	fmt.Printf("%5d", i)
	// }
	// fmt.Println()
	// for i := 0; i <= max; i++ {
	// 	fmt.Printf("%5d", i)

	// 	for j := 0; j <= max; j++ {
	// 		fmt.Printf("%5d",i*j)
	// 	}
	// 	fmt.Println()
	// }
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// var (
	//  sum int
	// 	i = 1
	// )

	// for {
	// 	if i > 5{
	// 		break
	// 	}
	// 	sum += 1
	// 	i++
	// }
	// fmt. Println(sum)

	/////////////////////////////////////////////////////////////////////////////////
	// SWITCH EXERCISES

	// EXERCISE: Richter Scale
	//
	//  1. Get the earthquake magnitude from the command-line.
	//  2. Display its corresponding description.
	//
	// ---------------------------------------------------------
	// MAGNITUDE                    DESCRIPTION
	// ---------------------------------------------------------
	// Less than 2.0                micro
	// 2.0 and less than 3.0        very minor
	// 3.0 and less than 4.0        minor
	// 4.0 and less than 5.0        light
	// 5.0 and less than 6.0        moderate
	// 6.0 and less than 7.0        strong
	// 7.0 and less than 8.0        major
	// 8.0 and less than 10.0       great
	// 10.0 or more                 massive
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Give me the magnitude of the earthquake.
	//
	//  go run main.go ABC
	//    I couldn't get that, sorry.
	//
	//  go run main.go 0.5
	//    0.50 is micro
	//
	//  go run main.go 2.5
	//    2.50 is very minor
	//
	//  go run main.go 3
	//    3.00 is minor
	//
	//  go run main.go 4.5
	//    4.50 is light
	//
	//  go run main.go 5
	//    5.00 is moderate
	//
	//  go run main.go 6
	//    6.00 is strong
	//
	//  go run main.go 7
	//    7.00 is major
	//
	//  go run main.go 8
	//    8.00 is great
	//
	//  go run main.go 11
	//    11.00 is massive

	// args := os.Args
	// if len(args) != 2 {
	// 	fmt.Printf("Give me the magnitude of the earthquake.")
	// 	return
	// }

	// richter, err := strconv.ParseFloat(args[1], 64)
	// if err != nil {
	// 	fmt.Printf("I couldn't get that, sorry.")
	// 	return
	// }

	// var desc string
	// r := richter

	// switch {
	// case r < 2:
	// 	desc = "micro"
	// case r >= 2 && r < 3:
	// 	desc = "very minor"
	// case r >= 3 && r < 4:
	// 	desc = "minor"
	// case r >= 4 && r < 5:
	// 	desc = "light"
	// case r >= 5 && r < 6:
	// 	desc = "moderate"
	// case r >= 6 && r < 7:
	// 	desc = "strong"
	// case r >= 7 && r < 8:
	// 	desc = "major"
	// case r >= 8 && r < 9:
	// 	desc = "great"
	// default:
	// 	desc = "massive"
	// }
	// fmt.Printf("%.2f is %s\n", richter, desc)
	/////////////////////////////////////////////////////////////////////////////////

	// ---------------------------------------------------------
	// EXERCISE: Convert
	//
	//  Convert the if statement to a switch statement.
	// ---------------------------------------------------------
	// const (
	// 	usage       = "Usage: [username] [password]"
	// 	errUser     = "Access denied for %q.\n"
	// 	errPwd      = "Invalid password for %q.\n"
	// 	accessOK    = "Access granted to %q.\n"
	// 	user, user2 = "jack", "inanc"
	// 	pass, pass2 = "1888", "1879"
	// )
	// args := os.Args

	// if len(args) != 3 {
	// 	fmt.Println(usage)
	// 	return
	// }

	// u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }

	// REFACTORED INTO SWITCH
	// switch  {
	// 	case u != user && u != user2:
	// 		fmt.Printf(errUser, u)
	// 	case u == user && p == pass:
	// 		fallthrough
	// 	case u == user2 && p == pass2:
	// 		fmt.Printf(accessOK, u)
	// 	default:
	// 		fmt.Printf(errPwd, u)
	// }
	//---------------------------------------------------------------------------------

	//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	// EXERCISE: String Manipulator
	//
	//  1. Get the operation as the first argument.
	//
	//  2. Get the string to be manipulated as the 2nd argument.
	//
	// HINT
	//  You can find the manipulation functions in the strings
	//  package Go documentation (ToLower, ToUpper, Title).
	//
	// EXPECTED OUTPUT
	//
	//  go run main.go
	//    [command] [string]
	//
	//    Available commands: lower, upper and title
	//
	//  go run main.go lower 'OMG!'
	//    omg!
	//
	//  go run main.go upper 'omg!'
	//    OMG!
	//
	//  go run main.go title "mr. charles darwin"
	//    Mr. Charles Darwin
	//
	//  go run main.go genius "mr. charles darwin"
	//    Unknown command: "genius"

	// const usage = `[command] [string]
	// Availabe commands: lower, upper and title`
	// args := os.Args
	// if len(os.Args) != 3 {
	// 	fmt.Printf(usage)
	// 	return
	// }

	// cmd, str := args[1], args[2]

	// switch cmd {
	// case "lower":
	// 	fmt.Printf(strings.ToLower(str))
	// case "upper":
	// 	fmt.Println(strings.ToUpper(str))
	// case "title":
	// 	fmt.Println(strings.Title(str))
	// default:
	// 	fmt.Printf("unknown comand: %q\n", cmd)
	// }
	//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

	//||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
	// EXERCISE: Days in a Month
	//
	//  Refactor the previous exercise from the if statement
	//  section.
	//
	//  "Print the number of days in a given month."
	//
	//  Use a switch statement instead of if statements.
	// if len(os.Args) != 2 {
	// 	fmt.Println("Give me a month name")
	// 	return
	// }

	// year := time.Now().Year()
	// leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// days, month := 28, os.Args[1]

	// 	if m := strings.ToLower(month); m == "april" ||
	// 		m == "june" ||
	// 		m == "september" ||
	// 		m == "november" {
	// 		days = 30
	// 	} else if m == "january" ||
	// 		m == "march" ||
	// 		m == "may" ||
	// 		m == "july" ||
	// 		m == "august" ||
	// 		m == "october" ||
	// 		m == "december" {
	// 		days = 31
	// 	} else if m == "february" {
	// 		if leap {
	// 			days = 29
	// 		}
	// 	} else {
	// 		fmt.Printf("%q is not a month.\n", month)
	// 		return
	// 	}

	// 	fmt.Printf("%q has %d days.\n", month, days)
	// }

	// switch m := strings.ToLower(month); m {
	// case "april", "june", "september", "november":
	// 	fmt.Println("days = 30")
	// case "january", "march", "may", "july", "august", "october", "december":
	// 	fmt.Println("days = 30")
	// case "februry":
	// 	if leap{
	// 		days =29
	// 	}
	// default:
	// 	fmt.Printf("%q has %d days.\n", month, days)
	// }

	//||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	// h := time.Now().Hour()
	// fmt.Println("current time is", h)

	// switch {
	// case h < 12:
	// 	fmt.Printf("Good morning")
	// case h >= 12:
	// 	fmt.Printf("Good afternoon")
	// case h >= 18:
	// 	fmt.Printf("good evening")
	// default:
	// 	fmt.Printf("good night")
	// }
	// FALLTHROUGH KEYWORD
	// i := 142
	// switch {
	// case i > 100:
	// 	fmt.Printf("big")
	// 	fallthrough
	// case i > 0:
	// 	fmt.Printf("positive")
	// 	fallthrough
	// default:
	// 	fmt.Printf("number")
	// }
	// i := -10

	// switch  {
	// case i > 0:
	// 	fmt.Printf("positive")
	// case i < 0:
	// 	fmt.Printf("negative")
	// default:
	// 	fmt.Printf("zero")
	// }

	// city := os.Args[1]

	// switch city {
	// case "Abuja":
	// 	fmt.Printf("nigeria")
	// case "Accra":
	// 	fmt.Printf("Ghana")
	// default:
	// 	fmt.Printf("where?")
	// }

	/////////////////////////////////////////////////////////////////////////////////////////////
	// HANDLING ERROR EXERCISES
	// 1. RATINGS MOVIE

	//  1. Get the age from the command-line.
	//
	//  2. Return one of the following messages if the age is:
	//     -> Above 17         : "R-Rated"
	//     -> Between 13 and 17: "PG-13"
	//     -> Below 13         : "PG-Rated"
	//
	// RESTRICTIONS:
	//  1. If age data is wrong or absent let the user know.
	//  2. Do not accept negative age.
	//
	// BONUS:
	//  1. BONUS: Use if statements only twice throughout your program.
	//  2. BONUS: Use an if statement only once.
	//
	// EXPECTED OUTPUT
	//  go run main.go 18
	//    R-Rated
	//
	//  go run main.go 17
	//    PG-13
	//
	//  go run main.go 12
	//    PG-Rated
	//
	//  go run main.go
	//    Requires age
	//
	//  go run main.go -5
	//    Wrong age: "-5"

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Requires age")
	// 	return
	// }

	// age, err := strconv.Atoi(os.Args[1])

	// if err != nil || age < 0 {
	// 	fmt.Printf(`wrong age: %q` +"\n", os.Args[1])
	// } else if age > 17 {
	// 	fmt.Printf("R-RATED")
	// } else if age >= 13 && age <= 17 {
	// 	fmt.Printf("PG-13")
	// }else if age < 13 {
	// 	fmt.Printf("PG-RATED")
	// }

	// 2
	// EXERCISE: Odd or Even
	//
	//  1. Get a number from the command-line.
	//
	//  2. Find whether the number is odd, even and divisible by 8.
	//
	// RESTRICTION
	//  Handle the error. If the number is not a valid number,
	//  or it's not provided, let the user know.
	//
	// EXPECTED OUTPUT
	//  go run main.go 16
	//    16 is an even number and it's divisible by 8
	//
	//  go run main.go 4
	//    4 is an even number
	//
	//  go run main.go 3
	//    3 is an odd number
	//
	//  go run main.go
	//    Pick a number
	//
	//  go run main.go ABC
	//    "ABC" is not a number

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Pick a number")
	// 	return
	// }

	// n, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("%d is not a number\n", n)
	// 	return
	// }

	// if n%8 == 0 {
	// 	fmt.Printf("%d is not an even number and it's divisible by 8\n", n)
	// } else if n%2 ==0 {
	// 	fmt.Printf("%d is an even number\n", n)
	// } else {
	// 	fmt.Printf("5d is an odd number")
	// }

	// XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	//. 3
	// EXERCISE: Leap Year
	//
	//  Find out whether a given year is a leap year or not.
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Give me a year number
	//
	//  go run main.go eighties
	//    "eighties" is not a valid year.
	//  go run main.go 2018
	//    2018 is not a leap year.
	//
	//  go run main.go 2100
	//    2100 is not a leap year.
	//
	//  go run main.go 2019
	//    2019 is not a leap year.
	//
	//  go run main.go 2020
	//    2020 is a leap year.
	//
	//  go run main.go 2024
	//    2024 is a leap year.

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Give me a year number")
	// 	return
	// }

	// year, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("%q is not a valid year\n", os.Args[1])
	// 	return
	// }

	// var leap bool
	// if year%400 == 0 {
	// 	leap = true
	// } else if year%100 == 0 {
	// 	leap =false
	// } else if year%4 == 0{
	// 	leap = true
	// }

	// if leap {
	// 	fmt.Printf("%d is a leap yaer.\n", year)
	// } else {
	// 	fmt.Printf("%d is not a leap.\n", year)
	// }
	// XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

	//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	// 4.
	// EXERCISE: Days in a Month
	//
	//  Print the number of days in a given month.
	//
	// RESTRICTIONS
	//  1. On a leap year, february should print 29. Otherwise, 28.
	//
	//     Set your computer clock to 2020 to see whether it works.
	//
	//  2. It should work case-insensitive. See below.
	//
	//     Search on Google: golang pkg strings ToLower
	//
	//  3. Get the current year using the time.Now()
	//
	//     Search on Google: golang pkg time now year
	//
	//
	// EXPECTED OUTPUT
	//
	//  -----------------------------------------
	//  Your solution should not accept invalid months
	//  -----------------------------------------
	//  go run main.go
	//    Give me a month name
	//
	//  go run main.go sheep
	//    "sheep" is not a month.
	//
	//  -----------------------------------------
	//  Your solution should handle the leap years
	//  -----------------------------------------
	//  go run main.go january
	//    "january" has 31 days.
	//
	//  go run main.go february
	//    "february" has 28 days.
	//
	//  go run main.go march
	//    "march" has 31 days.
	//
	//  go run main.go april
	//    "april" has 30 days.
	//
	//  go run main.go may
	//    "may" has 31 days.
	//
	//  go run main.go june
	//    "june" has 30 days.
	//
	//  go run main.go july
	//    "july" has 31 days.
	//
	//  go run main.go august
	//    "august" has 31 days.
	//
	//  go run main.go september
	//    "september" has 30 days.
	//
	//  go run main.go october
	//    "october" has 31 days.
	//
	//  go run main.go november
	//    "november" has 30 days.
	//
	//  go run main.go december
	//    "december" has 31 days.
	//
	//  -----------------------------------------
	//  Your solution should be case insensitive
	//  -----------------------------------------
	//  go run main.go DECEMBER
	//    "DECEMBER" has 31 days.
	//
	//  go run main.go dEcEmBeR
	//    "dEcEmBeR" has 31 days.

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Give me a month name")
	// 	return
	// }

	// // Get the current time and find out whether it is a leap year
	// year := time.Now().Year()
	// leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// days := 28
	// month := os.Args[1]

	// // case insensitive
	// m := strings.ToLower(month)
	// if m == "april" || m == "june" || m == "september" || m == "november" {
	// 	days = 30
	// } else if m == "january" || m == "march" || m == "may" || m == "august" || m == "october" || m == "december" {
	// 	days = 31
	// } else if m == "february" && leap {
	// 	if leap {
	// 		days = 29
	// 	}
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// 	return
	// }

	// fmt.Printf("%q has %d days.\n", month, days)
	//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

	//////////////////////////////////////////////////////////////////////////////////////////////////

	//------------------------------------------
	// ERROR HANDLING
	// age := os.Args[1]

	// n, err := strconv.Atoi(age)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Success: Converted %q to %d.\n", age, n)

	// arg := os.Args[1]

	// feet, err := strconv.ParseFloat(arg, 64)
	// if err != nil {
	// 	fmt.Printf(`error: %q is not a number`, arg)
	// 	return
	// }
	// meters := feet * 0.3048
	// fmt.Printf(" %g feet is %g converted to meters\n", feet, meters)

	//-------------------------------------------

	//--------------------------------------------------------------
	// const (
	// 	usage       = "Usage: [username] [password]"
	// 	errUser     = "Access denied for %q.\n"
	// 	errPwd      = "Invalid password for %q.\n"
	// 	accessOK    = "Access granted to %q.\n"
	// 	user, user2 = "isaac", "pee"
	// 	pass, pass2 = "2711", "1127"
	// )

	// args := os.Args

	// if len(args) != 3 {
	// 	fmt.Println(usage)
	// 	return
	// }

	// u, p := args[1], args[2]

	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }
	//----------------------------------------------------------

	//----------------------------------------------------------
	// args := os.Args
	// if len(os.Args) != 3 {
	// 	fmt.Println("Usage: [Username] [Password]")
	// 	return
	// }

	// u, p := args[1], args[2]

	// if u != "Isaac" {
	// 	fmt.Printf("Access denied for %q\n", u)
	// } else if p != "2711" {
	// 	fmt.Printf("Invalid password for %q\n", u)
	// } else {
	// 	fmt.Printf("Access granted to %q\n", u)
	// }
	//---------------------------------------------------------

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
	// 	var (
	// 		args = os.Args
	// 		l = len(args) - 1
	// 	)

	// if l == 0 {
	// 	fmt.Printf("Give me args")
	// } else if l == 1 {
	// 	fmt.Printf("There is one: %q\n", args[1])
	// } else if l == 2 {
	// 	fmt.Printf(
	// 		`There are two: "%s %s"`+"\n",
	// 		args[1], args[2],
	// 	)
	// } else {
	// 	fmt.Printf("There are %d arguments\n", 1)
	// }

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
