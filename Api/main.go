package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct{
	Msg string `json:"message"`
}

func apiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	resp := Response{Msg: "Welcome to Api Powered By Abhay"}
	json.NewEncoder(w).Encode(resp)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Abhay")
}

func main(){
	http.HandleFunc("/",helloHandler)
	http.HandleFunc("/api",apiHandler)
	http.ListenAndServe(":8080",nil)
}