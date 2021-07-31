package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Jsn, Home Page")
}

func Page1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"its page 1")
}

func handleRequest(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/page1",Page1)
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func main(){
	handleRequest()
}