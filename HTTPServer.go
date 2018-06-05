package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/adhi",handler2)
	http.ListenAndServe(":8000",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Adhi \n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Satria \n")
}