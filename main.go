// First we import necessary packages
// "fmt" for formatted I/O functions.
// "log" for logging fatal errors.
// "net/http" for building HTTP servers.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// This function handles form submissions.
// It gets called whenever a client sends a POST request to "/form" URL.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Try to parse form data from the request. If there's an error, report it to the client.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// If there's no error, send a success message.
	fmt.Fprintf(w, "POST request successful\n")
	// Extract form values for "name" and "address".
	name := r.FormValue("name")
	address := r.FormValue("address")
	// And send them back to the client.
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// This function returns a simple "Hello!" message.
// It gets called when a client sends a GET request to "/hello" URL.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// If the request URL path is not "/hello", return a 404 status.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// If the request method is not GET, return an error.
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Otherwise, write a "Hello!" message to the response.
	fmt.Fprintf(w, "Hello!")
}

// The main function where the server starts.
func main() {
	// Create a file server which serves files out of the "./static" directory.
	// When the client makes a request to "/", the server will respond with
	// the file from the "./static" directory that matches the requested URL.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register the form handler function for the "/form" URL.
	http.HandleFunc("/form", formHandler)
	// Register the hello handler function for the "/hello" URL.
	http.HandleFunc("/hello", helloHandler)

	// Print a start-up message and start the server.
	// If the server fails to start, log the error and exit the program.
	fmt.Println("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
