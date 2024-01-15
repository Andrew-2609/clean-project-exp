package webserver

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	Router        *http.ServeMux
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        http.NewServeMux(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
}

func (ws *WebServer) Start() {
	for path, handler := range ws.Handlers {
		ws.Router.HandleFunc(path, handler)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", ws.WebServerPort), ws.Router); err != nil {
		panic(err)
	}
}
