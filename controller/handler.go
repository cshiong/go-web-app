package controller

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"

	"net/http"
	"log"
	"html/template"
"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"

//	"code.comcast.com/VariousArtists/cronaas/core"
	"fmt"
)



type MyHandler struct {
	templatePath string
	staticDir string
}

type Page struct{
	Title string
	Content interface{}
}

type CronaasJob struct{
	Shard string
	Id string
	Job string
	Tag string
	Time string
}

func NewMyHandler() *MyHandler{
	cwd, _ := os.Getwd()
	var h MyHandler
	h.templatePath =filepath.Join(cwd,"/www/templates/")
	h.staticDir ="www"
	return &h
}
/*
load static pages using ReadFile
 */
func (h *MyHandler) LoadStaticPage1(w http.ResponseWriter, r *http.Request){
	log.Println("in LoadStaticPage1()")
	path := r.URL.Path[1:]
	var filePath string
	if strings.Contains(path,"www") {
		filePath = path
	}else{
		filePath = filepath.Join("www", path)
	}
	//only allow go to www folder.
	log.Printf("url path:%s\n",filePath)
	data, err := ioutil.ReadFile(filePath)
	if err == nil {
		var contentType string
		switch {
		case strings.HasSuffix(filePath, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(filePath, ".html"):
			contentType ="text/html"
		default:
			contentType ="test/plain"
		}
		w.Header().Add("Content Type",contentType)
		w.Write(data)
	}else{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


}
/*
load static pages using fileServer
 */
func (h *MyHandler) LoadStaticPage(w http.ResponseWriter, r *http.Request){
	log.Println("in LoadStaticPage()")
	fs := http.FileServer(http.Dir(h.staticDir))

	http.Handle("/v1/", http.StripPrefix("/www/", fs))


}

func (h *MyHandler) LoadHomePage(w http.ResponseWriter, r *http.Request){
    log.Println("in LoadHomePage()")
	p :="VA Administor Tool"
	h.renderTemplate(w,"index",p)

}

func (h *MyHandler)renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	log.Println("in renderTemplate()")
	log.Println(h.templatePath)
	t, err := template.ParseFiles(filepath.Join(h.templatePath, tmpl + ".tmpl"))
	//need abs filepath
	if err == nil {
		t.Execute(w, p)
	}else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}


func (h *MyHandler) GetTables(w http.ResponseWriter, r *http.Request){

	log.Println("in GetTables()")
	region :="us-west-2"

	svc := dynamodb.New(session.New((&aws.Config{Region: aws.String(region)})))
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("error:%s\n",err)
		return
	}
	h.renderTemplate(w,"tables",result.TableNames)
}

func (h *MyHandler) GetTablesContents(w http.ResponseWriter, r *http.Request){
	log.Println("in GetTableContents()")

	job1:= CronaasJob {"22",
		"2643c528-b784-4faa-90a9-3f665132d26c",
		`{"ID":"2643c528-b784-4faa-90a9-3f665132d26c","schedule":"* * * * *","tag":"ggl","url":"http://localhost:5000/test","body":"","method":"POST","header":nil,"contentType":"application/json","options":{"op2":"opt2","opp1":"op1v"},"randomMins":0,"ExecuteAt":"2015-11-05T02:15:00Z","OriginalExecuteAt":"2015-11-05T02:15:00Z","Retry":0,"Owner":"","Policy":"","Disabled":false,"Recurring":true}`,
		"ggl",
		"2015-11-05T02:15:00Z",
	}
	job2:= CronaasJob {"1",
		"22243g8-b7e2-4faa-90a9-3f665132d26c",
		`{"ID":"2643c528-b784-4faa-90a9-3f665132d26c","schedule":"* * * * *","tag":"ggl","url":"http://localhost:5000/test","body":"","method":"POST","header":nil,"contentType":"application/json","options":{"op2":"opt2","opp1":"op1v"},"randomMins":0,"ExecuteAt":"2015-11-05T02:15:00Z","OriginalExecuteAt":"2015-11-05T02:15:00Z","Retry":0,"Owner":"","Policy":"","Disabled":false,"Recurring":true}`,
		"ggl",
		"2015-11-05T02:15:00Z",
	}
	var jobs []CronaasJob
	jobs = append(jobs,job1,job2)
	fmt.Printf("jobs:%v\n",jobs)
	h.renderTemplate(w,"items",jobs)

/*
	region :="us-west-2"
	svc := dynamodb.New(session.New((&aws.Config{Region: aws.String(region)})))

	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{ // Required
			"cronaas-qa": { // Required
				Keys: []map[string]*dynamodb.AttributeValue{ // Required
					{ // Required
						"Key": { // Required
							S:    aws.String("StringAttributeValue"),
							SS: []*string{
								aws.String("StringAttributeValue"), // Required
								// More values...
							},
						},
						// More values...
					},
					// More values...
				},
				ConsistentRead: aws.Bool(true),
				ExpressionAttributeNames: map[string]*string{
					"Key": aws.String("AttributeName"), // Required
					aws.String("shard"), // Required
					aws.String("id"), // Required
					aws.String("job"), // Required
					aws.String("tag"), // Required
					aws.String("time"), // Required.
				},
				ProjectionExpression: aws.String("ProjectionExpression"),
			},
			// More values...
		},
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
	}
	resp, err := svc.BatchGetItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}
*/

	//result, err := svc.BatchGetItemPages( )

}