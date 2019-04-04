package elastic

type elasticsearchMappingProperty struct {
	Name string
	Type string
	Last bool
}

var elasticsearchTemplate = `
{
  "version": 1,
  "index_patterns": ["raccoon-events*"],
  "settings": {
    "number_of_shards": 5
  },
  "mappings": {
    "_doc": {
      "properties": {
	{{- range .}}
	"{{.Name}}": { "type": "{{.Type}}" }{{if not .Last}},{{end}}
	{{- end}}
      }
    }
  }
}
`
