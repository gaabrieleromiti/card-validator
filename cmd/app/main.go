package main

import (
	"ccv/pkg"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", pkg.LunhHandler)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
		fmt.Println(err)
	}


}
