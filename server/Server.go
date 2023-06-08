package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature list goest here")
	_, err := writer.Write([]byte(placeholder))
	check(err)
	writer.WriteHeader(http.StatusOK)
}
func getHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature get list")
	_, err := writer.Write([]byte(placeholder))
	check(err)
	writer.WriteHeader(http.StatusOK)
}
func postHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature post in list")
	_, err := writer.Write([]byte(placeholder))
	check(err)
	writer.WriteHeader(http.StatusOK)
}
func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}
