package main


import (

	"net/http"
	"code.comcast.com/VariousArtists/go-web-app/controller"
)


func main() {


	 myHandler := controller.NewMyHandler()
	//load static files

    http.HandleFunc("/",myHandler.LoadStaticPage1)
	//http.HandleFunc("/", myHandler.LoadStaticPage)
	http.HandleFunc("/dynamo/tables", myHandler.GetTables)
	http.HandleFunc("/dynamo/tables/",myHandler.GetTablesContents)
	http.HandleFunc("/dynamo/form", myHandler.DisplayForm)
	http.HandleFunc("/dynamo/formresult", myHandler.GetFormResult)
	http.HandleFunc("/dynamo/chart", myHandler.DisplayChart)


	http.ListenAndServe(":8080", nil)
}