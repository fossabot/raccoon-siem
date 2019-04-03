package main

import (
	"flag"
	"github.com/tephrocactus/raccoon-siem/sdk/normalization"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type eventFieldMeta struct {
	Name string
	Kind string
	Set  bool
	Time bool
}

var packageImportTemplate = `package normalization

// 
// THIS FILE IS GENERATED. DO NOT EDIT!
//

import (
	"strings"
	"gopkg.in/vmihailenco/msgpack.v4"
	"github.com/francoispqt/gojay"
)
`

func main() {
	outPath := flag.String("out", "", "")
	flag.Parse()

	_ = os.Remove(*outPath)

	outFile, err := os.OpenFile(*outPath, os.O_CREATE|os.O_RDWR, 0655)
	if err != nil {
		panic(err)
	}

	meta := scanEvent()
	templates := []string{
		packageImportTemplate,
		gettersTemplate,
		settersTemplate,
		encodeMsgpackTemplate,
		decodeMsgpackTemplate,
		encodeJSONTemplate,
		decodeJSONTemplate,
	}

	for _, t := range templates {
		tpl, err := template.New("").Parse(t)
		if err != nil {
			panic(err)
		}

		if err := tpl.Execute(outFile, meta); err != nil {
			panic(err)
		}
	}
}

func scanEvent() (result []eventFieldMeta) {
	e := normalization.Event{}
	rt := reflect.TypeOf(e)

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)

		isLabel := strings.Index(f.Name, "User") == 0 && strings.Index(f.Name, "Label") > -1
		isTimestamp := strings.Index(f.Name, "Timestamp") > -1

		result = append(result, eventFieldMeta{
			Name: f.Name,
			Kind: f.Type.Name(),
			Set:  f.Tag.Get("set") != "",
			Time: isTimestamp && !isLabel,
		})
	}

	return
}
