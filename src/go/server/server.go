package server

import (
	"fmt"
	logger "github.com/buhduh/go-logger"
	"io/fs"
	"net/http"
	"templategenerator/constants"
	"templategenerator/document"
)

type Server struct {
	myLogger       logger.Logger
	port           string
	templateFS     fs.FS
	documentHelper *document.DocumentHelper
}

func (s *Server) mainHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := NewTemplate(s.templateFS, "root.html.tpl")
	if err != nil {
		http.Error(
			w, fmt.Sprintf("error parsing template: '%s'", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err = tpl.Execute(w, s.templateFS, nil); err != nil {
		http.Error(
			w, fmt.Sprintf("error parsing template: '%s'", err),
			http.StatusInternalServerError,
		)
		return
	}
}

func (s *Server) spikeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := NewTemplate(s.templateFS, "spike.html.tpl")
	if err != nil {
		http.Error(
			w, fmt.Sprintf("error parsing template: '%s'", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err = tpl.Execute(w, s.templateFS, nil); err != nil {
		http.Error(
			w, fmt.Sprintf("error parsing template: '%s'", err),
			http.StatusInternalServerError,
		)
		return
	}
}

func NewServer(
	port string, templateFS fs.FS,
	myLogger logger.Logger, docHelper *document.DocumentHelper,
) *Server {
	toRet := &Server{
		port:           port,
		templateFS:     templateFS,
		myLogger:       myLogger,
		documentHelper: docHelper,
	}
	http.HandleFunc("/", toRet.mainHandler)
	myLogger.Debugf("initialized root handler at '/'")
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir(constants.RES_ROOT))))
	myLogger.Debugf("initialized resource path at: '%s'", constants.RES_ROOT)

	http.HandleFunc("/spike", toRet.spikeHandler)
	myLogger.Debugf("initialized spike handler")

	http.HandleFunc("/ajax", toRet.ajaxHandler)
	myLogger.Debugf("initialized ajax handler")

	return toRet
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}
