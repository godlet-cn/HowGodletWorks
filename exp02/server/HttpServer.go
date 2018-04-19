package server

import (
	"fmt"
	"github.com/godlet-cn/HowGodletWorks/exp02/processor"
	"github.com/godlet-cn/HowGodletWorks/exp02/request"
	"github.com/godlet-cn/HowGodletWorks/exp02/response"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const SHUTDOWN_COMMAND string = "/SHUTDOWN"

//HttpServer is working to serve client connection
type HttpServer struct {
	port     int
	listener net.Listener
	shutdown bool
}

func (httpServer *HttpServer) init() {
	httpServer.shutdown = false
}

//Await start waiting for client connection
func (httpServer *HttpServer) Await() {
	httpServer.port = 8056
	addr := "localhost:8056"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	httpServer.listener = listener

	defer httpServer.Close()

	for {
		//call listener.Accept() will hang up server process to wait for a new client
		conn, err := httpServer.listener.Accept()

		if err != nil {
			log.Fatal(err)
			break
		}

		go httpServer.HandleConnection(conn)
	}
}

//Close function will shutdown HttpServer
func (httpServer *HttpServer) Close() {
	if httpServer.listener != nil {
		fmt.Println(time.Now().Local(), "Server is shutting down")
		httpServer.listener.Close()
	}
}

//HandleConnection handle a client connection
func (httpServer *HttpServer) HandleConnection(conn net.Conn) {

	//conn will be closed finally whatever this function does
	defer conn.Close()

	req := request.NewRequest(conn)
	req.Parse()

	resp := response.NewResponse(conn)
	resp.SetRequest(req)

	if strings.HasPrefix(req.Uri, "/servlet/") {
		var processor processor.ServletProcessor1
		processor.Process(req, resp)
	} else {
		var processor processor.StaticResourceProcessor
		processor.Process(req, resp)
	}

	httpServer.shutdown = req.Uri == SHUTDOWN_COMMAND
	if httpServer.shutdown {
		httpServer.Close()
	}
}
