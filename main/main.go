package main


import (

	"net/http"
	"code.comcast.com/VariousArtists/go-web-app/controller"
)


func main() {


	 myHandler := controller.NewMyHandler()

	http.HandleFunc("/", myHandler.LoadHomePage)
	http.ListenAndServe(":8080", nil)
}