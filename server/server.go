package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kochman/partyplayer/config"
	"github.com/kochman/partyplayer/controller"
	"github.com/kochman/partyplayer/log"
	"github.com/kochman/partyplayer/websocket"
)

type Server struct {
	cfg        config.Config
	controller controller.Controller
	handler    http.Handler
}

func New(cfg *config.Config, controller *controller.Controller) (*Server, error) {
	s := &Server{
		cfg:        *cfg,
		controller: *controller,
	}

	wsServer, err := websocket.NewServer(s.newConnHandler)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Mount("/ws", wsServer)

	s.handler = router
	return s, nil
}

func (s *Server) Run() {
	if err := http.ListenAndServe(s.cfg.ListenURL, s.handler); err != nil {
		log.WithError(err).Error("Unable to serve.")
	}
}

type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (s *Server) newConnHandler(conn *websocket.Conn) {
	log.Debugf("New connection from %s", conn.ClientAddr)

	playlists := s.controller.Playlists()
	data, err := json.Marshal(playlists)
	if err != nil {
		log.WithError(err).Error("Unable to marshal playlists.")
		return
	}

	err = conn.Send(data)
	if err != nil {
		log.WithError(err).Error("Unable to send.")
		return
	}

	go func() {
		for incomingMsg := range conn.ReceiveChan {
			log.Infof("got message: [%s]", string(incomingMsg))
		}
	}()
}
