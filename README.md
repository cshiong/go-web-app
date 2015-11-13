# go-web-app

a full stack web app implemented with goLang; 
---
mainly focus on how to use the Go standard template lib.(text/template and html/template).
A truly MVC pattern.
Templates are just plain text with additional markup. All the markup is stored between {{ and }} tags.
The data that we feed to a template is called dot. We refer to it in our tags with . (hence the name).
 Dot can be pretty much anything: a simple value, a slice, a structure, and so on. 
 If dot is a structure, you can use standard Go variable syntax to access each component of dot.also it contextually aware(see items.tmpl for e.g)
 Think of templates as if they were Unix pipes: you feed some data into the template, and it produces a text document or HTML document based on that data.
 html/tempalte

#Goal

1. render static pages.
2. templates with simple data.
3. template with complex of data.(slice, map etc.)
3. template contains another template.
4. template with builtin and customer functions.
5. template with script.(support 3rd part lib etc.)
6. template integrated with Bootstrap.


#Template Files

1.tables.tmpl
 shows the .tag value is pipeline and build in functions. e.g 
 `````
 {{printf "%s" .Title | len}}
 `````

2.items.tmpl
  shows the simply way to work on list of data.
  `````
  {{range .}}
        <tr>
            <td>{{.Shard}}</td>
            <td>{{.Id}}</td>
            <td>{{.Job}}</td>
            <td>{{.Tag}}</td>
            <td>{{.Time}}</td>
       </tr>
   {{end}}
  `````
  

3.layout.tmpl
 as a main layout which take template "content" as content for display in the main div.
 `````
  {{template "content" .}}
 
  `````
 

4.layoutContent.tmpl, form.tmpl, formresult.tmpl,chart.tmpl
All kinds contents can be passed to the layout.tmpl to display different contents.
form.tmpl and formresult.tmpl also show how the script works in template
`````
{{define "content"}}

`````


5.itemsLayout.tmpl and itemsContent.tmpl 
Used to show how customer function works on template.
`````
"{{if even $index}}{{else}}alt{{end}}"

`````


6.chart.tmpl
integrated with Google chart lib in the script to display a sample pie chart.
`````
 <script type="text/javascript" src="https://www.google.com/jsapi"></script>
 
`````

#Third part information
Bootstrap --- http://www.w3schools.com/bootstrap/default.asp


