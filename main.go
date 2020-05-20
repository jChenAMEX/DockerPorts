package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainfun(w http.ResponseWriter, r *http.Request) {
	log.Println("main server running on :8090")
	fmt.Fprintf(w, "running from main function")
}
func funone(w http.ResponseWriter, r *http.Request) {
	log.Println(" server one running on :8091")
	fmt.Fprintf(w, "running from funone function")
}
func funtwo(w http.ResponseWriter, r *http.Request) {
	log.Println(" server two running on :8092")
	fmt.Fprintf(w, "running from funtwo function")
}
func main() {
	fmt.Println("Docker MultiPorts")
	server := http.NewServeMux()
	server.HandleFunc("/", mainfun)

	serverone := http.NewServeMux()
	serverone.HandleFunc("/", funone)
	go func() {
		http.ListenAndServe(":8091", serverone)
	}()

	servertwo := http.NewServeMux()
	servertwo.HandleFunc("/", funtwo)
	go func() {
		http.ListenAndServe(":8092", servertwo)
	}()

	http.ListenAndServe(":8090", server)

}
