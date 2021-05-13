package document

import (
	//"fmt"
	logger "github.com/buhduh/go-logger"
	//"gopkg.in/yaml.v3"
)

type DocumentHelper struct {
	logger logger.Logger
}

func NewDocumentHelper(logger logger.Logger) *DocumentHelper {
	return &DocumentHelper{
		logger: logger,
	}
}

type Document struct {
	Name         string
	Contents     *string
	DocumentTree *Tree
}

func (d *DocumentHelper) NewDocument(name, rawDoc string) (*Document, error) {
	d.logger.Debugf("creating a new document: '%s'", name)
	return &Document{
		Name:         name,
		Contents:     &rawDoc,
		DocumentTree: hardCodedTree(),
	}, nil
}
