package server

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/iamvineettiwari/go-web-server/http"
	"github.com/iamvineettiwari/go-web-server/http/status"
	"github.com/iamvineettiwari/go-web-server/router"
)

type HttpServer struct {
	ListenAddr string
	Listener   net.Listener
	serverLock chan struct{}
	staticPath string
	router     *router.Router
}

func NewHttpServer(address string) *HttpServer {
	return &HttpServer{
		ListenAddr: address,
		serverLock: make(chan struct{}),
		staticPath: filepath.Join(BASE_PATH, "public"),
		router:     router.NewRouter(),
	}
}

func (hs *HttpServer) SetStaticPath(path string) {
	hs.staticPath = path
}

func (hs *HttpServer) GetStaticPath() string {
	return hs.staticPath
}

func (hs *HttpServer) GetRouter() *router.Router {
	return hs.router
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
	response := http.GetResponseWriter(conn)

	if request.Method == http.GET && strings.Contains(request.Path, ".") {
		hs.processStaticRequest(request, response)
		return
	}

	handler, params, err := hs.router.Resolve(request.Method, request.Path)

	if err != nil {
		hs.processStaticRequest(request, response)
		return
	}

	parseErr := request.ParseRequest(params)

	if parseErr != nil {
		response.Send([]byte(parseErr.Error()), status.HTTP_400_BAD_REQUEST)
		return
	}

	handler(request, response)
}

func (hs *HttpServer) processStaticRequest(request *http.Request, response *http.Response) {
	file, err := hs.processFilePath(request.Path)

	if err != nil {
		response.Send([]byte("Page not found"), status.HTTP_404_NOT_FOUND)
		return
	}

	dataBuff := bytes.Buffer{}

	n, err := io.Copy(&dataBuff, file)

	if err != nil {
		response.Send([]byte("Something went wrong"), status.HTTP_500_SERVER_ERROR)
		return
	}

	response.Send(dataBuff.Bytes()[:n], status.HTTP_200_OK)
	return
}

func (hs *HttpServer) processFilePath(requestPath string) (*os.File, error) {
	curWd, wdErr := os.Getwd()

	if wdErr != nil {
		return nil, wdErr
	}

	path := filepath.Join(curWd, filepath.Join(hs.staticPath, requestPath))

	if strings.HasSuffix(requestPath, "/") {
		path = filepath.Join(path, "index.html")
	}

	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		path = filepath.Join(path, "index.html")
	}

	info, err = os.Stat(path)

	if err != nil {
		return nil, err
	}

	if info.Size() == 0 || info.IsDir() {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}
