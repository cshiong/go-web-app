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
	"net/url"
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

func even(num int) bool{
	return num % 2 == 0
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
	h.renderTemplate(w,Page{"homepage",p},"index")

}

func (h *MyHandler)renderTemplate(w http.ResponseWriter, p Page, tmpls ...string) {
	log.Println("in renderTemplate()")
	t, err := template.ParseFiles(tmpls ...)
	//need abs filepath
	if err == nil {
		t.Execute(w, p)
	}else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}


func (h *MyHandler)renderTemplateWithFunc(w http.ResponseWriter, p Page, tmpls ...string) {
	log.Println("in renderTemplateWithFunc()")
	/*t, err := template.ParseFiles(tmpls ...)
	//need abs filepath

	//can't add the func after the parse, need have the funMap first

	t.Funcs(template.FuncMap{
		"even": even,
	})*/

	var funcMap = template.FuncMap{
		"even": even,
	}
	path :=strings.Split(tmpls[0],"/")
	baseName := path[len(path)-1]
	fmt.Printf("baseName:%s\n",baseName)
    //New() needs to have the same basename as the "top level" template,but not the filePath just the name.
	t := template.Must(template.New(baseName).Funcs(funcMap).ParseFiles(tmpls ...))

	t.Execute(w, p)


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
	h.renderTemplate(w,Page{"Dynamo Tables",result.TableNames},filepath.Join(h.templatePath,"tables.tmpl"))
}

func (h *MyHandler) GetTablesContents(w http.ResponseWriter, r *http.Request){
	log.Println("in GetTableContents()")

	path :=r.URL.Path
    fmt.Printf("table path:%s\n",path)
	p :=strings.Split(path, "/")
	table :=p[3]
	fmt.Printf("table:%s\n",table)

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
	content := Page{table,jobs}
	fmt.Printf("jobs:%v\n",jobs)

	/*h.renderTemplateWithFunc(w, content, filepath.Join(h.templatePath,"itemsLayout.tmpl"),filepath.Join(h.templatePath,"itemsContent.tmpl"),
		filepath.Join(h.templatePath,"footer.tmpl"))
*/
	h.renderTemplateWithFunc(w, content, filepath.Join(h.templatePath,"layout.tmpl"),filepath.Join(h.templatePath,"LayoutContent.tmpl"),
		filepath.Join(h.templatePath,"footer.tmpl"))

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


func (h *MyHandler) DisplayForm(w http.ResponseWriter, r *http.Request){
	log.Println("in DisplayForm()")

	content := Page{"display form","dummy"}
	h.renderTemplate(w, content, filepath.Join(h.templatePath,"layout.tmpl"),filepath.Join(h.templatePath,"form.tmpl"))

}

// our user context input type
type UserContextInput struct {
	Tag              string
	Script           string
	Style            string
	AttributeLink    string
	AttributeDouble  string
	AttributeSingle  string
	AttributeOnEvent string
	AttributeSrc     string
}


// parses the form values into our custom structure.
func getFormValues(userInput *url.Values) (contextInput *UserContextInput, err error) {
	log.Println("in getFormValues()")

	// create a new UserContextInput struct to hold user input
	contextInput = new(UserContextInput)

	// iterate over the form values assigning each one to our struct.
	for key, value := range *userInput {
		switch key {
		case "tag":
			contextInput.Tag = value[0]
		case "script":
			contextInput.Script = value[0]
		case "style":
			contextInput.Style = value[0]
		case "attr_double":
			contextInput.AttributeDouble = value[0]
		case "attr_single":
			contextInput.AttributeSingle = value[0]
		case "attr_onevent":
			contextInput.AttributeOnEvent = value[0]
		case "attr_src":
			contextInput.AttributeSrc = value[0]
		default:
			return nil, fmt.Errorf("Unknown value passed in form!") // no funny business.
		}
	}
	return contextInput, nil // yes, it is safe to return a pointer in Go
}

func (h *MyHandler) GetFormResult(w http.ResponseWriter, r *http.Request) {
	log.Println("in GetFormResult()")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form values!", http.StatusInternalServerError)
		return
	}
	// get the form values
	userInput, err := getFormValues(&r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	content := Page{"display form", userInput}

	h.renderTemplate(w, content, filepath.Join(h.templatePath, "layout.tmpl"), filepath.Join(h.templatePath, "formResult.tmpl"))
}