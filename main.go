package main

import (
	"fmt"
	"log"
	"net/http"
)

// writting form handler
func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
	fmt.Fprintf(w, "ParseFrom() err: %v", err)
	return
	}
	fmt.Fprintf(w, "Post request succesffuly submitted")
	name := r.FormValue("name")
	address := r.FormValue("address")
	//phone := r.Formvalue = ("phone")

	fmt.Fprintf(w, "Name  = %s\n", name)
	fmt.Fprintf(w, "Adresss  = %s\n", address)
	
	}



// hello handler
func helloHandler(w http.ResponseWriter, r *http.Request ){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	// check the request from the browser
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello there")
}


func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler )

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}
}