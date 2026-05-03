package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValue(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("The sum of 2 + 2 is %d", sum))
}

// addValue adds two integers and return the sum
func addValue(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

// main is the main function
func main() {
	//fmt.Println("Hello World")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
