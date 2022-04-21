package main

import (
	"api-http-mock-testing/api"
	"fmt"
	"net/http"
)

func main() {

	service := api.Service{
		HttpClient: http.DefaultClient,
	}

	fmt.Println("listening....")
	http.HandleFunc("/", service.GetRandomUser)

	panic(http.ListenAndServe("localhost:7000", nil))
}
