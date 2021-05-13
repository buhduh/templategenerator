package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"templategenerator/document"
)

type AjaxRequest struct {
	Action string            `json:action`
	Data   map[string]string `json:data`
}

type DocumentRequest struct {
	contents string
	name     string
}

func (s *Server) ajaxHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("error reading body with error: '%s'", err),
			http.StatusBadRequest,
		)
		return
	}
	ajaxRequest := AjaxRequest{}
	err = json.Unmarshal(body, &ajaxRequest)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("could not unmarshal request into an AjaxRequest struct with error: '%s'", err),
			http.StatusBadRequest,
		)
		return
	}
	s.myLogger.Debugf("ajax request for %s", ajaxRequest.Action)
	switch ajaxRequest.Action {
	case "parseYAML":
		var doc, name string
		var ok bool
		if doc, ok = ajaxRequest.Data["contents"]; !ok {
			msg := "request body does not contain a contents key for newDocument"
			s.myLogger.Error(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		if name, ok = ajaxRequest.Data["name"]; !ok {
			msg := "request body does not contain a name key for newDocument"
			s.myLogger.Error(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		s.myLogger.Debugf("ajax request for newDocument, name: '%s'", name)
		newDoc, err := s.documentHelper.NewDocument(name, doc)
		if err != nil {
			msg := fmt.Sprintf("failed generating a new document with error, '%s'", err)
			s.myLogger.Error(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		//TODO, this is hardcoded for now
		treeJSON, err := newDoc.DocumentTree.ToJSON()
		if err != nil {
			msg := fmt.Sprintf("failed converting document tree to josn with error, '%s'", err)
			s.myLogger.Error(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		if n, err := w.Write(treeJSON); n != len(treeJSON) || err != nil {
			var msg string
			if n != len(treeJSON) {
				msg = fmt.Sprintf(
					"failed writing document tree's json, only wrote %d bytes of %d total",
					n, len(treeJSON),
				)
			} else {
				msg = fmt.Sprintf("writing response json failed with error, '%s'", err)
			}
			s.myLogger.Errorf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		return
	default:
		msg := fmt.Sprintf("unrecognized action, '%s'", ajaxRequest.Action)
		s.myLogger.Error(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}
}
