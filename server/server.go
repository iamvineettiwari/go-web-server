package server

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"

	"github.com/iamvineettiwari/go-web-server/http"
	"github.com/iamvineettiwari/go-web-server/http/status"
)

type HttpServer struct {
	ListenAddr string
	Listener   net.Listener
	serverLock chan struct{}
}

func NewHttpServer(address string) *HttpServer {
	return &HttpServer{
		ListenAddr: address,
		serverLock: make(chan struct{}),
	}
}

func (hs *HttpServer) Start() error {
	listener, err := net.Listen("tcp", hs.ListenAddr)

	if err != nil {
		return err
	}

	hs.Listener = listener

	go hs.acceptConnection()

	<-hs.serverLock
	return nil
}

func (hs *HttpServer) acceptConnection() {
	for {
		conn, err := hs.Listener.Accept()

		if err != nil {
			log.Println("Error while accepting connection : ", err)
			continue
		}

		go hs.serve(conn)
	}
}

func (hs *HttpServer) closeConnection(conn net.Conn) {
	conn.Close()
}

func (hs *HttpServer) serve(conn net.Conn) {
	defer hs.closeConnection(conn)

	buffer := make([]byte, 6048)

	n, err := conn.Read(buffer)

	if err != nil {
		log.Println("Error occured while reading data : ", err)
		return
	}

	data := buffer[:n]
	request := http.Decode(data)

	response, respStatus, err := hs.processRequest(request)
	var respData []byte

	if err != nil {
		respData = http.Encode(response, respStatus)
	} else {
		respData = http.Encode(response, respStatus)
	}

	conn.Write(respData)
}

func (hs *HttpServer) processRequest(request *http.Request) ([]byte, status.HttpStatus, error) {
	curDirectory, err := os.Getwd()

	if err != nil {
		return []byte("Something went wrong"), status.HTTP_500_SERVER_ERROR, err
	}

	path := curDirectory + "/www" + request.Path

	if request.Path == "/" {
		path += "index.html"
	}

	info, err := os.Stat(path)

	if err != nil {
		return []byte("File not found"), status.HTTP_404_NOT_FOUND, err
	}

	if info.Size() == 0 {
		return []byte("Invalid request"), status.HTTP_400_BAD_REQUEST, err
	}

	file, err := os.Open(path)

	if err != nil {
		return []byte("File not found"), status.HTTP_404_NOT_FOUND, err
	}

	dataBuff := bytes.Buffer{}

	n, err := io.Copy(&dataBuff, file)

	if err != nil {
		return []byte("Something went wrong"), status.HTTP_500_SERVER_ERROR, err
	}

	return dataBuff.Bytes()[:n], status.HTTP_200_OK, nil
}
