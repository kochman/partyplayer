package websocket

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var errNotABinaryMessage = errors.New("Not a binary message")
var ErrConnClosed = errors.New("Connection is closed")

// Server is a WebSocket server.
type Server struct {
	lock           sync.Mutex
	connections    map[string]*Conn
	newConnHandler NewConnHandler
}

// NewServer creates and returns a new WebSocket server.
func NewServer(newConnHandler NewConnHandler) (*Server, error) {
	s := &Server{
		connections:    make(map[string]*Conn),
		newConnHandler: newConnHandler,
	}
	return s, nil
}

func (s *Server) newConnection(key string, wsConn *websocket.Conn) {
	c := newConnection(key, wsConn)
	s.lock.Lock()
	s.connections[key] = c
	s.lock.Unlock()

	// Call the new connection handler
	if s.newConnHandler != nil {
		s.newConnHandler(c)
	}
}

func (s *Server) Connection(key string) *Conn {
	s.lock.Lock()
	conn := s.connections[key]
	s.lock.Unlock()
	return conn
}

func (s *Server) Connections() []*Conn {
	conns := []*Conn{}
	s.lock.Lock()
	for _, conn := range s.connections {
		conns = append(conns, conn)
	}
	s.lock.Unlock()
	return conns
}

// ServeHTTP implements the http.Handler interface so that it can be placed
// inside an http.ServeMux or other router.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Unable to upgrade WebSocket connection:", err)
		return
	}

	u := uuid.NewV4()
	s.newConnection(u.String(), c)
}
