package handlers

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorln("unable upgrade to websocket", r.URL)
	}
	defer ws.Close()

}

func GetTasksRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handleConnections)

	return router
}
