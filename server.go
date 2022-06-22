package main
import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/hello"{
		http.Error(res,"404 NOT FOUND!",http.StatusNotFound)
		return
	}
	if req.Method !="GET" {
		http.Error(res,"THIS METHOD NOT SUPPORTED",http.StatusNotFound)
		return
	}
	fmt.Fprint(res,"HELLO")
}
func formHandler(res http.ResponseWriter, req *http.Request){
	if err :=req.ParseForm(); err!=nil {
		fmt.Fprintf(res,"Parse form error, %v",err)
	}
	fmt.Fprintf(res,"POST request successful \n")
	username := req.FormValue("username")
	email := req.FormValue("email")
	fmt.Fprintf(res,"Username :%s \n",username)
	fmt.Fprintf(res,"Email :%s \n",email)
}
func main(){
	http.Handle("/",http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting the server at port 8080\n")
	if err :=http.ListenAndServe(":8080", nil); err!=nil {
		log.Fatal(err)
	}
}