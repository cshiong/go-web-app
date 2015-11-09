package main


import (

	"net/http"
	"code.comcast.com/VariousArtists/go-web-app/controller"
)


func main() {


	 myHandler := controller.NewMyHandler()
	//load static files

    http.HandleFunc("/",myHandler.LoadStaticPage1)
	//http.HandleFunc("/", myHandler.LoadHomePage)
	http.HandleFunc("/dynamo/tables", myHandler.GetTables)
	http.HandleFunc("/dynamo/tables/",myHandler.GetTablesContents)
	http.ListenAndServe(":8080", nil)
}