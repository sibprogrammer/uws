package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	port := "8080"
	fmt.Println("Static server started:", "http://127.0.0.1:"+port)
	return http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
}
