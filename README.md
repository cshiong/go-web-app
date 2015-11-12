# go-web-app

a full stack web app implemented with goLang; mainly focus on how to use the Go standard template lib.
Templates are just plain text with additional markup. All the markup is stored between {{ and }} tags.
The data that we feed to a template is called dot. We refer to it in our tags with . (hence the name).
 Dot can be pretty much anything: a simple value, a slice, a structure, and so on. 
 If dot is a structure, you can use standard Go variable syntax to access each component of dot.
 Think of templates as if they were Unix pipes: you feed some data into the template, and it produces a text document or HTML document based on that data.

#Goal

1. render static pages.
2. templates with simple data.
3. template with list of data.(slice, map etc)
3. template contains another template.
4. template with customer functions.
5. template with script.

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
  
  