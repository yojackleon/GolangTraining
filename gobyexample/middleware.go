/*
Functions in Go are first-class citizens. What this means is that you can not only
create anonymous functions dynamically, but you can also pass functions as parameters
to a function. For example, when creating a web server it is common to provide a
function that processes a web request to a specific route.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", timed(hello))
	http.ListenAndServe(":3000", nil)

}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")

}
