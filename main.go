package main

import{
	"fmt"
	"log"
	"net/http"
}




func formHandler(w http.ResponseWriter, r *http.Request){
		if err := r.ParseForm(); err != nil {

		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
		}


fmt.Fprintf(w, "POST request successful")
namer.FormValue("name")
address = r.FormValue("address") 
fmt.Fprintf(w, "Name = %s\n", name) 
fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler (w http. Responsewriter, r *http.Request){ 
	if r.URL.Path = "/hello" {
	http.Error(w, "484 not found", http.StatusNotFound)
	return
}
if r.Method != "GET" {
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err :=http.ListenAndServer(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}