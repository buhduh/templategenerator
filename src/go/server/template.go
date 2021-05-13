package server

import (
	"html/template"
	"io"
	"io/fs"
	"templategenerator/constants"
)

type Template struct {
	name     string
	template *template.Template
	data     interface{}
}

type GlobalTemplateData struct {
	Title        string
	ResourceRoot string
}

func getGlobalTemplateData() GlobalTemplateData {
	return GlobalTemplateData{
		Title:        constants.APPLICATION,
		ResourceRoot: constants.RES_ROOT,
	}
}

func NewTemplate(fs fs.FS, fileName string) (*Template, error) {
	template := template.New(fileName)
	tmp, err := template.ParseFS(fs, fileName)
	if err != nil {
		return nil, err
	}
	return &Template{
		name:     fileName,
		template: tmp,
	}, nil
}

//TODO should i cache this? probably not for local dev
func (t *Template) Execute(w io.Writer, templateFS fs.FS, data interface{}) error {
	/*
		headerF, err := templateFS.Open(constants.DEF_HEADER_TEMPLATE)
		if err != nil {
			return err
		}
		headerB, err := ioutil.ReadAll(headerF)
		if err != nil {
			return err
		}
	*/
	var exeData = struct {
		G GlobalTemplateData
		D interface{}
		//Header template.HTML
	}{
		G: getGlobalTemplateData(),
		D: data,
		//Header: template.HTML(headerB),
	}
	return t.template.Execute(w, exeData)
}
