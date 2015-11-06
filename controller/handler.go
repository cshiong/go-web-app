package controller

import (


	"net/http"
/*	"log"
"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"*/
"html/template"
	"log"
	"path/filepath"
	"os"
)



type MyHandler struct {
	templatePath string
}

func NewMyHandler() *MyHandler{
	cwd, _ := os.Getwd()
	var h MyHandler
	h.templatePath =filepath.Join(cwd,"/www/templates/")
	return &h
}

func (h *MyHandler) LoadHomePage(w http.ResponseWriter, r *http.Request){
    log.Println("in LoadHomePage()")
	p :="Home page of monitor"
	h.renderTemplate(w,"index",p)

}

func (h *MyHandler)renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	log.Println(h.templatePath)
	t, err := template.ParseFiles(filepath.Join(h.templatePath, tmpl + ".html"))
	//need abs filepath
	if err == nil {
		t.Execute(w, p)
	}else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

/*
func (h *MyHandler) GetTables(rw http.ResponseWriter, r *http.Request){

	log.Println()
	region :=r.PathParams["region"]

	svc := dynamodb.New(session.New((&aws.Config{Region: aws.String(region)})))
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("error:%s\n",err)
		return
	}

}*/
