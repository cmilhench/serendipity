//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const blob = "assets.gen.go"

var tmpl = template.Must(template.New("").Funcs(map[string]interface{}{"conv": formatByteSlice}).Parse(`package asset
func init(){
	{{- range $name, $file := . }}
	Add("{{ $name }}", []byte{ {{ conv $file }} })
	{{- end }}
}
`))

func main() {
	fmt.Println("Generating embeded assets")
	prefix := "../../assets/data"
	resources := make(map[string][]byte)
	err := filepath.Walk(prefix, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if path.Base(p)[:1] == "." {
			return nil
		}
		r := filepath.ToSlash(strings.TrimPrefix(p, prefix))
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		resources[r] = b
		return nil
	})
	if err != nil {
		panic(err)
	}
	f, err := os.Create(blob)
	if err != nil {
		log.Fatal("Error creating blob file:", err)
	}
	defer f.Close()

	builder := &bytes.Buffer{}

	err = tmpl.Execute(builder, resources)
	if err != nil {
		log.Fatal("Error executing template", err)
	}

	data, err := format.Source(builder.Bytes())
	if err != nil {
		log.Fatal("Error formatting generated code", err)
	}
	err = ioutil.WriteFile(blob, data, os.ModePerm)
	if err != nil {
		log.Fatal("Error writing blob file", err)
	}
}

func formatByteSlice(sl []byte) string {
	builder := strings.Builder{}
	for _, v := range sl {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}
	return builder.String()
}
