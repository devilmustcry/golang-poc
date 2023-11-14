package controller

import "net/http"

func HandleHello(writer http.ResponseWriter, request *http.Request) {
	res, err := http.Get("http://localhost:9090/hello")
	if err != nil {
		panic(err)
	}
	res.Write(writer)
}
