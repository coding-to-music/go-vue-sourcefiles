package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println(Red("Function main.go listening on port 3000"))
	fmt.Println(Info("Hello World at the command line from server.go"))
	fmt.Println(Info("Invoke with go run main.go"))

	fmt.Println(Info("hello, Info world!"))
	fmt.Println(Warn("hello, Warn world!"))
	fmt.Println(Fata("hello, Fatal world!"))

	fmt.Println(Black("hello, Black world!"))
	fmt.Println(Red("hello, Red world!"))
	fmt.Println(Green("hello, Green world!"))
	fmt.Println(Yellow("hello, Yellow world!"))
	fmt.Println(Purple("hello, Purple world!"))
	fmt.Println(Magenta("hello, Magenta world!"))
	fmt.Println(White("hello, White world!"))
	fmt.Println(Teal("hello, Teal world!"))
	fmt.Println(Green("hello, Green need Blue world!"))

	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
