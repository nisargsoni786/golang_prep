package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

type HandleFnc func(http.ResponseWriter,*http.Request)   // initialization of HandleFnc

func logPanics(function HandleFnc) HandleFnc {        // function to handle errors
	return func(writer http.ResponseWriter, request *http.Request) {
	  defer func() {
		if x := recover(); x != nil {
		  log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
		}
	  }()
	function(writer, request)
   }
  }

func homePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprintf(w,"Jsn, Home Page")	
	io.WriteString(w,"<h1>Home with io</h1>")
}

func Page1(w http.ResponseWriter, request *http.Request){
	if request.Method=="GET"{
		fmt.Fprintf(w,"its page 1 GETTTTTTTTTT")
	}else{
		fmt.Fprintf(w,"its page 1 POOOSSSSSTTTTTTTTTT")
	}
}

func handleRequest(){
	// http.HandleFunc("/",homePage)   // rather than doing this we can add logpanics function so 
									   // we can handle run time errors 
	// http.HandleFunc("/page1",Page1)

	http.HandleFunc("/",logPanics(homePage))
	http.HandleFunc("/page1",logPanics(Page1))
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
	  log.Fatal("ListenAndServe: ", err.Error())
	}else{
		fmt.Println("Listening at http://localhost:3000")
	}
}

func main(){
	handleRequest()
}