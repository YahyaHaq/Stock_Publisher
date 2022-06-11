package api

import (
	"encoding/json"
	"net"
	"time"

	"github.com/BackendTest/util"
)

type Server struct {
	listner    net.Listener
	clientList []net.Conn
}

// Takes in a config object to create a new Server
func NewServer(config util.Config) (*Server, error) {
	conn, err := net.Listen(config.ConnectionType, config.ServerAddress)
	if err != nil {
		return nil, err
	}

	clientList := make([]net.Conn, 0, 10)

	server := &Server{
		conn,
		clientList,
	}

	return server, nil

}

// starts the server
func (server *Server) StartServer() {
	go func() {
		for {
			conn, err := server.listner.Accept()
			if err != nil {
				return
			}
			server.clientList = append(server.clientList, conn)

		}
	}()

}

// notifies all the clients in the servers client list
func (server *Server) notifyAll() {
	for _, client := range server.clientList {
		random_number := util.RandomInt(1, 10) - 1

		tick := Cache[random_number]
		tick.UpdateTick()

		jsonTick, err := json.Marshal(tick)
		if err != nil {
			return
		}

		jsonTick = append(jsonTick, byte('\n'))
		client.Write([]byte(jsonTick))
	}
}

// publishes the tick to all the clients with a time delay of 100 milisceonds
func (server *Server) PublishTick() {
	for {
		server.notifyAll()
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
