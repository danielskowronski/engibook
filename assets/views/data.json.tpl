{ 
	"notebooks": [
		{}{{range $k,$v:=.NotebooksList}}
		,{"id": {{$v.ID}}, "title": "{{$v.Title}}" }{{end}}
	],
	"notes": [
		{}{{range $k,$v:=.List}}
		,{"id": {{$v.ID}}, "notebook": {{$v.NotebookID}}, "title": "{{$v.Title}}", "body": "{{$v.Body}}" } {{end}}
	]
}