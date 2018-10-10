package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func trinixHandler(w http.ResponseWriter, r * http.Request){
	fmt.Println("/trinix")
	fmt.Fprintf(w, "Hi Iam Trinix Systemssas %s!", r.URL.Path[1:])	
}
func utHandler(w http.ResponseWriter, r * http.Request){
	fmt.Println("/ut")
	fmt.Fprintf(w, "Hi Iam UT %s!", r.URL.Path[1:])	
}
func utHandler(w http.ResponseWriter, r * http.Request){
	fmt.Println("/lawencon")
	fmt.Fprintf(w, "Hi Iam Lawencon  your path is %s!", r.URL.Path[1:])	
}

func main() {
	fmt.Println("Initialization ...")
    http.HandleFunc("/", handler)
    http.HandleFunc("/trinix", trinixHandler)
    http.HandleFunc("/ut", utHandler)
    fmt.Println("Listen at port 9000")
    log.Fatal(http.ListenAndServe(":9000", nil))

}
